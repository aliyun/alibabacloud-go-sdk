// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeEndUserIds(v []*string) *FilterUsersShrinkRequest
	GetExcludeEndUserIds() []*string
	SetFilter(v string) *FilterUsersShrinkRequest
	GetFilter() *string
	SetIncludeDesktopCount(v bool) *FilterUsersShrinkRequest
	GetIncludeDesktopCount() *bool
	SetIncludeDesktopGroupCount(v bool) *FilterUsersShrinkRequest
	GetIncludeDesktopGroupCount() *bool
	SetIncludeOrgInfo(v bool) *FilterUsersShrinkRequest
	GetIncludeOrgInfo() *bool
	SetIncludeSupportIdps(v bool) *FilterUsersShrinkRequest
	GetIncludeSupportIdps() *bool
	SetIsQueryAllSubOrgs(v bool) *FilterUsersShrinkRequest
	GetIsQueryAllSubOrgs() *bool
	SetMaxResults(v int64) *FilterUsersShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *FilterUsersShrinkRequest
	GetNextToken() *string
	SetOrderParamShrink(v string) *FilterUsersShrinkRequest
	GetOrderParamShrink() *string
	SetOrgId(v string) *FilterUsersShrinkRequest
	GetOrgId() *string
	SetOwnerType(v string) *FilterUsersShrinkRequest
	GetOwnerType() *string
	SetPropertyFilterParam(v []*FilterUsersShrinkRequestPropertyFilterParam) *FilterUsersShrinkRequest
	GetPropertyFilterParam() []*FilterUsersShrinkRequestPropertyFilterParam
	SetPropertyKeyValueFilterParam(v []*FilterUsersShrinkRequestPropertyKeyValueFilterParam) *FilterUsersShrinkRequest
	GetPropertyKeyValueFilterParam() []*FilterUsersShrinkRequestPropertyKeyValueFilterParam
	SetStatus(v int32) *FilterUsersShrinkRequest
	GetStatus() *int32
}

type FilterUsersShrinkRequest struct {
	// The list of usernames to be precisely excluded.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	// The string that is used for fuzzy search. You can use usernames and email addresses to perform fuzzy search. Wildcard characters (\\*) are supported for this parameter. For example, if you set this parameter to a\\*m, the usernames or an email addresses that start with a or end with m are returned.
	//
	// example:
	//
	// test
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return the number of cloud desktops that are assigned to the convenience user.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	IncludeDesktopCount *bool `json:"IncludeDesktopCount,omitempty" xml:"IncludeDesktopCount,omitempty"`
	// Specifies whether to return the number of cloud desktop pools that are assigned to the convenience user.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	IncludeDesktopGroupCount *bool `json:"IncludeDesktopGroupCount,omitempty" xml:"IncludeDesktopGroupCount,omitempty"`
	IncludeOrgInfo           *bool `json:"IncludeOrgInfo,omitempty" xml:"IncludeOrgInfo,omitempty"`
	IncludeSupportIdps       *bool `json:"IncludeSupportIdps,omitempty" xml:"IncludeSupportIdps,omitempty"`
	IsQueryAllSubOrgs        *bool `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
	// The number of entries per page. If you set this parameter to a value greater than 100, the system resets the value to 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The parameter that might affect the sorting logic.
	OrderParamShrink *string `json:"OrderParam,omitempty" xml:"OrderParam,omitempty"`
	// The ID of the organization.
	//
	// example:
	//
	// org-aliyun-wy-org-id
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The activation type of the convenience account.
	//
	// Valid values:
	//
	// 	- CreateFromManager: administrator-activated.
	//
	// 	- Normal: user-activated.
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The list of properties for fuzzy search.
	PropertyFilterParam []*FilterUsersShrinkRequestPropertyFilterParam `json:"PropertyFilterParam,omitempty" xml:"PropertyFilterParam,omitempty" type:"Repeated"`
	// The list of property names and property values.
	PropertyKeyValueFilterParam []*FilterUsersShrinkRequestPropertyKeyValueFilterParam `json:"PropertyKeyValueFilterParam,omitempty" xml:"PropertyKeyValueFilterParam,omitempty" type:"Repeated"`
	Status                      *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FilterUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequest) GetExcludeEndUserIds() []*string {
	return s.ExcludeEndUserIds
}

func (s *FilterUsersShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *FilterUsersShrinkRequest) GetIncludeDesktopCount() *bool {
	return s.IncludeDesktopCount
}

func (s *FilterUsersShrinkRequest) GetIncludeDesktopGroupCount() *bool {
	return s.IncludeDesktopGroupCount
}

func (s *FilterUsersShrinkRequest) GetIncludeOrgInfo() *bool {
	return s.IncludeOrgInfo
}

func (s *FilterUsersShrinkRequest) GetIncludeSupportIdps() *bool {
	return s.IncludeSupportIdps
}

func (s *FilterUsersShrinkRequest) GetIsQueryAllSubOrgs() *bool {
	return s.IsQueryAllSubOrgs
}

func (s *FilterUsersShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *FilterUsersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *FilterUsersShrinkRequest) GetOrderParamShrink() *string {
	return s.OrderParamShrink
}

func (s *FilterUsersShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *FilterUsersShrinkRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *FilterUsersShrinkRequest) GetPropertyFilterParam() []*FilterUsersShrinkRequestPropertyFilterParam {
	return s.PropertyFilterParam
}

func (s *FilterUsersShrinkRequest) GetPropertyKeyValueFilterParam() []*FilterUsersShrinkRequestPropertyKeyValueFilterParam {
	return s.PropertyKeyValueFilterParam
}

func (s *FilterUsersShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *FilterUsersShrinkRequest) SetExcludeEndUserIds(v []*string) *FilterUsersShrinkRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *FilterUsersShrinkRequest) SetFilter(v string) *FilterUsersShrinkRequest {
	s.Filter = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeDesktopCount(v bool) *FilterUsersShrinkRequest {
	s.IncludeDesktopCount = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeDesktopGroupCount(v bool) *FilterUsersShrinkRequest {
	s.IncludeDesktopGroupCount = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeOrgInfo(v bool) *FilterUsersShrinkRequest {
	s.IncludeOrgInfo = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeSupportIdps(v bool) *FilterUsersShrinkRequest {
	s.IncludeSupportIdps = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIsQueryAllSubOrgs(v bool) *FilterUsersShrinkRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetMaxResults(v int64) *FilterUsersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetNextToken(v string) *FilterUsersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOrderParamShrink(v string) *FilterUsersShrinkRequest {
	s.OrderParamShrink = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOrgId(v string) *FilterUsersShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOwnerType(v string) *FilterUsersShrinkRequest {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetPropertyFilterParam(v []*FilterUsersShrinkRequestPropertyFilterParam) *FilterUsersShrinkRequest {
	s.PropertyFilterParam = v
	return s
}

func (s *FilterUsersShrinkRequest) SetPropertyKeyValueFilterParam(v []*FilterUsersShrinkRequestPropertyKeyValueFilterParam) *FilterUsersShrinkRequest {
	s.PropertyKeyValueFilterParam = v
	return s
}

func (s *FilterUsersShrinkRequest) SetStatus(v int32) *FilterUsersShrinkRequest {
	s.Status = &v
	return s
}

func (s *FilterUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type FilterUsersShrinkRequestPropertyFilterParam struct {
	// The ID of the property.
	//
	// example:
	//
	// 123
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The IDs of the property values.
	//
	// example:
	//
	// test
	PropertyValueIds *string `json:"PropertyValueIds,omitempty" xml:"PropertyValueIds,omitempty"`
}

func (s FilterUsersShrinkRequestPropertyFilterParam) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersShrinkRequestPropertyFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) GetPropertyValueIds() *string {
	return s.PropertyValueIds
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) SetPropertyId(v int64) *FilterUsersShrinkRequestPropertyFilterParam {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) SetPropertyValueIds(v string) *FilterUsersShrinkRequestPropertyFilterParam {
	s.PropertyValueIds = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) Validate() error {
	return dara.Validate(s)
}

type FilterUsersShrinkRequestPropertyKeyValueFilterParam struct {
	// The property name.
	//
	// example:
	//
	// job
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property values.
	//
	// example:
	//
	// dev
	PropertyValues *string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty"`
}

func (s FilterUsersShrinkRequestPropertyKeyValueFilterParam) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersShrinkRequestPropertyKeyValueFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) GetPropertyValues() *string {
	return s.PropertyValues
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) SetPropertyKey(v string) *FilterUsersShrinkRequestPropertyKeyValueFilterParam {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) SetPropertyValues(v string) *FilterUsersShrinkRequestPropertyKeyValueFilterParam {
	s.PropertyValues = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) Validate() error {
	return dara.Validate(s)
}
