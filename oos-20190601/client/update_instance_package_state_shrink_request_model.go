// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePackageStateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigureAction(v string) *UpdateInstancePackageStateShrinkRequest
	GetConfigureAction() *string
	SetInstanceId(v string) *UpdateInstancePackageStateShrinkRequest
	GetInstanceId() *string
	SetParametersShrink(v string) *UpdateInstancePackageStateShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *UpdateInstancePackageStateShrinkRequest
	GetRegionId() *string
	SetTemplateName(v string) *UpdateInstancePackageStateShrinkRequest
	GetTemplateName() *string
	SetTemplateVersion(v string) *UpdateInstancePackageStateShrinkRequest
	GetTemplateVersion() *string
}

type UpdateInstancePackageStateShrinkRequest struct {
	// The operation type.
	//
	// Valid values:
	//
	// 	- uninstall
	//
	// 	- install
	//
	// This parameter is required.
	//
	// example:
	//
	// install
	ConfigureAction *string `json:"ConfigureAction,omitempty" xml:"ConfigureAction,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1jaxa2bs4bps7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The parameters for installing or uninstalling the extensions.
	//
	// example:
	//
	// {"username": "xx"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateInstancePackageStateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePackageStateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateShrinkRequest) GetConfigureAction() *string {
	return s.ConfigureAction
}

func (s *UpdateInstancePackageStateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstancePackageStateShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateInstancePackageStateShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstancePackageStateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateInstancePackageStateShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateInstancePackageStateShrinkRequest) SetConfigureAction(v string) *UpdateInstancePackageStateShrinkRequest {
	s.ConfigureAction = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetInstanceId(v string) *UpdateInstancePackageStateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetParametersShrink(v string) *UpdateInstancePackageStateShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetRegionId(v string) *UpdateInstancePackageStateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetTemplateName(v string) *UpdateInstancePackageStateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetTemplateVersion(v string) *UpdateInstancePackageStateShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
