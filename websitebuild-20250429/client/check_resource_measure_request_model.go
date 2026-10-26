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
	SetDesiredTotal(v int64) *CheckResourceMeasureRequest
	GetDesiredTotal() *int64
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
	// The ID of the owner, such as the website ID or Alibaba Cloud user ID.
	//
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// The type of the owner ID. Valid values: siteId and uid.
	//
	// example:
	//
	// USER
	BelongIdType *string `json:"BelongIdType,omitempty" xml:"BelongIdType,omitempty"`
	// The ESP business type.
	//
	// example:
	//
	// 1
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The desired total for APP_WORKSPACE. Valid values: 0 to 100000. If this parameter is specified, it takes precedence over ResourceValue and is not added to the existing usage. For new calls, you only need to specify ResourceValue. This parameter is not used for APP or USER.
	//
	// example:
	//
	// 21
	DesiredTotal *int64 `json:"DesiredTotal,omitempty" xml:"DesiredTotal,omitempty"`
	// The ESP business ID.
	//
	// example:
	//
	// p20202933455
	EspBizId *string `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	// The order component parameters. If this parameter is empty, EspBizId and BizType cannot be empty.
	//
	// example:
	//
	// {"siteversion":"test"}
	OrderComponentParams *string `json:"OrderComponentParams,omitempty" xml:"OrderComponentParams,omitempty"`
	// The resource identifier.
	//
	// example:
	//
	// InspirationTokens
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The resource value. This parameter is empty by default and is required only for feature resources. Valid values: 0: Not supported. 1: Supported.
	//
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

func (s *CheckResourceMeasureRequest) GetDesiredTotal() *int64 {
	return s.DesiredTotal
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

func (s *CheckResourceMeasureRequest) SetDesiredTotal(v int64) *CheckResourceMeasureRequest {
	s.DesiredTotal = &v
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
