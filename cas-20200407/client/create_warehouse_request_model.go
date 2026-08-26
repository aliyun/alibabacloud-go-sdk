// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBiz(v string) *CreateWarehouseRequest
	GetBiz() *string
	SetName(v string) *CreateWarehouseRequest
	GetName() *string
	SetType(v string) *CreateWarehouseRequest
	GetType() *string
}

type CreateWarehouseRequest struct {
	// The scenarios of the repository.
	//
	// example:
	//
	// contract
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The name of the certificate repository.
	//
	// example:
	//
	// MyCertificateWarehouse
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the certificate repository.
	//
	// example:
	//
	// pcaCertificate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseRequest) GoString() string {
	return s.String()
}

func (s *CreateWarehouseRequest) GetBiz() *string {
	return s.Biz
}

func (s *CreateWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *CreateWarehouseRequest) GetType() *string {
	return s.Type
}

func (s *CreateWarehouseRequest) SetBiz(v string) *CreateWarehouseRequest {
	s.Biz = &v
	return s
}

func (s *CreateWarehouseRequest) SetName(v string) *CreateWarehouseRequest {
	s.Name = &v
	return s
}

func (s *CreateWarehouseRequest) SetType(v string) *CreateWarehouseRequest {
	s.Type = &v
	return s
}

func (s *CreateWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
