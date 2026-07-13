// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceResponseBody
	GetCode() *string
	SetData(v *UpdateInstanceResponseBodyData) *UpdateInstanceResponseBody
	GetData() *UpdateInstanceResponseBodyData
	SetHttpStatusCode(v int32) *UpdateInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceResponseBody
	GetSuccess() *bool
}

type UpdateInstanceResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// request-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceResponseBody) GetData() *UpdateInstanceResponseBodyData {
	return s.Data
}

func (s *UpdateInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceResponseBody) SetCode(v string) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetData(v *UpdateInstanceResponseBodyData) *UpdateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v bool) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceResponseBodyData struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceResponseBodyData) SetInstanceId(v string) *UpdateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) SetInstanceName(v string) *UpdateInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
