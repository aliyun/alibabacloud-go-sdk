// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtmpKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRtmpKeyResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteRtmpKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteRtmpKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRtmpKeyResponseBody
	GetSuccess() *bool
}

type DeleteRtmpKeyResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRtmpKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtmpKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRtmpKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRtmpKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRtmpKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRtmpKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRtmpKeyResponseBody) SetCode(v string) *DeleteRtmpKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRtmpKeyResponseBody) SetErrorMessage(v string) *DeleteRtmpKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRtmpKeyResponseBody) SetRequestId(v string) *DeleteRtmpKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRtmpKeyResponseBody) SetSuccess(v bool) *DeleteRtmpKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRtmpKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
