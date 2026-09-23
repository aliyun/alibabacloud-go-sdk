// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCreditSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddCreditSeatsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddCreditSeatsResponseBody
	GetCode() *string
	SetData(v *AddCreditSeatsResponseBodyData) *AddCreditSeatsResponseBody
	GetData() *AddCreditSeatsResponseBodyData
	SetMessage(v string) *AddCreditSeatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCreditSeatsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddCreditSeatsResponseBody
	GetSuccess() *bool
}

type AddCreditSeatsResponseBody struct {
	AccessDeniedDetail *string                         `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *AddCreditSeatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCreditSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCreditSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *AddCreditSeatsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddCreditSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddCreditSeatsResponseBody) GetData() *AddCreditSeatsResponseBodyData {
	return s.Data
}

func (s *AddCreditSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCreditSeatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCreditSeatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddCreditSeatsResponseBody) SetAccessDeniedDetail(v string) *AddCreditSeatsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddCreditSeatsResponseBody) SetCode(v string) *AddCreditSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *AddCreditSeatsResponseBody) SetData(v *AddCreditSeatsResponseBodyData) *AddCreditSeatsResponseBody {
	s.Data = v
	return s
}

func (s *AddCreditSeatsResponseBody) SetMessage(v string) *AddCreditSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *AddCreditSeatsResponseBody) SetRequestId(v string) *AddCreditSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCreditSeatsResponseBody) SetSuccess(v bool) *AddCreditSeatsResponseBody {
	s.Success = &v
	return s
}

func (s *AddCreditSeatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddCreditSeatsResponseBodyData struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OrderId    *int64    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s AddCreditSeatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddCreditSeatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddCreditSeatsResponseBodyData) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *AddCreditSeatsResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *AddCreditSeatsResponseBodyData) SetInstanceId(v []*string) *AddCreditSeatsResponseBodyData {
	s.InstanceId = v
	return s
}

func (s *AddCreditSeatsResponseBodyData) SetOrderId(v int64) *AddCreditSeatsResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *AddCreditSeatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
