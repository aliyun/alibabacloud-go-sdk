// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListChartNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *ListChartNamespaceRequest
	GetNamespaceName() *string
	SetNamespaceStatus(v string) *ListChartNamespaceRequest
	GetNamespaceStatus() *string
	SetPageNo(v int32) *ListChartNamespaceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListChartNamespaceRequest
	GetPageSize() *int32
}

type ListChartNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`: The namespace is normal.
	//
	// 	- `DELETING`: The namespace is being deleted.
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChartNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChartNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListChartNamespaceRequest) GetNamespaceStatus() *string {
	return s.NamespaceStatus
}

func (s *ListChartNamespaceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChartNamespaceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChartNamespaceRequest) SetInstanceId(v string) *ListChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChartNamespaceRequest) SetNamespaceName(v string) *ListChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListChartNamespaceRequest) SetNamespaceStatus(v string) *ListChartNamespaceRequest {
	s.NamespaceStatus = &v
	return s
}

func (s *ListChartNamespaceRequest) SetPageNo(v int32) *ListChartNamespaceRequest {
	s.PageNo = &v
	return s
}

func (s *ListChartNamespaceRequest) SetPageSize(v int32) *ListChartNamespaceRequest {
	s.PageSize = &v
	return s
}

func (s *ListChartNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
