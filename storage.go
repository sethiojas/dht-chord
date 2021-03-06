package chord

// dataStore is an alias to map data structure
// with string type keys and string type values
type dataStore map[string]string

// Save a key-value pair
func (data dataStore) set(key, value string) {
	data[key] = value
}

// Return the value associated with the given key
func (data dataStore) get(key string) (string, bool) {
	value, ok := data[key]
	return value, ok
}

// Delete key-value pairs
func (data dataStore) del(keys []string) {
	for _, key := range keys {
		delete(data, key)
	}
}
