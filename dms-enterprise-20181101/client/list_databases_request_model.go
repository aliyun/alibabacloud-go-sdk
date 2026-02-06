// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDatabasesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDatabasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatabasesRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListDatabasesRequest
	GetSearchKey() *string
	SetTid(v int64) *ListDatabasesRequest
	GetTid() *int64
}

type ListDatabasesRequest struct {
	// The ID of the instance. The valid value is returned if you call the ListInstances operation. The instance ID is not the ID of the RDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 10
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatabasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatabasesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListDatabasesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDatabasesRequest) SetInstanceId(v string) *ListDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesRequest) SetPageNumber(v int32) *ListDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabasesRequest) SetPageSize(v int32) *ListDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabasesRequest) SetSearchKey(v string) *ListDatabasesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListDatabasesRequest) SetTid(v int64) *ListDatabasesRequest {
	s.Tid = &v
	return s
}

func (s *ListDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
