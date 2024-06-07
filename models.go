package semrush

var url string = `https://api.semrush.com/apis/v4-raw/`

//////////////////////////////////////////
/*			UPDATE REQUESTS				*/
//////////////////////////////////////////

type EditLocations struct {
	Locations []EditLocation `json:"locations"`
}

type EditLocation struct {
	LocationID            string                     `json:"-"`
	Name                  string                     `json:"locationName"`
	City                  string                     `json:"city"`
	Address               string                     `json:"address"`
	Phone                 string                     `json:"phone"`
	AdditionalAddressInfo *string                    `json:"additionalAddressInfo"`
	Zip                   *string                    `json:"zip"`
	Region                *string                    `json:"region"`
	WebsiteURL            *string                    `json:"websiteUrl"`
	ReOpenDate            *string                    `json:"reopenDate"`
	BusinessHours         map[string][]ScheduleTimes `json:"businessHours"`
	HolidayHours          *[]HolidayHours            `json:"holidayHours"`
}

//////////////////////////////////////////
/*			BROWSE REQUESTS				*/
//////////////////////////////////////////

type BrowseLocation struct {
	Page *string
	Size *string
}

type ReadLocation struct {
	LocationID string
}

//////////////////////////////////////////
/*			RESPONSE STRUCTS			*/
//////////////////////////////////////////

type LocationResponse struct {
	RequestID string   `json:"requestId"`
	Data      Location `json:"data"`
	Error     Error    `json:"error"`
}

type LocationsResponse struct {
	RequestID string        `json:"requestId"`
	Data      LocationsData `json:"data"`
	Error     Error         `json:"error"`
}

//////////////////////////////////////////
/*			OTHER STRUCTS				*/
//////////////////////////////////////////

type Error struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Details *[]ErrorDetails `json:"details"`
}

type ErrorDetails struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Field   *string `json:"field"`
	Index   *string `json:"index"`
}

type LocationsData struct {
	TotalElements *int       `json:"totalElements"`
	Page          *int       `json:"page"`
	TotalPages    *int       `json:"totalPages"`
	Content       []Location `json:"content"`
}

type Location struct {
	ID                    string                     `json:"id"`
	Name                  string                     `json:"locationName"`
	City                  string                     `json:"city"`
	Address               string                     `json:"address"`
	AdditionalAddressInfo string                     `json:"additionalAddressInfo"`
	Phone                 string                     `json:"phone"`
	Zip                   *string                    `json:"zip"`
	Region                *string                    `json:"region"`
	WebsiteURL            *string                    `json:"websiteUrl"`
	ReOpenDate            *string                    `json:"reopenDate"`
	BusinessHours         map[string][]ScheduleTimes `json:"businessHours"`
	HolidayHours          *[]HolidayHours            `json:"holidayHours"`
}

type HolidayHours struct {
	Type  string           `json:"type"`
	Day   string           `json:"day"`
	Times *[]ScheduleTimes `json:"times"`
}

type ScheduleTimes struct {
	From string `json:"from"`
	To   string `json:"to"`
}
