// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfiguration(v *GetAccessConfigurationResponseBodyAccessConfiguration) *GetAccessConfigurationResponseBody
	GetAccessConfiguration() *GetAccessConfigurationResponseBodyAccessConfiguration
	SetRequestId(v string) *GetAccessConfigurationResponseBody
	GetRequestId() *string
}

type GetAccessConfigurationResponseBody struct {
	// The access configuration information.
	AccessConfiguration *GetAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D5E40508-483B-52F6-993C-D880B0F87591
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationResponseBody) GetAccessConfiguration() *GetAccessConfigurationResponseBodyAccessConfiguration {
	return s.AccessConfiguration
}

func (s *GetAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessConfigurationResponseBody) SetAccessConfiguration(v *GetAccessConfigurationResponseBodyAccessConfiguration) *GetAccessConfigurationResponseBody {
	s.AccessConfiguration = v
	return s
}

func (s *GetAccessConfigurationResponseBody) SetRequestId(v string) *GetAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessConfigurationResponseBody) Validate() error {
	if s.AccessConfiguration != nil {
		if err := s.AccessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessConfigurationResponseBodyAccessConfiguration struct {
	// The access configuration ID.
	//
	// example:
	//
	// ac-00ccule7tadaijxc****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// VPC-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access configuration was created.
	//
	// example:
	//
	// 2021-06-30T09:39:44Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access configuration.
	//
	// example:
	//
	// This is an access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial access page.
	//
	// The URL of the initial page that is displayed when a CloudSSO user uses the access configuration to access an account in a resource directory.
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
	// The tags.
	Tags []*GetAccessConfigurationResponseBodyAccessConfigurationTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the access configuration was last modified.
	//
	// example:
	//
	// 2021-07-26T03:02:11Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAccessConfigurationResponseBodyAccessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetAccessConfigurationResponseBodyAccessConfiguration) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetDescription() *string {
	return s.Description
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetRelayState() *string {
	return s.RelayState
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetSessionDuration() *int32 {
	return s.SessionDuration
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetStatusNotifications() []*string {
	return s.StatusNotifications
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetTags() []*GetAccessConfigurationResponseBodyAccessConfigurationTags {
	return s.Tags
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationId(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationId = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationName(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationName = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetCreateTime(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.CreateTime = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetDescription(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.Description = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetRelayState(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.RelayState = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetSessionDuration(v int32) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetStatusNotifications(v []*string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.StatusNotifications = v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetTags(v []*GetAccessConfigurationResponseBodyAccessConfigurationTags) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.Tags = v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) SetUpdateTime(v string) *GetAccessConfigurationResponseBodyAccessConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfiguration) Validate() error {
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

type GetAccessConfigurationResponseBodyAccessConfigurationTags struct {
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

func (s GetAccessConfigurationResponseBodyAccessConfigurationTags) String() string {
	return dara.Prettify(s)
}

func (s GetAccessConfigurationResponseBodyAccessConfigurationTags) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationResponseBodyAccessConfigurationTags) GetKey() *string {
	return s.Key
}

func (s *GetAccessConfigurationResponseBodyAccessConfigurationTags) GetValue() *string {
	return s.Value
}

func (s *GetAccessConfigurationResponseBodyAccessConfigurationTags) SetKey(v string) *GetAccessConfigurationResponseBodyAccessConfigurationTags {
	s.Key = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfigurationTags) SetValue(v string) *GetAccessConfigurationResponseBodyAccessConfigurationTags {
	s.Value = &v
	return s
}

func (s *GetAccessConfigurationResponseBodyAccessConfigurationTags) Validate() error {
	return dara.Validate(s)
}
