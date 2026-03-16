// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckImageExistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CheckImageExistsRequest
	GetInstanceName() *string
	SetPicName(v string) *CheckImageExistsRequest
	GetPicName() *string
	SetProductId(v string) *CheckImageExistsRequest
	GetProductId() *string
}

type CheckImageExistsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s CheckImageExistsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckImageExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckImageExistsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CheckImageExistsRequest) GetPicName() *string {
	return s.PicName
}

func (s *CheckImageExistsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CheckImageExistsRequest) SetInstanceName(v string) *CheckImageExistsRequest {
	s.InstanceName = &v
	return s
}

func (s *CheckImageExistsRequest) SetPicName(v string) *CheckImageExistsRequest {
	s.PicName = &v
	return s
}

func (s *CheckImageExistsRequest) SetProductId(v string) *CheckImageExistsRequest {
	s.ProductId = &v
	return s
}

func (s *CheckImageExistsRequest) Validate() error {
	return dara.Validate(s)
}
