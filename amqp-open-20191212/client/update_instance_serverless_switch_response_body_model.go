// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceServerlessSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateInstanceServerlessSwitchResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateInstanceServerlessSwitchResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateInstanceServerlessSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceServerlessSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceServerlessSwitchResponseBody
	GetSuccess() *bool
}

type UpdateInstanceServerlessSwitchResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CCBB1225-C392-480E-8C7F-D09AB2CD2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceServerlessSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceServerlessSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceServerlessSwitchResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateInstanceServerlessSwitchResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstanceServerlessSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceServerlessSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceServerlessSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceServerlessSwitchResponseBody) SetCode(v int32) *UpdateInstanceServerlessSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponseBody) SetData(v bool) *UpdateInstanceServerlessSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponseBody) SetMessage(v string) *UpdateInstanceServerlessSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponseBody) SetRequestId(v string) *UpdateInstanceServerlessSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponseBody) SetSuccess(v bool) *UpdateInstanceServerlessSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
