// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserProvisioningConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateUserProvisioningConfigurationRequest
	GetDirectoryId() *string
	SetNewDefaultLandingPage(v string) *UpdateUserProvisioningConfigurationRequest
	GetNewDefaultLandingPage() *string
	SetNewSessionDuration(v int32) *UpdateUserProvisioningConfigurationRequest
	GetNewSessionDuration() *int32
}

type UpdateUserProvisioningConfigurationRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new default URL for a CloudSSO user who logs on to the Alibaba Cloud Management Console.
	//
	// Default value: https://homenew.console.aliyun.com.
	//
	// example:
	//
	// https://home.console.aliyun.com/home/dashboard/ProductAndService
	NewDefaultLandingPage *string `json:"NewDefaultLandingPage,omitempty" xml:"NewDefaultLandingPage,omitempty"`
	// The new duration of the logon session.
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
	NewSessionDuration *int32 `json:"NewSessionDuration,omitempty" xml:"NewSessionDuration,omitempty"`
}

func (s UpdateUserProvisioningConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateUserProvisioningConfigurationRequest) GetNewDefaultLandingPage() *string {
	return s.NewDefaultLandingPage
}

func (s *UpdateUserProvisioningConfigurationRequest) GetNewSessionDuration() *int32 {
	return s.NewSessionDuration
}

func (s *UpdateUserProvisioningConfigurationRequest) SetDirectoryId(v string) *UpdateUserProvisioningConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationRequest) SetNewDefaultLandingPage(v string) *UpdateUserProvisioningConfigurationRequest {
	s.NewDefaultLandingPage = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationRequest) SetNewSessionDuration(v int32) *UpdateUserProvisioningConfigurationRequest {
	s.NewSessionDuration = &v
	return s
}

func (s *UpdateUserProvisioningConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
