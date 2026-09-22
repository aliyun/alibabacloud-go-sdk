// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAppsByUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAuthorizedAppsByUserRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *ListAuthorizedAppsByUserRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *ListAuthorizedAppsByUserRequest
	GetAppInstanceGroupName() *string
	SetAppName(v string) *ListAuthorizedAppsByUserRequest
	GetAppName() *string
	SetEndUserId(v string) *ListAuthorizedAppsByUserRequest
	GetEndUserId() *string
	SetPageNumber(v int32) *ListAuthorizedAppsByUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedAppsByUserRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAuthorizedAppsByUserRequest
	GetProductType() *string
}

type ListAuthorizedAppsByUserRequest struct {
	// The application ID used to filter results. Fuzzy match by containment is used. This parameter can be combined with other filter parameters. You can obtain the application ID from the Apps list returned by the [GetAppInstanceGroup](https://help.aliyun.com/document_detail/600836.html) operation.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID used to filter results. Fuzzy match by containment is used. This parameter can be combined with other filter parameters. Call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the delivery group ID.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name used to filter results. Fuzzy match by name is used. This parameter can be combined with other filter parameters.
	//
	// example:
	//
	// OfficeApp
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application name used to filter results. Fuzzy match by name is used. This parameter can be combined with other filter parameters.
	//
	// example:
	//
	// OfficeApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The username to query. **Required**. The user must already exist under the current account. Call the [DescribeUsers](https://help.aliyun.com/document_detail/436936.html) operation to obtain the username.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The page number of the results. **Required**. The value starts from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. **Required**. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. **Required**. The value is case-insensitive.
	//
	// This operation queries per-application authorization records. This authorization method applies to WUYING Cloud Application delivery groups. Valid values:
	//
	// - CloudApp: WUYING Cloud Application.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListAuthorizedAppsByUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppsByUserRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppsByUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAuthorizedAppsByUserRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAuthorizedAppsByUserRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListAuthorizedAppsByUserRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAuthorizedAppsByUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListAuthorizedAppsByUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedAppsByUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedAppsByUserRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAuthorizedAppsByUserRequest) SetAppId(v string) *ListAuthorizedAppsByUserRequest {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) SetAppInstanceGroupId(v string) *ListAuthorizedAppsByUserRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) SetAppInstanceGroupName(v string) *ListAuthorizedAppsByUserRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) SetAppName(v string) *ListAuthorizedAppsByUserRequest {
	s.AppName = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) SetEndUserId(v string) *ListAuthorizedAppsByUserRequest {
	s.EndUserId = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) SetPageNumber(v int32) *ListAuthorizedAppsByUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) SetPageSize(v int32) *ListAuthorizedAppsByUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) SetProductType(v string) *ListAuthorizedAppsByUserRequest {
	s.ProductType = &v
	return s
}

func (s *ListAuthorizedAppsByUserRequest) Validate() error {
	return dara.Validate(s)
}
