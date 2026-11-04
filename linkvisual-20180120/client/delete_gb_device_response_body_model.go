// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGbDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGbDeviceResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteGbDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteGbDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGbDeviceResponseBody
	GetSuccess() *bool
}

type DeleteGbDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGbDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGbDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGbDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGbDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteGbDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGbDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGbDeviceResponseBody) SetCode(v string) *DeleteGbDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGbDeviceResponseBody) SetErrorMessage(v string) *DeleteGbDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteGbDeviceResponseBody) SetRequestId(v string) *DeleteGbDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGbDeviceResponseBody) SetSuccess(v bool) *DeleteGbDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGbDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
