// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtmpKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRtmpKeyResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdateRtmpKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateRtmpKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRtmpKeyResponseBody
	GetSuccess() *bool
}

type UpdateRtmpKeyResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRtmpKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtmpKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRtmpKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRtmpKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateRtmpKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRtmpKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRtmpKeyResponseBody) SetCode(v string) *UpdateRtmpKeyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRtmpKeyResponseBody) SetErrorMessage(v string) *UpdateRtmpKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRtmpKeyResponseBody) SetRequestId(v string) *UpdateRtmpKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRtmpKeyResponseBody) SetSuccess(v bool) *UpdateRtmpKeyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRtmpKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
