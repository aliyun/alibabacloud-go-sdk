// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListPublishedAppInstanceGroupRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *ListPublishedAppInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *ListPublishedAppInstanceGroupRequest
	GetAppInstanceGroupName() *string
	SetAppName(v string) *ListPublishedAppInstanceGroupRequest
	GetAppName() *string
	SetExcludeUserId(v string) *ListPublishedAppInstanceGroupRequest
	GetExcludeUserId() *string
	SetOfficeSiteId(v string) *ListPublishedAppInstanceGroupRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *ListPublishedAppInstanceGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishedAppInstanceGroupRequest
	GetPageSize() *int32
	SetProductType(v string) *ListPublishedAppInstanceGroupRequest
	GetProductType() *string
}

type ListPublishedAppInstanceGroupRequest struct {
	// The ID of the application used to filter delivery groups. Partial matching is supported. The delivery group must contain a deployed application that meets the condition. If this parameter is not specified or is set to an empty string, the results are not filtered by application ID. If this parameter is specified together with `AppName`, both conditions must be met by the same application.
	//
	// This condition does not restrict the returned `Apps` list to the matched applications.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the delivery group used for filtering. Partial matching is supported. You can pass the complete ID or a continuous segment of it. If this parameter is not specified or is set to an empty string, the results are not filtered by ID. This parameter can be used together with other filter conditions, and the results must meet all conditions.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group used for filtering. Partial matching is supported. For example, if you pass `Office applications`, delivery groups whose names contain this text are matched. If this parameter is not specified or is set to an empty string, the results are not filtered by name. If this parameter is specified together with the delivery group ID, both the ID and the name must be matched.
	//
	// example:
	//
	// Office applications
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The name of the application used to filter delivery groups. Partial matching is supported. The delivery group must contain a deployed application whose name contains the specified text. If this parameter is not specified or is set to an empty string, the results are not filtered by application name. If this parameter is specified together with `AppId`, both conditions must be met by the same application.
	//
	// This condition does not restrict the returned `Apps` list to the matched applications.
	//
	// example:
	//
	// Office applications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The username to be excluded based on existing authorizations. Exact match by full username is supported, such as `alice`. If this parameter is specified, delivery groups for which all applications have been directly authorized to the user are excluded. If this parameter is not specified or is set to an empty string, the results are not filtered by user authorization status.
	//
	// **Authorizations for a single application or access permissions obtained through user groups are not fully evaluated by this condition.*	- Do not equate the returned results with a list of delivery groups to which the user has no access at all.
	//
	// example:
	//
	// alice
	ExcludeUserId *string `json:"ExcludeUserId,omitempty" xml:"ExcludeUserId,omitempty"`
	// The ID of the office network.
	//
	// example:
	//
	// cn-beijing+dir-172301****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number. This parameter is required. Start the query from page `1` and use it in conjunction with `PageSize`. Keep other filter conditions unchanged when querying subsequent pages.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of delivery groups to return per page. This parameter is required. Valid values: `1` to `100`. Explicitly specify this parameter and do not use the default values of other query operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. This parameter is required. The value must be the same as the product type of the delivery groups to be queried. Only published delivery groups of this product type are returned. If an unrecognized value is passed, a parameter error is returned.
	//
	// Valid values:
	//
	// - `CloudApp`: Wuying Cloud App.
	//
	// - `CloudBrowser`: Cloud browser.
	//
	// - `WuyingServer`: Enterprise Edition workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Linggou container workstation.
	//
	// - `WuyingWorkstationTeam`: Team Edition Linggou container workstation.
	//
	// - `WuyingWorkstationBusiness`: Dedicated Edition Linggou container workstation.
	//
	// - `AndroidCloud`: Cloud phone.
	//
	// - `AIAgent`: AgentBay (AI agent).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListPublishedAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInstanceGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedAppInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListPublishedAppInstanceGroupRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListPublishedAppInstanceGroupRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListPublishedAppInstanceGroupRequest) GetExcludeUserId() *string {
	return s.ExcludeUserId
}

func (s *ListPublishedAppInstanceGroupRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListPublishedAppInstanceGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishedAppInstanceGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListPublishedAppInstanceGroupRequest) SetAppId(v string) *ListPublishedAppInstanceGroupRequest {
	s.AppId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *ListPublishedAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *ListPublishedAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetAppName(v string) *ListPublishedAppInstanceGroupRequest {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetExcludeUserId(v string) *ListPublishedAppInstanceGroupRequest {
	s.ExcludeUserId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetOfficeSiteId(v string) *ListPublishedAppInstanceGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetPageNumber(v int32) *ListPublishedAppInstanceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetPageSize(v int32) *ListPublishedAppInstanceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) SetProductType(v string) *ListPublishedAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *ListPublishedAppInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
