package notion

import (
	"encoding/json"
)

// /api/v3/loadPageChunk request
type loadPageChunkRequest struct {
	PageID          string `json:"pageId"`
	Limit           int    `json:"limit"`
	Cursor          cursor `json:"cursor"`
	VerticalColumns bool   `json:"verticalColumns"`
}

type cursor struct {
	Stack [][]stack `json:"stack"`
}

type stack struct {
	Table string `json:"table"`
	ID    string `json:"id"`
	Index int    `json:"index"`
}

// /api/v3/loadPageChunk response
type loadPageChunkResponse struct {
	RecordMap recordMap `json:"recordMap"`
	Cursor    cursor    `json:"cursor"`
}

type recordMap struct {
	Blocks          map[string]*BlockWithRole          `json:"block"`
	Space           map[string]interface{}             `json:"space"` // TODO: figure out the type
	Users           map[string]*notionUserInfo         `json:"notion_user"`
	Collections     map[string]*CollectionWithRole     `json:"collection"`
	CollectionViews map[string]*CollectionViewWithRole `json:"collection_view"`
	// TDOO: there might be more records types
}

// CollectionViewWithRole describes a role and a collection view
type CollectionViewWithRole struct {
	Role  string         `json:"role"`
	Value CollectionView `json:"value"`
}

// CollectionView describes a collection
type CollectionView struct {
	Alive       bool                  `json:"alive"`
	Format      *CollectionViewFormat `json:"format"`
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	PageSort    []string              `json:"page_sort"`
	ParentID    string                `json:"parent_id"`
	ParentTable string                `json:"parent_table"`
	Query       *CollectionViewQuery  `json:"query"`
	Type        string                `json:"type"`
	Version     int                   `json:"version"`
}

// CollectionViewFormat describes a fomrat of a collection view
type CollectionViewFormat struct {
	TableProperties []TableProperty `json:"table_properties"`
	TableWrap       bool            `json:"table_wrap"`
}

// TableProperty describes a property of the table
type TableProperty struct {
	Property string `json:"property"`
	Visible  bool   `json:"visible"`
	Width    *int   `json:"width,omitempty"`
}

// CollectionViewQuery describes a query
type CollectionViewQuery struct {
	Aggregate []*AggregateQuery `json:"aggregate"`
}

// AggregateQuery describes an aggregate query
type AggregateQuery struct {
	AggregationType string `json:"aggregation_type"`
	ID              string `json:"id"`
	Property        string `json:"property"`
	Type            string `json:"type"`
	ViewType        string `json:"view_type"`
}

// CollectionWithRole describes a collection
type CollectionWithRole struct {
	Role  string              `json:"role"`
	Value *CollectionWithRole `json:"value"`
}

// Collection describes a collection
type Collection struct {
	Alive              bool                             `json:"alive"`
	Format             *CollectionFormat                `json:"format"`
	ID                 string                           `json:"id"`
	Name               [][]string                       `json:"name"`
	ParentID           string                           `json:"parent_id"`
	ParentTable        string                           `json:"parent_table"`
	ColumnNameToSchema map[string]*CollectionColumnInfo `json:"schema"`
	Version            int                              `json:"version"`
}

// CollectionFormat describes format of a collection
type CollectionFormat struct {
	CollectionPageProperties []*CollectionPageProperty `json:"collection_page_properties"`
}

// CollectionPageProperty describes properties of a collection
type CollectionPageProperty struct {
	Property string `json:"property"`
	Visible  bool   `json:"visible"`
}

// CollectionColumnInfo describes a info of a collection column
type CollectionColumnInfo struct {
	Name    string                    `json:"name"`
	Options []*CollectionColumnOption `json:"options"`
	Type    string                    `json:"type"`
}

// CollectionColumnOption describes options for a collection column
type CollectionColumnOption struct {
	Color string `json:"color"`
	ID    string `json:"id"`
	Value string `json:"value"`
}

type notionUserInfo struct {
	Role  string `json:"role"`
	Value *User  `json:"value"`
}

// User describes a user
type User struct {
	Email                     string `json:"email"`
	FamilyName                string `json:"family_name"`
	GivenName                 string `json:"given_name"`
	ID                        string `json:"id"`
	Locale                    string `json:"locale"`
	MobileOnboardingCompleted bool   `json:"mobile_onboarding_completed"`
	OnboardingCompleted       bool   `json:"onboarding_completed"`
	ProfilePhoto              string `json:"profile_photo"`
	TimeZone                  string `json:"time_zone"`
	Version                   int    `json:"version"`
}

func parseLoadPageChunk(d []byte) (*loadPageChunkResponse, error) {
	var rsp loadPageChunkResponse
	err := json.Unmarshal(d, &rsp)
	if err != nil {
		dbg("parseLoadPageChunk: json.Unmarshal() failed with '%s'\n", err)
		return nil, err
	}
	return &rsp, nil
}
