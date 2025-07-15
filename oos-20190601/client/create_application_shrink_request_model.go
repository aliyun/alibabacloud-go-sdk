// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfigShrink(v string) *CreateApplicationShrinkRequest
	GetAlarmConfigShrink() *string
	SetApplicationSource(v string) *CreateApplicationShrinkRequest
	GetApplicationSource() *string
	SetClientToken(v string) *CreateApplicationShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateApplicationShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateApplicationShrinkRequest
	GetName() *string
	SetRegionId(v string) *CreateApplicationShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateApplicationShrinkRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateApplicationShrinkRequest
	GetServiceId() *string
	SetTagsShrink(v string) *CreateApplicationShrinkRequest
	GetTagsShrink() *string
}

type CreateApplicationShrinkRequest struct {
	// The configurations of application alerts.
	AlarmConfigShrink *string `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty"`
	// The source of application.
	//
	// example:
	//
	// {"Platform":"github","Owner":"githubUser","RepoName":"githubUser/repoName"}
	ApplicationSource *string `json:"ApplicationSource,omitempty" xml:"ApplicationSource,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// TF-CreateApplication-1647587475-84104b89-eba5-47a8-b2fd-807b8b7d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the Compute Nest service that corresponds to the application template.
	//
	// example:
	//
	// service-79538e30e44541b699d8
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) GetAlarmConfigShrink() *string {
	return s.AlarmConfigShrink
}

func (s *CreateApplicationShrinkRequest) GetApplicationSource() *string {
	return s.ApplicationSource
}

func (s *CreateApplicationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApplicationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateApplicationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateApplicationShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateApplicationShrinkRequest) SetAlarmConfigShrink(v string) *CreateApplicationShrinkRequest {
	s.AlarmConfigShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetApplicationSource(v string) *CreateApplicationShrinkRequest {
	s.ApplicationSource = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetClientToken(v string) *CreateApplicationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDescription(v string) *CreateApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetName(v string) *CreateApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetRegionId(v string) *CreateApplicationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetResourceGroupId(v string) *CreateApplicationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetServiceId(v string) *CreateApplicationShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTagsShrink(v string) *CreateApplicationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
