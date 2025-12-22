// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserDefinedSgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyUserDefinedSgResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyUserDefinedSgResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyUserDefinedSgResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyUserDefinedSgResponseBody
	GetSuccess() *bool
}

type ModifyUserDefinedSgResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyUserDefinedSgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserDefinedSgResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserDefinedSgResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyUserDefinedSgResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyUserDefinedSgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserDefinedSgResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyUserDefinedSgResponseBody) SetCode(v int32) *ModifyUserDefinedSgResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyUserDefinedSgResponseBody) SetMessage(v string) *ModifyUserDefinedSgResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyUserDefinedSgResponseBody) SetRequestId(v string) *ModifyUserDefinedSgResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserDefinedSgResponseBody) SetSuccess(v bool) *ModifyUserDefinedSgResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyUserDefinedSgResponseBody) Validate() error {
	return dara.Validate(s)
}
