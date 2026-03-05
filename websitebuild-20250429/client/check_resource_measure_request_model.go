// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourceMeasureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBelongId(v string) *CheckResourceMeasureRequest
	GetBelongId() *string
	SetBelongIdType(v string) *CheckResourceMeasureRequest
	GetBelongIdType() *string
	SetBizType(v string) *CheckResourceMeasureRequest
	GetBizType() *string
	SetEspBizId(v string) *CheckResourceMeasureRequest
	GetEspBizId() *string
	SetOrderComponentParams(v string) *CheckResourceMeasureRequest
	GetOrderComponentParams() *string
	SetResourceCode(v string) *CheckResourceMeasureRequest
	GetResourceCode() *string
	SetResourceValue(v int64) *CheckResourceMeasureRequest
	GetResourceValue() *int64
}

type CheckResourceMeasureRequest struct {
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// example:
	//
	// USER
	BelongIdType *string `json:"BelongIdType,omitempty" xml:"BelongIdType,omitempty"`
	// example:
	//
	// 1
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// p20202933455
	EspBizId *string `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	// example:
	//
	// {"siteversion":"test"}
	OrderComponentParams *string `json:"OrderComponentParams,omitempty" xml:"OrderComponentParams,omitempty"`
	// example:
	//
	// InspirationTokens
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// example:
	//
	// 1
	ResourceValue *int64 `json:"ResourceValue,omitempty" xml:"ResourceValue,omitempty"`
}

func (s CheckResourceMeasureRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceMeasureRequest) GoString() string {
	return s.String()
}

func (s *CheckResourceMeasureRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *CheckResourceMeasureRequest) GetBelongIdType() *string {
	return s.BelongIdType
}

func (s *CheckResourceMeasureRequest) GetBizType() *string {
	return s.BizType
}

func (s *CheckResourceMeasureRequest) GetEspBizId() *string {
	return s.EspBizId
}

func (s *CheckResourceMeasureRequest) GetOrderComponentParams() *string {
	return s.OrderComponentParams
}

func (s *CheckResourceMeasureRequest) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *CheckResourceMeasureRequest) GetResourceValue() *int64 {
	return s.ResourceValue
}

func (s *CheckResourceMeasureRequest) SetBelongId(v string) *CheckResourceMeasureRequest {
	s.BelongId = &v
	return s
}

func (s *CheckResourceMeasureRequest) SetBelongIdType(v string) *CheckResourceMeasureRequest {
	s.BelongIdType = &v
	return s
}

func (s *CheckResourceMeasureRequest) SetBizType(v string) *CheckResourceMeasureRequest {
	s.BizType = &v
	return s
}

func (s *CheckResourceMeasureRequest) SetEspBizId(v string) *CheckResourceMeasureRequest {
	s.EspBizId = &v
	return s
}

func (s *CheckResourceMeasureRequest) SetOrderComponentParams(v string) *CheckResourceMeasureRequest {
	s.OrderComponentParams = &v
	return s
}

func (s *CheckResourceMeasureRequest) SetResourceCode(v string) *CheckResourceMeasureRequest {
	s.ResourceCode = &v
	return s
}

func (s *CheckResourceMeasureRequest) SetResourceValue(v int64) *CheckResourceMeasureRequest {
	s.ResourceValue = &v
	return s
}

func (s *CheckResourceMeasureRequest) Validate() error {
	return dara.Validate(s)
}
