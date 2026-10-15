// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlansInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSavingsPlansInstanceResponseBody
	GetCode() *string
	SetData(v *CreateSavingsPlansInstanceResponseBodyData) *CreateSavingsPlansInstanceResponseBody
	GetData() *CreateSavingsPlansInstanceResponseBodyData
	SetMessage(v string) *CreateSavingsPlansInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSavingsPlansInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSavingsPlansInstanceResponseBody
	GetSuccess() *bool
}

type CreateSavingsPlansInstanceResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateSavingsPlansInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSavingsPlansInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlansInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlansInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSavingsPlansInstanceResponseBody) GetData() *CreateSavingsPlansInstanceResponseBodyData {
	return s.Data
}

func (s *CreateSavingsPlansInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSavingsPlansInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSavingsPlansInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSavingsPlansInstanceResponseBody) SetCode(v string) *CreateSavingsPlansInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetData(v *CreateSavingsPlansInstanceResponseBodyData) *CreateSavingsPlansInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetMessage(v string) *CreateSavingsPlansInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetRequestId(v string) *CreateSavingsPlansInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) SetSuccess(v bool) *CreateSavingsPlansInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSavingsPlansInstanceResponseBodyData struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateSavingsPlansInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlansInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlansInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateSavingsPlansInstanceResponseBodyData) SetOrderId(v int64) *CreateSavingsPlansInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
