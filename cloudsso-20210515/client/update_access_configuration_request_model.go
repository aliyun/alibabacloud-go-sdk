// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *UpdateAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *UpdateAccessConfigurationRequest
	GetDirectoryId() *string
	SetNewDescription(v string) *UpdateAccessConfigurationRequest
	GetNewDescription() *string
	SetNewRelayState(v string) *UpdateAccessConfigurationRequest
	GetNewRelayState() *string
	SetNewSessionDuration(v int32) *UpdateAccessConfigurationRequest
	GetNewSessionDuration() *int32
}

type UpdateAccessConfigurationRequest struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new description of the access configuration.
	//
	// The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// This is an access configuration.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new initial web page
	//
	// that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// The web page must be a page of the Alibaba Cloud Management Console.
	//
	// example:
	//
	// https://cloudsso.console.aliyun.com
	NewRelayState *string `json:"NewRelayState,omitempty" xml:"NewRelayState,omitempty"`
	// The new duration of a session
	//
	// in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	//
	// Valid values: 900 to 43200. The value 900 indicates 15 minutes. The value 43200 indicates 12 hours.
	//
	// example:
	//
	// 3600
	NewSessionDuration *int32 `json:"NewSessionDuration,omitempty" xml:"NewSessionDuration,omitempty"`
}

func (s UpdateAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *UpdateAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateAccessConfigurationRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateAccessConfigurationRequest) GetNewRelayState() *string {
	return s.NewRelayState
}

func (s *UpdateAccessConfigurationRequest) GetNewSessionDuration() *int32 {
	return s.NewSessionDuration
}

func (s *UpdateAccessConfigurationRequest) SetAccessConfigurationId(v string) *UpdateAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetDirectoryId(v string) *UpdateAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetNewDescription(v string) *UpdateAccessConfigurationRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetNewRelayState(v string) *UpdateAccessConfigurationRequest {
	s.NewRelayState = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) SetNewSessionDuration(v int32) *UpdateAccessConfigurationRequest {
	s.NewSessionDuration = &v
	return s
}

func (s *UpdateAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
