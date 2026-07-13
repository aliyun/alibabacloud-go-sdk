// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceResponseBody
	GetCode() *string
	SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody
	GetData() *CreateInstanceResponseBodyData
	SetHttpStatusCode(v int32) *CreateInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *CreateInstanceResponseBody
	GetInstanceId() *string
	SetMessage(v string) *CreateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceResponseBody
	GetSuccess() *bool
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// InstanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// req-create-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceResponseBody) GetData() *CreateInstanceResponseBodyData {
	return s.Data
}

func (s *CreateInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v int32) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceResponseBodyData struct {
	// example:
	//
	// agentteams-test-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// open-create-agentteams-test-001-req-create-001
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateInstanceResponseBodyData) SetInstanceId(v string) *CreateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetOrderId(v string) *CreateInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
