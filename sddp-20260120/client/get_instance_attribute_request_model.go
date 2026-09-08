// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceAttributeRequest
	GetInstanceId() *string
	SetLang(v string) *GetInstanceAttributeRequest
	GetLang() *string
	SetProductCode(v string) *GetInstanceAttributeRequest
	GetProductCode() *string
	SetProductId(v int64) *GetInstanceAttributeRequest
	GetProductId() *int64
}

type GetInstanceAttributeRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId   *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s GetInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceAttributeRequest) GetLang() *string {
	return s.Lang
}

func (s *GetInstanceAttributeRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetInstanceAttributeRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetInstanceAttributeRequest) SetInstanceId(v string) *GetInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceAttributeRequest) SetLang(v string) *GetInstanceAttributeRequest {
	s.Lang = &v
	return s
}

func (s *GetInstanceAttributeRequest) SetProductCode(v string) *GetInstanceAttributeRequest {
	s.ProductCode = &v
	return s
}

func (s *GetInstanceAttributeRequest) SetProductId(v int64) *GetInstanceAttributeRequest {
	s.ProductId = &v
	return s
}

func (s *GetInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
