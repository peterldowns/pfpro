import SwiftUI

func loadConfig() -> Config? {
    let decoder = JSONDecoder()
    let raw = config_open()!
    let json_string = String(cString: raw)
    let json_data = json_string.data(using: .utf8)!
    return try? decoder.decode(Config.self, from: json_data)
}

struct Config: Hashable, Codable {
    let Path: String
    let Directives: [Directive]?
    private enum CodingKeys: String, CodingKey {
        case Path, Directives
    }
}

struct Directive: Hashable, Codable, Identifiable {
    var Alias: String
    var Port: Int
    // Allows editing of text fields while inside a dynamic list.
    // Automatically generated at instantiation, never serialized.
    let id: UUID = UUID()
    private enum CodingKeys: String, CodingKey {
        case Alias, Port
    }
}

extension [Directive] {
    // Port descending
    // Name ascending
    func sorted() -> [Directive] {
        return self.sorted(by: {
            ($1.Port, $0.Alias) < ($0.Port, $1.Alias)
        })
    }
}
