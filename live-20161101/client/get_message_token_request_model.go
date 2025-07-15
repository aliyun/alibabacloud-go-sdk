// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMessageTokenRequest
	GetAppId() *string
	SetDeviceId(v string) *GetMessageTokenRequest
	GetDeviceId() *string
	SetDeviceType(v string) *GetMessageTokenRequest
	GetDeviceType() *string
	SetUserId(v string) *GetMessageTokenRequest
	GetUserId() *string
}

type GetMessageTokenRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the device. Each device has a unique ID. You can specify a custom ID. The ID can be up to 64 characters in length and can contain lowercase letters, digits, underscores (_), and hyphen (-). You can specify multiple device IDs. We recommend that you obtain the IDs from the devices and pass the IDs to the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The type of the device. Valid values:
	//
	// 	- ios
	//
	// 	- android
	//
	// 	- web
	//
	// 	- pc
	//
	// This parameter is required.
	//
	// example:
	//
	// android
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The ID of the user. Each user has a unique ID in the application. The ID can be up to 32 characters in length and can contain lowercase letters, digits, underscores (_), and periods (.). You can specify multiple user IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMessageTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageTokenRequest) GoString() string {
	return s.String()
}

func (s *GetMessageTokenRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageTokenRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetMessageTokenRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *GetMessageTokenRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetMessageTokenRequest) SetAppId(v string) *GetMessageTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageTokenRequest) SetDeviceId(v string) *GetMessageTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *GetMessageTokenRequest) SetDeviceType(v string) *GetMessageTokenRequest {
	s.DeviceType = &v
	return s
}

func (s *GetMessageTokenRequest) SetUserId(v string) *GetMessageTokenRequest {
	s.UserId = &v
	return s
}

func (s *GetMessageTokenRequest) Validate() error {
	return dara.Validate(s)
}
