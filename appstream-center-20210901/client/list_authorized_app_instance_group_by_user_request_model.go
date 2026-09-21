// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAppInstanceGroupByUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetAppInstanceGroupName() *string
	SetAppName(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetAppName() *string
	SetEndUserId(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetEndUserId() *string
	SetPageNumber(v int32) *ListAuthorizedAppInstanceGroupByUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedAppInstanceGroupByUserRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetProductType() *string
}

type ListAuthorizedAppInstanceGroupByUserRequest struct {
	// The application ID. Fuzzy matching is supported: delivery groups that contain a deployed application whose ID includes the specified string are returned. You can obtain the application ID from the Apps list returned by this operation.
	//
	// If this parameter is not specified, no filtering by application ID is applied.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID. Fuzzy matching is supported: delivery groups whose IDs contain the specified string are returned. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the delivery group ID.
	//
	// If this parameter is not specified, no filtering by delivery group ID is applied.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name. Fuzzy matching is supported. For example, if you set this parameter to `Office App`, delivery groups whose names contain `Office App` (such as `My Office App` or `Office App A`) are returned.
	//
	// If this parameter is not specified, no filtering by delivery group name is applied.
	//
	// example:
	//
	// Office App
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application name. Fuzzy matching is supported: delivery groups that contain a deployed application whose name includes the specified string are returned.
	//
	// If this parameter is not specified, no filtering by application name is applied.
	//
	// example:
	//
	// Office App
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The username. An **exact match*	- is performed on the username to query the delivery groups for which the user has been granted delivery group-level authorization.
	//
	// > This parameter is required. If this parameter is not specified, the error code `InvalidParameter.UserId` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of delivery groups to return per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. The value must match the product type of the delivery groups to query. Only delivery groups of the specified product type are returned.
	//
	// Valid values:
	//
	// - CloudApp: WUYING Cloud Application.
	//
	// - CloudBrowser: cloud browser.
	//
	// - WuyingServer: Enterprise Edition workstation.
	//
	// - WuyingWorkstation: Personal Edition Lingjun container workstation.
	//
	// - WuyingWorkstationTeam: Team Edition Lingjun container workstation.
	//
	// - WuyingWorkstationBusiness: Dedicated Edition Lingjun container workstation.
	//
	// - AndroidCloud: cloud phone.
	//
	// - AIAgent: AgentBay (AI agent).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListAuthorizedAppInstanceGroupByUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppInstanceGroupByUserRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetAppId(v string) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetAppInstanceGroupId(v string) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetAppInstanceGroupName(v string) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetAppName(v string) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.AppName = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetEndUserId(v string) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.EndUserId = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetPageNumber(v int32) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetPageSize(v int32) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetProductType(v string) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.ProductType = &v
	return s
}

func (s *ListAuthorizedAppInstanceGroupByUserRequest) Validate() error {
	return dara.Validate(s)
}
