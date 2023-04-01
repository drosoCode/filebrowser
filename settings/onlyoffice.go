package settings

// OnlyOffice contains the onlyoffice server connection settings of the app.
type OnlyOffice struct {
	Url string `json:"url"`
	JWT string `json:"jwt"`
}
