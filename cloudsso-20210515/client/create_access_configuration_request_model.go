// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationName(v string) *CreateAccessConfigurationRequest
	GetAccessConfigurationName() *string
	SetDescription(v string) *CreateAccessConfigurationRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateAccessConfigurationRequest
	GetDirectoryId() *string
	SetRelayState(v string) *CreateAccessConfigurationRequest
	GetRelayState() *string
	SetSessionDuration(v int32) *CreateAccessConfigurationRequest
	GetSessionDuration() *int32
	SetTags(v []*CreateAccessConfigurationRequestTags) *CreateAccessConfigurationRequest
	GetTags() []*CreateAccessConfigurationRequestTags
}

type CreateAccessConfigurationRequest struct {
	// The name of the access configuration.
	//
	// Format: contains letters, digits, or hyphens (-).
	//
	// Length: up to 32 characters.
	//
	// example:
	//
	// ECS-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The description of the access configuration.
	//
	// Length: up to 1024 characters.
	//
	// example:
	//
	// This is an access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The initial access page.
	//
	// The page address that a CloudSSO user initially accesses when using the access configuration to access an account in a resource directory.
	//
	// The page must be an Alibaba Cloud Management Console page. Default value: empty, which indicates that the user is redirected to the homepage of the Alibaba Cloud Management Console.
	//
	// example:
	//
	// https://cloudsso.console.aliyun.com
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The session duration.
	//
	// The maximum duration of a session when a CloudSSO user uses the access configuration to access an account in a resource directory.
	//
	// Unit: seconds.
	//
	// Valid values: 900 to 43200 (15 minutes to 12 hours).
	//
	// Default value: 3600 (1 hour).
	//
	// example:
	//
	// 3600
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The list of tags.
	Tags []*CreateAccessConfigurationRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationRequest) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *CreateAccessConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateAccessConfigurationRequest) GetRelayState() *string {
	return s.RelayState
}

func (s *CreateAccessConfigurationRequest) GetSessionDuration() *int32 {
	return s.SessionDuration
}

func (s *CreateAccessConfigurationRequest) GetTags() []*CreateAccessConfigurationRequestTags {
	return s.Tags
}

func (s *CreateAccessConfigurationRequest) SetAccessConfigurationName(v string) *CreateAccessConfigurationRequest {
	s.AccessConfigurationName = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetDescription(v string) *CreateAccessConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetDirectoryId(v string) *CreateAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetRelayState(v string) *CreateAccessConfigurationRequest {
	s.RelayState = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetSessionDuration(v int32) *CreateAccessConfigurationRequest {
	s.SessionDuration = &v
	return s
}

func (s *CreateAccessConfigurationRequest) SetTags(v []*CreateAccessConfigurationRequestTags) *CreateAccessConfigurationRequest {
	s.Tags = v
	return s
}

func (s *CreateAccessConfigurationRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAccessConfigurationRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAccessConfigurationRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessConfigurationRequestTags) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateAccessConfigurationRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateAccessConfigurationRequestTags) SetKey(v string) *CreateAccessConfigurationRequestTags {
	s.Key = &v
	return s
}

func (s *CreateAccessConfigurationRequestTags) SetValue(v string) *CreateAccessConfigurationRequestTags {
	s.Value = &v
	return s
}

func (s *CreateAccessConfigurationRequestTags) Validate() error {
	return dara.Validate(s)
}
