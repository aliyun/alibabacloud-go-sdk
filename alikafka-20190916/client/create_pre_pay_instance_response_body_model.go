// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePrePayInstanceResponseBody
	GetCode() *int32
	SetData(v *CreatePrePayInstanceResponseBodyData) *CreatePrePayInstanceResponseBody
	GetData() *CreatePrePayInstanceResponseBodyData
	SetMessage(v string) *CreatePrePayInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePrePayInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePrePayInstanceResponseBody
	GetSuccess() *bool
}

type CreatePrePayInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePrePayInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E57A8862-DF68-4055-8E55-B80CB4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePrePayInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePrePayInstanceResponseBody) GetData() *CreatePrePayInstanceResponseBodyData {
	return s.Data
}

func (s *CreatePrePayInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePrePayInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrePayInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePrePayInstanceResponseBody) SetCode(v int32) *CreatePrePayInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetData(v *CreatePrePayInstanceResponseBodyData) *CreatePrePayInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetMessage(v string) *CreatePrePayInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetRequestId(v string) *CreatePrePayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetSuccess(v bool) *CreatePrePayInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePrePayInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePrePayInstanceResponseBodyData struct {
	// example:
	//
	// alikafka_post-cn-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 236972661xxxx
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreatePrePayInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreatePrePayInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreatePrePayInstanceResponseBodyData) SetInstanceId(v string) *CreatePrePayInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreatePrePayInstanceResponseBodyData) SetOrderId(v int64) *CreatePrePayInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreatePrePayInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
