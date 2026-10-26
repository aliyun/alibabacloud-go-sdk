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
	// The database instance identity. You can obtain the value from Items[].InstanceId in the response of ListDataMaskingInstances.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The cloud service code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two parameters must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud service type identifier. Valid values:
	//
	// - 5: RDS
	//
	// - 8: POLARDB
	//
	// - 27: PolarDBX2
	//
	// If this parameter is specified together with ProductCode, the values must match.
	//
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
