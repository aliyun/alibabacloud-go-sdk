// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateChatSessionRequest
	GetInstanceId() *string
	SetLicense(v string) *CreateChatSessionRequest
	GetLicense() *string
	SetPlatform(v string) *CreateChatSessionRequest
	GetPlatform() *string
}

type CreateChatSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rmq-cn-9t946y43m1d
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// b9be4b25c2d38c409c376ffd2372be1
	License *string `json:"license,omitempty" xml:"license,omitempty"`
	// example:
	//
	// Web | Android | iOS
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
}

func (s CreateChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateChatSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateChatSessionRequest) GetLicense() *string {
	return s.License
}

func (s *CreateChatSessionRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateChatSessionRequest) SetInstanceId(v string) *CreateChatSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChatSessionRequest) SetLicense(v string) *CreateChatSessionRequest {
	s.License = &v
	return s
}

func (s *CreateChatSessionRequest) SetPlatform(v string) *CreateChatSessionRequest {
	s.Platform = &v
	return s
}

func (s *CreateChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
