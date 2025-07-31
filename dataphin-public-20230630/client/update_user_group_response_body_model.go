// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserGroupResponseBody
	GetCode() *string
	SetData(v bool) *UpdateUserGroupResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateUserGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateUserGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUserGroupResponseBody
	GetSuccess() *bool
}

type UpdateUserGroupResponseBody struct {
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

func (s UpdateUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateUserGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateUserGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserGroupResponseBody) SetCode(v string) *UpdateUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetData(v bool) *UpdateUserGroupResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetHttpStatusCode(v int32) *UpdateUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetMessage(v string) *UpdateUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetSuccess(v bool) *UpdateUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
