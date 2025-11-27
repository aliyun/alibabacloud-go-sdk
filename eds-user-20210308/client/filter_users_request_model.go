// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeEndUserIds(v []*string) *FilterUsersRequest
	GetExcludeEndUserIds() []*string
	SetFilter(v string) *FilterUsersRequest
	GetFilter() *string
	SetIncludeDesktopCount(v bool) *FilterUsersRequest
	GetIncludeDesktopCount() *bool
	SetIncludeDesktopGroupCount(v bool) *FilterUsersRequest
	GetIncludeDesktopGroupCount() *bool
	SetIncludeOrgInfo(v bool) *FilterUsersRequest
	GetIncludeOrgInfo() *bool
	SetIncludeSupportIdps(v bool) *FilterUsersRequest
	GetIncludeSupportIdps() *bool
	SetIsQueryAllSubOrgs(v bool) *FilterUsersRequest
	GetIsQueryAllSubOrgs() *bool
	SetMaxResults(v int64) *FilterUsersRequest
	GetMaxResults() *int64
	SetNextToken(v string) *FilterUsersRequest
	GetNextToken() *string
	SetOrderParam(v *FilterUsersRequestOrderParam) *FilterUsersRequest
	GetOrderParam() *FilterUsersRequestOrderParam
	SetOrgId(v string) *FilterUsersRequest
	GetOrgId() *string
	SetOwnerType(v string) *FilterUsersRequest
	GetOwnerType() *string
	SetPropertyFilterParam(v []*FilterUsersRequestPropertyFilterParam) *FilterUsersRequest
	GetPropertyFilterParam() []*FilterUsersRequestPropertyFilterParam
	SetPropertyKeyValueFilterParam(v []*FilterUsersRequestPropertyKeyValueFilterParam) *FilterUsersRequest
	GetPropertyKeyValueFilterParam() []*FilterUsersRequestPropertyKeyValueFilterParam
	SetStatus(v int32) *FilterUsersRequest
	GetStatus() *int32
}

type FilterUsersRequest struct {
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
	// Specifies whether to return the organization information.
	IncludeOrgInfo *bool `json:"IncludeOrgInfo,omitempty" xml:"IncludeOrgInfo,omitempty"`
	// Specifies whether to return the supported logon types.
	IncludeSupportIdps *bool `json:"IncludeSupportIdps,omitempty" xml:"IncludeSupportIdps,omitempty"`
	// Specifies whether to query all sub-organizations.
	IsQueryAllSubOrgs *bool `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
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
	OrderParam *FilterUsersRequestOrderParam `json:"OrderParam,omitempty" xml:"OrderParam,omitempty" type:"Struct"`
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
	PropertyFilterParam []*FilterUsersRequestPropertyFilterParam `json:"PropertyFilterParam,omitempty" xml:"PropertyFilterParam,omitempty" type:"Repeated"`
	// The list of property names and property values.
	PropertyKeyValueFilterParam []*FilterUsersRequestPropertyKeyValueFilterParam `json:"PropertyKeyValueFilterParam,omitempty" xml:"PropertyKeyValueFilterParam,omitempty" type:"Repeated"`
	// The status.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FilterUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersRequest) GoString() string {
	return s.String()
}

func (s *FilterUsersRequest) GetExcludeEndUserIds() []*string {
	return s.ExcludeEndUserIds
}

func (s *FilterUsersRequest) GetFilter() *string {
	return s.Filter
}

func (s *FilterUsersRequest) GetIncludeDesktopCount() *bool {
	return s.IncludeDesktopCount
}

func (s *FilterUsersRequest) GetIncludeDesktopGroupCount() *bool {
	return s.IncludeDesktopGroupCount
}

func (s *FilterUsersRequest) GetIncludeOrgInfo() *bool {
	return s.IncludeOrgInfo
}

func (s *FilterUsersRequest) GetIncludeSupportIdps() *bool {
	return s.IncludeSupportIdps
}

func (s *FilterUsersRequest) GetIsQueryAllSubOrgs() *bool {
	return s.IsQueryAllSubOrgs
}

func (s *FilterUsersRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *FilterUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *FilterUsersRequest) GetOrderParam() *FilterUsersRequestOrderParam {
	return s.OrderParam
}

func (s *FilterUsersRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *FilterUsersRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *FilterUsersRequest) GetPropertyFilterParam() []*FilterUsersRequestPropertyFilterParam {
	return s.PropertyFilterParam
}

func (s *FilterUsersRequest) GetPropertyKeyValueFilterParam() []*FilterUsersRequestPropertyKeyValueFilterParam {
	return s.PropertyKeyValueFilterParam
}

func (s *FilterUsersRequest) GetStatus() *int32 {
	return s.Status
}

func (s *FilterUsersRequest) SetExcludeEndUserIds(v []*string) *FilterUsersRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *FilterUsersRequest) SetFilter(v string) *FilterUsersRequest {
	s.Filter = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeDesktopCount(v bool) *FilterUsersRequest {
	s.IncludeDesktopCount = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeDesktopGroupCount(v bool) *FilterUsersRequest {
	s.IncludeDesktopGroupCount = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeOrgInfo(v bool) *FilterUsersRequest {
	s.IncludeOrgInfo = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeSupportIdps(v bool) *FilterUsersRequest {
	s.IncludeSupportIdps = &v
	return s
}

func (s *FilterUsersRequest) SetIsQueryAllSubOrgs(v bool) *FilterUsersRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *FilterUsersRequest) SetMaxResults(v int64) *FilterUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *FilterUsersRequest) SetNextToken(v string) *FilterUsersRequest {
	s.NextToken = &v
	return s
}

func (s *FilterUsersRequest) SetOrderParam(v *FilterUsersRequestOrderParam) *FilterUsersRequest {
	s.OrderParam = v
	return s
}

func (s *FilterUsersRequest) SetOrgId(v string) *FilterUsersRequest {
	s.OrgId = &v
	return s
}

func (s *FilterUsersRequest) SetOwnerType(v string) *FilterUsersRequest {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersRequest) SetPropertyFilterParam(v []*FilterUsersRequestPropertyFilterParam) *FilterUsersRequest {
	s.PropertyFilterParam = v
	return s
}

func (s *FilterUsersRequest) SetPropertyKeyValueFilterParam(v []*FilterUsersRequestPropertyKeyValueFilterParam) *FilterUsersRequest {
	s.PropertyKeyValueFilterParam = v
	return s
}

func (s *FilterUsersRequest) SetStatus(v int32) *FilterUsersRequest {
	s.Status = &v
	return s
}

func (s *FilterUsersRequest) Validate() error {
	if s.OrderParam != nil {
		if err := s.OrderParam.Validate(); err != nil {
			return err
		}
	}
	if s.PropertyFilterParam != nil {
		for _, item := range s.PropertyFilterParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PropertyKeyValueFilterParam != nil {
		for _, item := range s.PropertyKeyValueFilterParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FilterUsersRequestOrderParam struct {
	// The field that you want to sort by.
	//
	// Valid values:
	//
	// 	- EndUserId: the username.
	//
	// 	- id: the ID of the primary key.
	//
	// 	- gmt_created: the creation time.
	//
	// example:
	//
	// id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The direction of the sort.
	//
	// Valid values:
	//
	// 	- ASC: the ascending order.
	//
	// 	- DESC (default): the descending order.
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s FilterUsersRequestOrderParam) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersRequestOrderParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestOrderParam) GetOrderField() *string {
	return s.OrderField
}

func (s *FilterUsersRequestOrderParam) GetOrderType() *string {
	return s.OrderType
}

func (s *FilterUsersRequestOrderParam) SetOrderField(v string) *FilterUsersRequestOrderParam {
	s.OrderField = &v
	return s
}

func (s *FilterUsersRequestOrderParam) SetOrderType(v string) *FilterUsersRequestOrderParam {
	s.OrderType = &v
	return s
}

func (s *FilterUsersRequestOrderParam) Validate() error {
	return dara.Validate(s)
}

type FilterUsersRequestPropertyFilterParam struct {
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

func (s FilterUsersRequestPropertyFilterParam) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersRequestPropertyFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestPropertyFilterParam) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *FilterUsersRequestPropertyFilterParam) GetPropertyValueIds() *string {
	return s.PropertyValueIds
}

func (s *FilterUsersRequestPropertyFilterParam) SetPropertyId(v int64) *FilterUsersRequestPropertyFilterParam {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersRequestPropertyFilterParam) SetPropertyValueIds(v string) *FilterUsersRequestPropertyFilterParam {
	s.PropertyValueIds = &v
	return s
}

func (s *FilterUsersRequestPropertyFilterParam) Validate() error {
	return dara.Validate(s)
}

type FilterUsersRequestPropertyKeyValueFilterParam struct {
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

func (s FilterUsersRequestPropertyKeyValueFilterParam) String() string {
	return dara.Prettify(s)
}

func (s FilterUsersRequestPropertyKeyValueFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) GetPropertyValues() *string {
	return s.PropertyValues
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) SetPropertyKey(v string) *FilterUsersRequestPropertyKeyValueFilterParam {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) SetPropertyValues(v string) *FilterUsersRequestPropertyKeyValueFilterParam {
	s.PropertyValues = &v
	return s
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) Validate() error {
	return dara.Validate(s)
}
