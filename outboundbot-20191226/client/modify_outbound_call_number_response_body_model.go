// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOutboundCallNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyOutboundCallNumberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyOutboundCallNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyOutboundCallNumberResponseBody
	GetMessage() *string
	SetOutboundCallNumber(v *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) *ModifyOutboundCallNumberResponseBody
	GetOutboundCallNumber() *ModifyOutboundCallNumberResponseBodyOutboundCallNumber
	SetRequestId(v string) *ModifyOutboundCallNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyOutboundCallNumberResponseBody
	GetSuccess() *bool
}

type ModifyOutboundCallNumberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	OutboundCallNumber *ModifyOutboundCallNumberResponseBodyOutboundCallNumber `json:"OutboundCallNumber,omitempty" xml:"OutboundCallNumber,omitempty" type:"Struct"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyOutboundCallNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOutboundCallNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOutboundCallNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyOutboundCallNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyOutboundCallNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyOutboundCallNumberResponseBody) GetOutboundCallNumber() *ModifyOutboundCallNumberResponseBodyOutboundCallNumber {
	return s.OutboundCallNumber
}

func (s *ModifyOutboundCallNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOutboundCallNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyOutboundCallNumberResponseBody) SetCode(v string) *ModifyOutboundCallNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBody) SetHttpStatusCode(v int32) *ModifyOutboundCallNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBody) SetMessage(v string) *ModifyOutboundCallNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBody) SetOutboundCallNumber(v *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) *ModifyOutboundCallNumberResponseBody {
	s.OutboundCallNumber = v
	return s
}

func (s *ModifyOutboundCallNumberResponseBody) SetRequestId(v string) *ModifyOutboundCallNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBody) SetSuccess(v bool) *ModifyOutboundCallNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBody) Validate() error {
	if s.OutboundCallNumber != nil {
		if err := s.OutboundCallNumber.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyOutboundCallNumberResponseBodyOutboundCallNumber struct {
	// example:
	//
	// 10088
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// fa0e21e9-caab-4629-9121-1e341243d599
	OutboundCallNumberId *string `json:"OutboundCallNumberId,omitempty" xml:"OutboundCallNumberId,omitempty"`
	// example:
	//
	// 10
	RateLimitCount *string `json:"RateLimitCount,omitempty" xml:"RateLimitCount,omitempty"`
	// example:
	//
	// 100
	RateLimitPeriod *string `json:"RateLimitPeriod,omitempty" xml:"RateLimitPeriod,omitempty"`
}

func (s ModifyOutboundCallNumberResponseBodyOutboundCallNumber) String() string {
	return dara.Prettify(s)
}

func (s ModifyOutboundCallNumberResponseBodyOutboundCallNumber) GoString() string {
	return s.String()
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) GetNumber() *string {
	return s.Number
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) GetOutboundCallNumberId() *string {
	return s.OutboundCallNumberId
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) GetRateLimitCount() *string {
	return s.RateLimitCount
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) GetRateLimitPeriod() *string {
	return s.RateLimitPeriod
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) SetNumber(v string) *ModifyOutboundCallNumberResponseBodyOutboundCallNumber {
	s.Number = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) SetOutboundCallNumberId(v string) *ModifyOutboundCallNumberResponseBodyOutboundCallNumber {
	s.OutboundCallNumberId = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) SetRateLimitCount(v string) *ModifyOutboundCallNumberResponseBodyOutboundCallNumber {
	s.RateLimitCount = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) SetRateLimitPeriod(v string) *ModifyOutboundCallNumberResponseBodyOutboundCallNumber {
	s.RateLimitPeriod = &v
	return s
}

func (s *ModifyOutboundCallNumberResponseBodyOutboundCallNumber) Validate() error {
	return dara.Validate(s)
}
