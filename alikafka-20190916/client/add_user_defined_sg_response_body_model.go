// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserDefinedSgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddUserDefinedSgResponseBody
	GetCode() *int32
	SetMessage(v string) *AddUserDefinedSgResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddUserDefinedSgResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddUserDefinedSgResponseBody
	GetSuccess() *bool
}

type AddUserDefinedSgResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserDefinedSgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserDefinedSgResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserDefinedSgResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddUserDefinedSgResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddUserDefinedSgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserDefinedSgResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserDefinedSgResponseBody) SetCode(v int32) *AddUserDefinedSgResponseBody {
	s.Code = &v
	return s
}

func (s *AddUserDefinedSgResponseBody) SetMessage(v string) *AddUserDefinedSgResponseBody {
	s.Message = &v
	return s
}

func (s *AddUserDefinedSgResponseBody) SetRequestId(v string) *AddUserDefinedSgResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserDefinedSgResponseBody) SetSuccess(v bool) *AddUserDefinedSgResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserDefinedSgResponseBody) Validate() error {
	return dara.Validate(s)
}
