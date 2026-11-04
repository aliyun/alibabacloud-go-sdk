// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtmpDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRtmpDeviceResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteRtmpDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteRtmpDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRtmpDeviceResponseBody
	GetSuccess() *bool
}

type DeleteRtmpDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRtmpDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtmpDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRtmpDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRtmpDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRtmpDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRtmpDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRtmpDeviceResponseBody) SetCode(v string) *DeleteRtmpDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRtmpDeviceResponseBody) SetErrorMessage(v string) *DeleteRtmpDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRtmpDeviceResponseBody) SetRequestId(v string) *DeleteRtmpDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRtmpDeviceResponseBody) SetSuccess(v bool) *DeleteRtmpDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRtmpDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
