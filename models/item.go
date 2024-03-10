package models

// Item represents an item in the Alfamidi inventory
type Item struct {
	ItemID      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id,omitempty"` // Optional for newItem
}

// Order represents a customer order at Alfamidi
type Order struct {
	OrderID      int    `json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"` // Use string for flexibility
	Items        []Item `json:"items"`
}
