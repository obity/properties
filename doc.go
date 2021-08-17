/*
Properties 是一个用于读写属性文件的Go库。Properties 可保存在流中或从流中加载。
属性列表中每个键及其对应值都是一个字符串。它是线程安全的：多个线程可以共享单个 Properties 对象而无需进行外部同步。

Properties is a Go library for reading and writing properties files.
The Properties model represents a persistent set of properties. The Properties can be saved to a stream or loaded from a stream. Each key and its corresponding value in the property list is a string.

It is thread-safe: multiple threads can share a single Properties object without the need for external synchronization.
*/
package properties
