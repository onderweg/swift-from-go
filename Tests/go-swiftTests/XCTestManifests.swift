import XCTest

#if !os(macOS)
public func allTests() -> [XCTestCaseEntry] {
    return [
        testCase(go_swiftTests.allTests),
    ]
}
#endif