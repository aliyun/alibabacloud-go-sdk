// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfiguration(v *UpdateAccessConfigurationResponseBodyAccessConfiguration) *UpdateAccessConfigurationResponseBody
	GetAccessConfiguration() *UpdateAccessConfigurationResponseBodyAccessConfiguration
	SetRequestId(v string) *UpdateAccessConfigurationResponseBody
	GetRequestId() *string
}

type UpdateAccessConfigurationResponseBody struct {
	// The information about the access configuration.
	AccessConfiguration *UpdateAccessConfigurationResponseBodyAccessConfiguration `json:"AccessConfiguration,omitempty" xml:"AccessConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B13E4EE-3853-5852-9165-597C32AD8FB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationResponseBody) GetAccessConfiguration() *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	return s.AccessConfiguration
}

func (s *UpdateAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccessConfigurationResponseBody) SetAccessConfiguration(v *UpdateAccessConfigurationResponseBodyAccessConfiguration) *UpdateAccessConfigurationResponseBody {
	s.AccessConfiguration = v
	return s
}

func (s *UpdateAccessConfigurationResponseBody) SetRequestId(v string) *UpdateAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBody) Validate() error {
	if s.AccessConfiguration != nil {
		if err := s.AccessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAccessConfigurationResponseBodyAccessConfiguration struct {
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
	// The initial web page
	//
	// that is displayed after a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// example:
	//
	// https://cloudsso.console.aliyun.com
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The duration of a session
	//
	// in which a CloudSSO user accesses an account in your resource directory by using the access configuration.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 3600
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The status notifications.
	StatusNotifications []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	// The time when the information about the access configuration was modified.
	//
	// example:
	//
	// 2021-11-02T10:10:01Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateAccessConfigurationResponseBodyAccessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessConfigurationResponseBodyAccessConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetDescription() *string {
	return s.Description
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetRelayState() *string {
	return s.RelayState
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetSessionDuration() *int32 {
	return s.SessionDuration
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetStatusNotifications() []*string {
	return s.StatusNotifications
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationId(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationId = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetAccessConfigurationName(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.AccessConfigurationName = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetCreateTime(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.CreateTime = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetDescription(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.Description = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetRelayState(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.RelayState = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetSessionDuration(v int32) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetStatusNotifications(v []*string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.StatusNotifications = v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) SetUpdateTime(v string) *UpdateAccessConfigurationResponseBodyAccessConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *UpdateAccessConfigurationResponseBodyAccessConfiguration) Validate() error {
	return dara.Validate(s)
}
