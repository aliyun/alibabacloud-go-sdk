// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAuthorizedUsersForInstanceRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListAuthorizedUsersForInstanceRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAuthorizedUsersForInstanceRequest
	GetPageSize() *string
	SetSearchKey(v string) *ListAuthorizedUsersForInstanceRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAuthorizedUsersForInstanceRequest
	GetTid() *int64
}

type ListAuthorizedUsersForInstanceRequest struct {
	// The ID of the instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used for the search.
	//
	// example:
	//
	// poc_test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAuthorizedUsersForInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizedUsersForInstanceRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAuthorizedUsersForInstanceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAuthorizedUsersForInstanceRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAuthorizedUsersForInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorizedUsersForInstanceRequest) SetInstanceId(v string) *ListAuthorizedUsersForInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetPageNumber(v string) *ListAuthorizedUsersForInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetPageSize(v string) *ListAuthorizedUsersForInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetSearchKey(v string) *ListAuthorizedUsersForInstanceRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetTid(v int64) *ListAuthorizedUsersForInstanceRequest {
	s.Tid = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) Validate() error {
	return dara.Validate(s)
}
