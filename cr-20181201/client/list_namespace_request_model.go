// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *ListNamespaceRequest
	GetNamespaceName() *string
	SetNamespaceStatus(v string) *ListNamespaceRequest
	GetNamespaceStatus() *string
	SetPageNo(v int32) *ListNamespaceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListNamespaceRequest
	GetPageSize() *int32
}

type ListNamespaceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-94klsruryslx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// test-namespace
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`
	//
	// 	- `DELETING`
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespaceRequest) GoString() string {
	return s.String()
}

func (s *ListNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListNamespaceRequest) GetNamespaceStatus() *string {
	return s.NamespaceStatus
}

func (s *ListNamespaceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListNamespaceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespaceRequest) SetInstanceId(v string) *ListNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNamespaceRequest) SetNamespaceName(v string) *ListNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListNamespaceRequest) SetNamespaceStatus(v string) *ListNamespaceRequest {
	s.NamespaceStatus = &v
	return s
}

func (s *ListNamespaceRequest) SetPageNo(v int32) *ListNamespaceRequest {
	s.PageNo = &v
	return s
}

func (s *ListNamespaceRequest) SetPageSize(v int32) *ListNamespaceRequest {
	s.PageSize = &v
	return s
}

func (s *ListNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
