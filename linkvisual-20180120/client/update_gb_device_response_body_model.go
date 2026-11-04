// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGbDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGbDeviceResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdateGbDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateGbDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGbDeviceResponseBody
	GetSuccess() *bool
}

type UpdateGbDeviceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGbDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGbDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGbDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGbDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateGbDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGbDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGbDeviceResponseBody) SetCode(v string) *UpdateGbDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGbDeviceResponseBody) SetErrorMessage(v string) *UpdateGbDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateGbDeviceResponseBody) SetRequestId(v string) *UpdateGbDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGbDeviceResponseBody) SetSuccess(v bool) *UpdateGbDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGbDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
