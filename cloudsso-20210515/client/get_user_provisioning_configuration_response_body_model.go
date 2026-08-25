// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserProvisioningConfigurationResponseBody
	GetRequestId() *string
	SetUserProvisioningConfiguration(v *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) *GetUserProvisioningConfigurationResponseBody
	GetUserProvisioningConfiguration() *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration
}

type GetUserProvisioningConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 66898413-EB80-556D-9429-06FE3548F672
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The global configurations of the RAM user provisioning.
	UserProvisioningConfiguration *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration `json:"UserProvisioningConfiguration,omitempty" xml:"UserProvisioningConfiguration,omitempty" type:"Struct"`
}

func (s GetUserProvisioningConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserProvisioningConfigurationResponseBody) GetUserProvisioningConfiguration() *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	return s.UserProvisioningConfiguration
}

func (s *GetUserProvisioningConfigurationResponseBody) SetRequestId(v string) *GetUserProvisioningConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBody) SetUserProvisioningConfiguration(v *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) *GetUserProvisioningConfigurationResponseBody {
	s.UserProvisioningConfiguration = v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBody) Validate() error {
	if s.UserProvisioningConfiguration != nil {
		if err := s.UserProvisioningConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration struct {
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
	// d-00fc2p61****
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
	// 10
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetDefaultLandingPage() *string {
	return s.DefaultLandingPage
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetSessionDuration() *int32 {
	return s.SessionDuration
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetCreateTime(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.CreateTime = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDefaultLandingPage(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DefaultLandingPage = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetDirectoryId(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetSessionDuration(v int32) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.SessionDuration = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) SetUpdateTime(v string) *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *GetUserProvisioningConfigurationResponseBodyUserProvisioningConfiguration) Validate() error {
	return dara.Validate(s)
}
