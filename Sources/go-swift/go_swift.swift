@_cdecl("sayHello") // export to C as `sayHello``
public func sayHello(namePtr: UnsafePointer<CChar>?) {    
    // Creates a new string by copying the null-terminated UTF-8 data (C String) 
    // referenced by the given pointer.
    let name = String(cString: namePtr!)
    print("Welcome, \(name)!")
}
