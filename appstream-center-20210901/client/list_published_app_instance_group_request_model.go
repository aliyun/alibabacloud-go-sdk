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
	SetPageNumber(v int32) *ListPublishedAppInstanceGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishedAppInstanceGroupRequest
	GetPageSize() *int32
	SetProductType(v string) *ListPublishedAppInstanceGroupRequest
	GetProductType() *string
}

type ListPublishedAppInstanceGroupRequest struct {
	// The application ID used for filtering delivery groups. Substring matching is supported. The delivery group must contain a deployed application that matches the condition. If this parameter is not specified or is set to an empty string, no filtering by application ID is applied. When specified together with `AppName`, the same application must satisfy both conditions.
	//
	// This condition does not trim the returned `Apps` list.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID used for filtering. Substring matching is supported. You can pass in a full ID or a consecutive segment of the ID. If this parameter is not specified or is set to an empty string, no filtering by ID is applied. This parameter can be used together with other filter conditions. Results must satisfy all conditions simultaneously.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name used for filtering. Substring matching is supported. For example, if you pass in `OfficeApps`, delivery groups whose names contain this text are matched. If this parameter is not specified or is set to an empty string, no filtering by name is applied. When specified together with the delivery group ID, both the ID and name must match.
	//
	// example:
	//
	// OfficeApps
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application name used for filtering delivery groups. Substring matching is supported. The delivery group must contain a deployed application whose name includes the specified text. If this parameter is not specified or is set to an empty string, no filtering by application name is applied. When specified together with `AppId`, the same application must satisfy both conditions.
	//
	// This condition does not trim the returned `Apps` list.
	//
	// example:
	//
	// OfficeApps
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The username to exclude based on existing authorization. Exact username matching is used, for example, `alice`. When specified, delivery groups in which all applications have been directly authorized to this user are excluded. If this parameter is not specified or is set to an empty string, no exclusion based on user authorization is applied.
	//
	// **Authorization granted for individual applications only, or access permissions obtained through user groups, is not fully evaluated by this condition.*	- Do not treat the returned results as a complete list of delivery groups that the user has no access permissions to.
	//
	// example:
	//
	// alice
	ExcludeUserId *string `json:"ExcludeUserId,omitempty" xml:"ExcludeUserId,omitempty"`
	// The page number. This parameter is required. Start from page `1` and use this parameter together with `PageSize`. Keep other filter conditions unchanged when querying subsequent pages.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of delivery groups to return per page. This parameter is required. Valid values: `1` to `100`. Unit: delivery groups. Specify this value explicitly and do not rely on default values from other query operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. This parameter is required. The value must match the product type of the delivery groups you want to query. Only published delivery groups of the specified product type are returned. A parameter error is returned if an unrecognized value is passed in.
	//
	// Valid values:
	//
	// - `CloudApp`: WUYING Cloud Application.
	//
	// - `CloudBrowser`: Cloud Browser.
	//
	// - `WuyingServer`: Enterprise Workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Lingjou Container Workstation.
	//
	// - `WuyingWorkstationTeam`: Team Edition Lingjou Container Workstation.
	//
	// - `WuyingWorkstationBusiness`: Dedicated Edition Lingjou Container Workstation.
	//
	// - `AndroidCloud`: Cloud Phone.
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
