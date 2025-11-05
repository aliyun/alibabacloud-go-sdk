// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCouponTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCouponTemplateResponseBody
	GetCode() *string
	SetData(v bool) *DeleteCouponTemplateResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteCouponTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCouponTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCouponTemplateResponseBody
	GetSuccess() *bool
}

type DeleteCouponTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A747A00F-E096-5244-88B3-3E474BAE3AE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCouponTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCouponTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCouponTemplateResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteCouponTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCouponTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCouponTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCouponTemplateResponseBody) SetCode(v string) *DeleteCouponTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCouponTemplateResponseBody) SetData(v bool) *DeleteCouponTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCouponTemplateResponseBody) SetMessage(v string) *DeleteCouponTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCouponTemplateResponseBody) SetRequestId(v string) *DeleteCouponTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCouponTemplateResponseBody) SetSuccess(v bool) *DeleteCouponTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCouponTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
