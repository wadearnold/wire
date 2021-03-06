/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Personal is personal demographic information
type Personal struct {
	// Identification Code:  * `1` - Passport Number * `2` - Tax Identification Number * `3` - Driver’s License Number * `4` - Alien Registration Number * `5` - Corporate Identification * `9` - Other Identification 
	IdentificationCode string `json:"identificationCode"`
	// Identifier
	Identifier string `json:"identifier"`
	// Name
	Name string `json:"name"`
	Address Address `json:"address"`
}
