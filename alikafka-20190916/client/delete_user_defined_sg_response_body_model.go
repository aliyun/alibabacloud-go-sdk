// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefinedSgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteUserDefinedSgResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteUserDefinedSgResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUserDefinedSgResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUserDefinedSgResponseBody
	GetSuccess() *bool
}

type DeleteUserDefinedSgResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserDefinedSgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefinedSgResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDefinedSgResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteUserDefinedSgResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUserDefinedSgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserDefinedSgResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserDefinedSgResponseBody) SetCode(v int32) *DeleteUserDefinedSgResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserDefinedSgResponseBody) SetMessage(v string) *DeleteUserDefinedSgResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserDefinedSgResponseBody) SetRequestId(v string) *DeleteUserDefinedSgResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserDefinedSgResponseBody) SetSuccess(v bool) *DeleteUserDefinedSgResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserDefinedSgResponseBody) Validate() error {
	return dara.Validate(s)
}
