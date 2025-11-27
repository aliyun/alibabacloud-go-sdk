// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePostPayInstanceResponseBody
	GetCode() *int32
	SetData(v *CreatePostPayInstanceResponseBodyData) *CreatePostPayInstanceResponseBody
	GetData() *CreatePostPayInstanceResponseBodyData
	SetMessage(v string) *CreatePostPayInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePostPayInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePostPayInstanceResponseBody
	GetSuccess() *bool
}

type CreatePostPayInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePostPayInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePostPayInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePostPayInstanceResponseBody) GetData() *CreatePostPayInstanceResponseBodyData {
	return s.Data
}

func (s *CreatePostPayInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePostPayInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePostPayInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePostPayInstanceResponseBody) SetCode(v int32) *CreatePostPayInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetData(v *CreatePostPayInstanceResponseBodyData) *CreatePostPayInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetMessage(v string) *CreatePostPayInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetRequestId(v string) *CreatePostPayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetSuccess(v bool) *CreatePostPayInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePostPayInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePostPayInstanceResponseBodyData struct {
	// example:
	//
	// alikafka_pre-cn-pe333xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 236972661580636
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreatePostPayInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreatePostPayInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreatePostPayInstanceResponseBodyData) SetInstanceId(v string) *CreatePostPayInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreatePostPayInstanceResponseBodyData) SetOrderId(v int64) *CreatePostPayInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreatePostPayInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
