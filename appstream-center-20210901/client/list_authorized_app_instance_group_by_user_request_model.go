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
	SetOfficeSiteId(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *ListAuthorizedAppInstanceGroupByUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedAppInstanceGroupByUserRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAuthorizedAppInstanceGroupByUserRequest
	GetProductType() *string
}

type ListAuthorizedAppInstanceGroupByUserRequest struct {
	// The application ID. Fuzzy match is supported. Delivery groups in which the IDs of deployed applications contain this string are returned. You can obtain the application ID from the Apps list returned by this operation.
	//
	// If this parameter is not specified, delivery groups are not filtered by application ID.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID. Fuzzy match is supported. Delivery groups whose IDs contain this string are returned. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID.
	//
	// If this parameter is not specified, delivery groups are not filtered by delivery group ID.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name. Fuzzy match is supported. For example, if you specify Office applications, delivery groups whose names contain Office applications (such as My office applications and Office applications A) are returned.
	//
	// If this parameter is not specified, delivery groups are not filtered by delivery group name.
	//
	// example:
	//
	// Office applications
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application name. Fuzzy match is supported. Delivery groups in which the names of deployed applications contain this string are returned.
	//
	// If this parameter is not specified, delivery groups are not filtered by application name.
	//
	// example:
	//
	// Office applications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The username. Exact match is used to query the delivery groups for which the user has been granted delivery group-level authorization.
	//
	// > This parameter is required. If this parameter is not specified, the InvalidParameter.UserId error code is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-shenzhen+dir-436909****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number. Pages start from 1.
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
	// The product type. The value must be the same as the product type of the delivery groups to be queried. Only delivery groups of this product type are returned.
	//
	// Valid values:
	//
	// - CloudApp: Wuying Cloud App.
	//
	// - CloudBrowser: Cloud Browser.
	//
	// - WuyingServer: Enterprise Edition workstation.
	//
	// - WuyingWorkstation: Personal Edition Linggou container workstation.
	//
	// - WuyingWorkstationTeam: Team Edition Linggou container workstation.
	//
	// - WuyingWorkstationBusiness: Exclusive Edition Linggou container workstation.
	//
	// - AndroidCloud: Cloud Phone.
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

func (s *ListAuthorizedAppInstanceGroupByUserRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
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

func (s *ListAuthorizedAppInstanceGroupByUserRequest) SetOfficeSiteId(v string) *ListAuthorizedAppInstanceGroupByUserRequest {
	s.OfficeSiteId = &v
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
