// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserProvisioningConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserProvisioningConfigurationResponseBody
	GetRequestId() *string
	SetUserProvisioningConfiguration(v *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) *UpdateUserProvisioningConfigurationResponseBody
	GetUserProvisioningConfiguration() *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration
}

type UpdateUserProvisioningConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BBC2ED1D-FAC5-3DF8-B63C-992B85B08DD9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The global configurations of the RAM user provisioning.
	UserProvisioningConfiguration *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration `json:"UserProvisioningConfiguration,omitempty" xml:"UserProvisioningConfiguration,omitempty" type:"Struct"`
}

func (s UpdateUserProvisioningConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserProvisioningConfigurationResponseBody) GetUserProvisioningConfiguration() *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	return s.UserProvisioningConfiguration
}

func (s *UpdateUserProvisioningConfigurationResponseBody) SetRequestId(v string) *UpdateUserProvisioningConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBody) SetUserProvisioningConfiguration(v *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) *UpdateUserProvisioningConfigurationResponseBody {
	s.UserProvisioningConfiguration = v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBody) Validate() error {
	if s.UserProvisioningConfiguration != nil {
		if err := s.UserProvisioningConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration struct {
	// The creation time.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default URL for a CloudSSO user who logs on to the Alibaba Cloud Management Console.
	//
	// Default value: https://homenew.console.aliyun.com.
	//
	// example:
	//
	// https://homenew.console.aliyun.com
	DefaultLandingPage *string `json:"DefaultLandingPage,omitempty" xml:"DefaultLandingPage,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The duration of the logon session.
	//
	// Unit: hours.
	//
	// Valid values: 1 to 24.
	//
	// Default value: 6.
	//
	// example:
	//
	// 6
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetDefaultLandingPage() *string {
	return s.DefaultLandingPage
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetSessionDuration() *int32 {
	return s.SessionDuration
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetCreateTime(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDefaultLandingPage(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DefaultLandingPage = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDirectoryId(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetSessionDuration(v int32) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetUpdateTime(v string) *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) Validate() error {
	return dara.Validate(s)
}
