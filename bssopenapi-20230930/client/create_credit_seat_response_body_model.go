// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCreditSeatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateCreditSeatResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateCreditSeatResponseBody
	GetCode() *string
	SetData(v *CreateCreditSeatResponseBodyData) *CreateCreditSeatResponseBody
	GetData() *CreateCreditSeatResponseBodyData
	SetMessage(v string) *CreateCreditSeatResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCreditSeatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCreditSeatResponseBody
	GetSuccess() *bool
}

type CreateCreditSeatResponseBody struct {
	AccessDeniedDetail *string                           `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *CreateCreditSeatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCreditSeatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditSeatResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCreditSeatResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateCreditSeatResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCreditSeatResponseBody) GetData() *CreateCreditSeatResponseBodyData {
	return s.Data
}

func (s *CreateCreditSeatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCreditSeatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCreditSeatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCreditSeatResponseBody) SetAccessDeniedDetail(v string) *CreateCreditSeatResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateCreditSeatResponseBody) SetCode(v string) *CreateCreditSeatResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCreditSeatResponseBody) SetData(v *CreateCreditSeatResponseBodyData) *CreateCreditSeatResponseBody {
	s.Data = v
	return s
}

func (s *CreateCreditSeatResponseBody) SetMessage(v string) *CreateCreditSeatResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCreditSeatResponseBody) SetRequestId(v string) *CreateCreditSeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCreditSeatResponseBody) SetSuccess(v bool) *CreateCreditSeatResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCreditSeatResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCreditSeatResponseBodyData struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OrderId    *int64    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateCreditSeatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditSeatResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCreditSeatResponseBodyData) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *CreateCreditSeatResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateCreditSeatResponseBodyData) SetInstanceId(v []*string) *CreateCreditSeatResponseBodyData {
	s.InstanceId = v
	return s
}

func (s *CreateCreditSeatResponseBodyData) SetOrderId(v int64) *CreateCreditSeatResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateCreditSeatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
