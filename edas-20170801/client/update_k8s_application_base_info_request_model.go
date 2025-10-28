// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sApplicationBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateK8sApplicationBaseInfoRequest
	GetAppId() *string
	SetDescription(v string) *UpdateK8sApplicationBaseInfoRequest
	GetDescription() *string
	SetEmail(v string) *UpdateK8sApplicationBaseInfoRequest
	GetEmail() *string
	SetOwner(v string) *UpdateK8sApplicationBaseInfoRequest
	GetOwner() *string
	SetPhoneNumber(v string) *UpdateK8sApplicationBaseInfoRequest
	GetPhoneNumber() *string
}

type UpdateK8sApplicationBaseInfoRequest struct {
	// The ID of the application that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4f038ddf-b27b-****-****-88e44375****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The description of the application. The description can be up to 256 characters in length.
	//
	// example:
	//
	// app for pre-production
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The email address of the application owner.
	//
	// example:
	//
	// mymail@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The owner of the application. The value can be up to 128 characters in length.
	//
	// example:
	//
	// Tom
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The phone number of the application owner.
	//
	// example:
	//
	// 1361234xxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s UpdateK8sApplicationBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sApplicationBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationBaseInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateK8sApplicationBaseInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateK8sApplicationBaseInfoRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateK8sApplicationBaseInfoRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateK8sApplicationBaseInfoRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateK8sApplicationBaseInfoRequest) SetAppId(v string) *UpdateK8sApplicationBaseInfoRequest {
	s.AppId = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoRequest) SetDescription(v string) *UpdateK8sApplicationBaseInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoRequest) SetEmail(v string) *UpdateK8sApplicationBaseInfoRequest {
	s.Email = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoRequest) SetOwner(v string) *UpdateK8sApplicationBaseInfoRequest {
	s.Owner = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoRequest) SetPhoneNumber(v string) *UpdateK8sApplicationBaseInfoRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
