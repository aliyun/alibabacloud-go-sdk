// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePackageStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigureAction(v string) *UpdateInstancePackageStateRequest
	GetConfigureAction() *string
	SetInstanceId(v string) *UpdateInstancePackageStateRequest
	GetInstanceId() *string
	SetParameters(v map[string]interface{}) *UpdateInstancePackageStateRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *UpdateInstancePackageStateRequest
	GetRegionId() *string
	SetTemplateName(v string) *UpdateInstancePackageStateRequest
	GetTemplateName() *string
	SetTemplateVersion(v string) *UpdateInstancePackageStateRequest
	GetTemplateVersion() *string
}

type UpdateInstancePackageStateRequest struct {
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s UpdateInstancePackageStateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePackageStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateRequest) GetConfigureAction() *string {
	return s.ConfigureAction
}

func (s *UpdateInstancePackageStateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstancePackageStateRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateInstancePackageStateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstancePackageStateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateInstancePackageStateRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateInstancePackageStateRequest) SetConfigureAction(v string) *UpdateInstancePackageStateRequest {
	s.ConfigureAction = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetInstanceId(v string) *UpdateInstancePackageStateRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetParameters(v map[string]interface{}) *UpdateInstancePackageStateRequest {
	s.Parameters = v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetRegionId(v string) *UpdateInstancePackageStateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetTemplateName(v string) *UpdateInstancePackageStateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetTemplateVersion(v string) *UpdateInstancePackageStateRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) Validate() error {
	return dara.Validate(s)
}
