// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCreditSeatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RenewCreditSeatResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *RenewCreditSeatResponseBody
	GetCode() *string
	SetData(v *RenewCreditSeatResponseBodyData) *RenewCreditSeatResponseBody
	GetData() *RenewCreditSeatResponseBodyData
	SetMessage(v string) *RenewCreditSeatResponseBody
	GetMessage() *string
	SetRequestId(v string) *RenewCreditSeatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewCreditSeatResponseBody
	GetSuccess() *bool
}

type RenewCreditSeatResponseBody struct {
	AccessDeniedDetail *string                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *RenewCreditSeatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewCreditSeatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewCreditSeatResponseBody) GoString() string {
	return s.String()
}

func (s *RenewCreditSeatResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RenewCreditSeatResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewCreditSeatResponseBody) GetData() *RenewCreditSeatResponseBodyData {
	return s.Data
}

func (s *RenewCreditSeatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenewCreditSeatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewCreditSeatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewCreditSeatResponseBody) SetAccessDeniedDetail(v string) *RenewCreditSeatResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RenewCreditSeatResponseBody) SetCode(v string) *RenewCreditSeatResponseBody {
	s.Code = &v
	return s
}

func (s *RenewCreditSeatResponseBody) SetData(v *RenewCreditSeatResponseBodyData) *RenewCreditSeatResponseBody {
	s.Data = v
	return s
}

func (s *RenewCreditSeatResponseBody) SetMessage(v string) *RenewCreditSeatResponseBody {
	s.Message = &v
	return s
}

func (s *RenewCreditSeatResponseBody) SetRequestId(v string) *RenewCreditSeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewCreditSeatResponseBody) SetSuccess(v bool) *RenewCreditSeatResponseBody {
	s.Success = &v
	return s
}

func (s *RenewCreditSeatResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewCreditSeatResponseBodyData struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewCreditSeatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewCreditSeatResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewCreditSeatResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewCreditSeatResponseBodyData) SetOrderId(v int64) *RenewCreditSeatResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewCreditSeatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
