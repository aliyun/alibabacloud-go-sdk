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
	// Interactive Messages application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Terminal device ID, uniquely representing a user terminal device, user-defined. It consists of lowercase letters, numbers, underscores (_), and hyphens (-), with a maximum length of 64 characters. Different terminal devices need to use different DeviceIds. We recommend obtaining it from the terminal device and passing it to the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Terminal device type. Valid values:
	//
	// - ios
	//
	// - android
	//
	// - web
	//
	// - pc
	//
	// This parameter is required.
	//
	// example:
	//
	// android
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// User UserId, user-defined, unique within the AppId. It consists of lowercase letters, numbers, underscores (_), and periods (.), with a maximum length of 32 characters. Different users need to use different UserIds.
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
