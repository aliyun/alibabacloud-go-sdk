// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserGroupSwitchResponseBody
	GetCode() *string
	SetData(v bool) *UpdateUserGroupSwitchResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateUserGroupSwitchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateUserGroupSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserGroupSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUserGroupSwitchResponseBody
	GetSuccess() *bool
}

type UpdateUserGroupSwitchResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserGroupSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserGroupSwitchResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateUserGroupSwitchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateUserGroupSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserGroupSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserGroupSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserGroupSwitchResponseBody) SetCode(v string) *UpdateUserGroupSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetData(v bool) *UpdateUserGroupSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetHttpStatusCode(v int32) *UpdateUserGroupSwitchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetMessage(v string) *UpdateUserGroupSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetRequestId(v string) *UpdateUserGroupSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetSuccess(v bool) *UpdateUserGroupSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
