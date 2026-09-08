// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataMaskingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CheckDataMaskingInstanceRequest
	GetInstanceId() *string
	SetLang(v string) *CheckDataMaskingInstanceRequest
	GetLang() *string
	SetProductCode(v string) *CheckDataMaskingInstanceRequest
	GetProductCode() *string
	SetProductId(v int64) *CheckDataMaskingInstanceRequest
	GetProductId() *int64
}

type CheckDataMaskingInstanceRequest struct {
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s CheckDataMaskingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataMaskingInstanceRequest) GoString() string {
	return s.String()
}

func (s *CheckDataMaskingInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckDataMaskingInstanceRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckDataMaskingInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CheckDataMaskingInstanceRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *CheckDataMaskingInstanceRequest) SetInstanceId(v string) *CheckDataMaskingInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckDataMaskingInstanceRequest) SetLang(v string) *CheckDataMaskingInstanceRequest {
	s.Lang = &v
	return s
}

func (s *CheckDataMaskingInstanceRequest) SetProductCode(v string) *CheckDataMaskingInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CheckDataMaskingInstanceRequest) SetProductId(v int64) *CheckDataMaskingInstanceRequest {
	s.ProductId = &v
	return s
}

func (s *CheckDataMaskingInstanceRequest) Validate() error {
	return dara.Validate(s)
}
