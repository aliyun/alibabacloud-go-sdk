// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtmpDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRtmpDeviceResponseBody
	GetCode() *string
	SetData(v *CreateRtmpDeviceResponseBodyData) *CreateRtmpDeviceResponseBody
	GetData() *CreateRtmpDeviceResponseBodyData
	SetErrorMessage(v string) *CreateRtmpDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateRtmpDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRtmpDeviceResponseBody
	GetSuccess() *bool
}

type CreateRtmpDeviceResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateRtmpDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRtmpDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRtmpDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRtmpDeviceResponseBody) GetData() *CreateRtmpDeviceResponseBodyData {
	return s.Data
}

func (s *CreateRtmpDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateRtmpDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRtmpDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRtmpDeviceResponseBody) SetCode(v string) *CreateRtmpDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetData(v *CreateRtmpDeviceResponseBodyData) *CreateRtmpDeviceResponseBody {
	s.Data = v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetErrorMessage(v string) *CreateRtmpDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetRequestId(v string) *CreateRtmpDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRtmpDeviceResponseBody) SetSuccess(v bool) *CreateRtmpDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRtmpDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRtmpDeviceResponseBodyData struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s CreateRtmpDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateRtmpDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CreateRtmpDeviceResponseBodyData) GetIotId() *string {
	return s.IotId
}

func (s *CreateRtmpDeviceResponseBodyData) GetStreamName() *string {
	return s.StreamName
}

func (s *CreateRtmpDeviceResponseBodyData) SetDeviceName(v string) *CreateRtmpDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *CreateRtmpDeviceResponseBodyData) SetIotId(v string) *CreateRtmpDeviceResponseBodyData {
	s.IotId = &v
	return s
}

func (s *CreateRtmpDeviceResponseBodyData) SetStreamName(v string) *CreateRtmpDeviceResponseBodyData {
	s.StreamName = &v
	return s
}

func (s *CreateRtmpDeviceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
