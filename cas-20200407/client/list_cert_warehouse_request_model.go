// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListCertWarehouseRequest
	GetCurrentPage() *int64
	SetInstanceId(v string) *ListCertWarehouseRequest
	GetInstanceId() *string
	SetName(v string) *ListCertWarehouseRequest
	GetName() *string
	SetShowSize(v int64) *ListCertWarehouseRequest
	GetShowSize() *int64
	SetType(v string) *ListCertWarehouseRequest
	GetType() *string
}

type ListCertWarehouseRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the certificate application repository.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the certificate application repository. Fuzzy match is supported.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The type of the certificate application repository. Valid values:
	//
	// 	- **ssl**: certificate application repository of SSL certificates
	//
	// 	- **uploadPCA**: certificate application repository of uploaded private certificates
	//
	// 	- **free**: certificate application repository of free certificates, available only on the China site (aliyun.com)
	//
	// 	- **aliyunPCA**: certificate application repository of private certificates purchased from Alibaba Cloud Private Certificate Authority (PCA), available only on the China site (aliyun.com)
	//
	// 	- **disable**: disabled certificate application repository
	//
	// example:
	//
	// aliyunPCA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCertWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertWarehouseRequest) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListCertWarehouseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCertWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *ListCertWarehouseRequest) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListCertWarehouseRequest) GetType() *string {
	return s.Type
}

func (s *ListCertWarehouseRequest) SetCurrentPage(v int64) *ListCertWarehouseRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCertWarehouseRequest) SetInstanceId(v string) *ListCertWarehouseRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCertWarehouseRequest) SetName(v string) *ListCertWarehouseRequest {
	s.Name = &v
	return s
}

func (s *ListCertWarehouseRequest) SetShowSize(v int64) *ListCertWarehouseRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCertWarehouseRequest) SetType(v string) *ListCertWarehouseRequest {
	s.Type = &v
	return s
}

func (s *ListCertWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
