// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *UpdateLivePackageChannelCredentialsRequest
	GetChannelName() *string
	SetGroupName(v string) *UpdateLivePackageChannelCredentialsRequest
	GetGroupName() *string
	SetRotateCredentials(v int32) *UpdateLivePackageChannelCredentialsRequest
	GetRotateCredentials() *int32
}

type UpdateLivePackageChannelCredentialsRequest struct {
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether to update the credentials. 1: updates the credentials of endpoint 1. 2: updates the credentials of endpoint 2. 3: updates the credentials of endpoints 1 and 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	RotateCredentials *int32 `json:"RotateCredentials,omitempty" xml:"RotateCredentials,omitempty"`
}

func (s UpdateLivePackageChannelCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelCredentialsRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelCredentialsRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateLivePackageChannelCredentialsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageChannelCredentialsRequest) GetRotateCredentials() *int32 {
	return s.RotateCredentials
}

func (s *UpdateLivePackageChannelCredentialsRequest) SetChannelName(v string) *UpdateLivePackageChannelCredentialsRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsRequest) SetGroupName(v string) *UpdateLivePackageChannelCredentialsRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsRequest) SetRotateCredentials(v int32) *UpdateLivePackageChannelCredentialsRequest {
	s.RotateCredentials = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
