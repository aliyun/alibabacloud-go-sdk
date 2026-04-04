// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserResourceMeasureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBelongId(v string) *CheckUserResourceMeasureRequest
	GetBelongId() *string
	SetBelongIdType(v string) *CheckUserResourceMeasureRequest
	GetBelongIdType() *string
	SetBizType(v string) *CheckUserResourceMeasureRequest
	GetBizType() *string
	SetEspBizId(v string) *CheckUserResourceMeasureRequest
	GetEspBizId() *string
	SetOrderComponentParams(v string) *CheckUserResourceMeasureRequest
	GetOrderComponentParams() *string
	SetResourceCode(v string) *CheckUserResourceMeasureRequest
	GetResourceCode() *string
	SetResourceValue(v int64) *CheckUserResourceMeasureRequest
	GetResourceValue() *int64
}

type CheckUserResourceMeasureRequest struct {
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

func (s CheckUserResourceMeasureRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUserResourceMeasureRequest) GoString() string {
	return s.String()
}

func (s *CheckUserResourceMeasureRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *CheckUserResourceMeasureRequest) GetBelongIdType() *string {
	return s.BelongIdType
}

func (s *CheckUserResourceMeasureRequest) GetBizType() *string {
	return s.BizType
}

func (s *CheckUserResourceMeasureRequest) GetEspBizId() *string {
	return s.EspBizId
}

func (s *CheckUserResourceMeasureRequest) GetOrderComponentParams() *string {
	return s.OrderComponentParams
}

func (s *CheckUserResourceMeasureRequest) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *CheckUserResourceMeasureRequest) GetResourceValue() *int64 {
	return s.ResourceValue
}

func (s *CheckUserResourceMeasureRequest) SetBelongId(v string) *CheckUserResourceMeasureRequest {
	s.BelongId = &v
	return s
}

func (s *CheckUserResourceMeasureRequest) SetBelongIdType(v string) *CheckUserResourceMeasureRequest {
	s.BelongIdType = &v
	return s
}

func (s *CheckUserResourceMeasureRequest) SetBizType(v string) *CheckUserResourceMeasureRequest {
	s.BizType = &v
	return s
}

func (s *CheckUserResourceMeasureRequest) SetEspBizId(v string) *CheckUserResourceMeasureRequest {
	s.EspBizId = &v
	return s
}

func (s *CheckUserResourceMeasureRequest) SetOrderComponentParams(v string) *CheckUserResourceMeasureRequest {
	s.OrderComponentParams = &v
	return s
}

func (s *CheckUserResourceMeasureRequest) SetResourceCode(v string) *CheckUserResourceMeasureRequest {
	s.ResourceCode = &v
	return s
}

func (s *CheckUserResourceMeasureRequest) SetResourceValue(v int64) *CheckUserResourceMeasureRequest {
	s.ResourceValue = &v
	return s
}

func (s *CheckUserResourceMeasureRequest) Validate() error {
	return dara.Validate(s)
}
