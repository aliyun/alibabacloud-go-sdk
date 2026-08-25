// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfiguration(v *CreateAccessConfigurationResponseBodyAccessConfiguration) *CreateAccessConfigurationResponseBody
	GetAccessConfiguration() *CreateAccessConfigurationResponseBodyAccessConfiguration
	SetRequestId(v string) *CreateAccessConfigurationResponseBody
	GetRequestId() *string
}

type CreateAccessConfigurationResponseBody struct {
	// The access configuration information.
	AccessConfiguration *CreateAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A3A41736-A050-50B6-ABC5-590F376A0044
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationResponseBody) GetAccessConfiguration() *CreateAccessConfigurationResponseBodyAccessConfiguration {
	return s.AccessConfiguration
}

func (s *CreateAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessConfigurationResponseBody) SetAccessConfiguration(v *CreateAccessConfigurationResponseBodyAccessConfiguration) *CreateAccessConfigurationResponseBody {
	s.AccessConfiguration = v
	return s
}

func (s *CreateAccessConfigurationResponseBody) SetRequestId(v string) *CreateAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessConfigurationResponseBody) Validate() error {
	if s.AccessConfiguration != nil {
		if err := s.AccessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccessConfigurationResponseBodyAccessConfiguration struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// ECS-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access configuration was created.
	//
	// example:
	//
	// 2021-11-02T08:44:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access configuration.
	//
	// example:
	//
	// This is an access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial access page.
	//
	// The page address that a CloudSSO user initially accesses when using the access configuration to access an account in a resource directory.
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
	// example:
	//
	// 3600
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The status notification information.
	StatusNotifications []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*CreateAccessConfigurationResponseBodyAccessConfigurationTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the access configuration was last modified.
	//
	// example:
	//
	// 2021-11-02T08:44:23Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateAccessConfigurationResponseBodyAccessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessConfigurationResponseBodyAccessConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetDescription() *string {
	return s.Description
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetRelayState() *string {
	return s.RelayState
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetSessionDuration() *int32 {
	return s.SessionDuration
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetStatusNotifications() []*string {
	return s.StatusNotifications
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetTags() []*CreateAccessConfigurationResponseBodyAccessConfigurationTags {
	return s.Tags
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationId(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationId = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationName(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationName = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetCreateTime(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.CreateTime = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetDescription(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.Description = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetRelayState(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.RelayState = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetSessionDuration(v int32) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetStatusNotifications(v []*string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.StatusNotifications = v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetTags(v []*CreateAccessConfigurationResponseBodyAccessConfigurationTags) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.Tags = v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) SetUpdateTime(v string) *CreateAccessConfigurationResponseBodyAccessConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfiguration) Validate() error {
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

type CreateAccessConfigurationResponseBodyAccessConfigurationTags struct {
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

func (s CreateAccessConfigurationResponseBodyAccessConfigurationTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessConfigurationResponseBodyAccessConfigurationTags) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationResponseBodyAccessConfigurationTags) GetKey() *string {
	return s.Key
}

func (s *CreateAccessConfigurationResponseBodyAccessConfigurationTags) GetValue() *string {
	return s.Value
}

func (s *CreateAccessConfigurationResponseBodyAccessConfigurationTags) SetKey(v string) *CreateAccessConfigurationResponseBodyAccessConfigurationTags {
	s.Key = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfigurationTags) SetValue(v string) *CreateAccessConfigurationResponseBodyAccessConfigurationTags {
	s.Value = &v
	return s
}

func (s *CreateAccessConfigurationResponseBodyAccessConfigurationTags) Validate() error {
	return dara.Validate(s)
}
