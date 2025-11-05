// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCouponResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelCouponResponseBody
	GetCode() *string
	SetData(v bool) *CancelCouponResponseBody
	GetData() *bool
	SetMessage(v string) *CancelCouponResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelCouponResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelCouponResponseBody
	GetSuccess() *bool
}

type CancelCouponResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data    *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelCouponResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCouponResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCouponResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelCouponResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelCouponResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelCouponResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCouponResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelCouponResponseBody) SetCode(v string) *CancelCouponResponseBody {
	s.Code = &v
	return s
}

func (s *CancelCouponResponseBody) SetData(v bool) *CancelCouponResponseBody {
	s.Data = &v
	return s
}

func (s *CancelCouponResponseBody) SetMessage(v string) *CancelCouponResponseBody {
	s.Message = &v
	return s
}

func (s *CancelCouponResponseBody) SetRequestId(v string) *CancelCouponResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCouponResponseBody) SetSuccess(v bool) *CancelCouponResponseBody {
	s.Success = &v
	return s
}

func (s *CancelCouponResponseBody) Validate() error {
	return dara.Validate(s)
}
