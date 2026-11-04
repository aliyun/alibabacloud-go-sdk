// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGbDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGbDeviceResponseBody
	GetCode() *string
	SetData(v *CreateGbDeviceResponseBodyData) *CreateGbDeviceResponseBody
	GetData() *CreateGbDeviceResponseBodyData
	SetErrorMessage(v string) *CreateGbDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateGbDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGbDeviceResponseBody
	GetSuccess() *bool
}

type CreateGbDeviceResponseBody struct {
	Code         *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateGbDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGbDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGbDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGbDeviceResponseBody) GetData() *CreateGbDeviceResponseBodyData {
	return s.Data
}

func (s *CreateGbDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateGbDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGbDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGbDeviceResponseBody) SetCode(v string) *CreateGbDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGbDeviceResponseBody) SetData(v *CreateGbDeviceResponseBodyData) *CreateGbDeviceResponseBody {
	s.Data = v
	return s
}

func (s *CreateGbDeviceResponseBody) SetErrorMessage(v string) *CreateGbDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateGbDeviceResponseBody) SetRequestId(v string) *CreateGbDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGbDeviceResponseBody) SetSuccess(v bool) *CreateGbDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGbDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGbDeviceResponseBodyData struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s CreateGbDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateGbDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGbDeviceResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateGbDeviceResponseBodyData) GetIotId() *string {
	return s.IotId
}

func (s *CreateGbDeviceResponseBodyData) SetDeviceName(v string) *CreateGbDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *CreateGbDeviceResponseBodyData) SetIotId(v string) *CreateGbDeviceResponseBodyData {
	s.IotId = &v
	return s
}

func (s *CreateGbDeviceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
