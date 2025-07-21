// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateInstanceNameResponseBody
	GetCode() *int32
	SetData(v string) *UpdateInstanceNameResponseBody
	GetData() *string
	SetMessage(v string) *UpdateInstanceNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceNameResponseBody
	GetSuccess() *bool
}

type UpdateInstanceNameResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned when an error occurs during the update of the instance name.
	//
	// example:
	//
	// InstanceNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6DC68EC9-0E76-5435-B8C0-FF9492B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned message that indicates the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateInstanceNameResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateInstanceNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceNameResponseBody) SetCode(v int32) *UpdateInstanceNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetData(v string) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetMessage(v string) *UpdateInstanceNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
