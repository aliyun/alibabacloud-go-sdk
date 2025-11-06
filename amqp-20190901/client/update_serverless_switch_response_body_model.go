// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerlessSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateServerlessSwitchResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateServerlessSwitchResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateServerlessSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServerlessSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateServerlessSwitchResponseBody
	GetSuccess() *bool
}

type UpdateServerlessSwitchResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateServerlessSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerlessSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerlessSwitchResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateServerlessSwitchResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateServerlessSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServerlessSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServerlessSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateServerlessSwitchResponseBody) SetCode(v int32) *UpdateServerlessSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServerlessSwitchResponseBody) SetData(v bool) *UpdateServerlessSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateServerlessSwitchResponseBody) SetMessage(v string) *UpdateServerlessSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServerlessSwitchResponseBody) SetRequestId(v string) *UpdateServerlessSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerlessSwitchResponseBody) SetSuccess(v bool) *UpdateServerlessSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateServerlessSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
