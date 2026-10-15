// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *UpdateAppPluginRequest
	GetCategory() *string
	SetConfigItems(v string) *UpdateAppPluginRequest
	GetConfigItems() *string
	SetDescription(v string) *UpdateAppPluginRequest
	GetDescription() *string
	SetExtend(v string) *UpdateAppPluginRequest
	GetExtend() *string
	SetHooks(v string) *UpdateAppPluginRequest
	GetHooks() *string
	SetIcon(v string) *UpdateAppPluginRequest
	GetIcon() *string
	SetPluginId(v string) *UpdateAppPluginRequest
	GetPluginId() *string
	SetPluginName(v string) *UpdateAppPluginRequest
	GetPluginName() *string
	SetPluginVersion(v string) *UpdateAppPluginRequest
	GetPluginVersion() *string
	SetSkillHeader(v string) *UpdateAppPluginRequest
	GetSkillHeader() *string
	SetTags(v string) *UpdateAppPluginRequest
	GetTags() *string
	SetVisibility(v string) *UpdateAppPluginRequest
	GetVisibility() *string
}

type UpdateAppPluginRequest struct {
	// The category.
	//
	// example:
	//
	// DataManagement
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The configuration form in React JSON Schema format.
	//
	// example:
	//
	// ***
	ConfigItems *string `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty"`
	// The application description.
	//
	// example:
	//
	// myapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Other extension information in JSON format for future parameter expansion.
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The hook definitions.
	//
	// example:
	//
	// []
	Hooks *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// The plug-in description.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01ANcLnQ1bCX7tZPVMJ_!!6000000003429-2-tps-100-100.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The gateway plug-in ID.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plug-in name.
	//
	// example:
	//
	// alisecguard
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The plug-in version.
	//
	// example:
	//
	// v2
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
	// The skill header information for model selection.
	//
	// example:
	//
	// {}
	SkillHeader *string `json:"SkillHeader,omitempty" xml:"SkillHeader,omitempty"`
	// The labels.
	//
	// example:
	//
	// []
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The visibility.
	//
	// example:
	//
	// private
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s UpdateAppPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPluginRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppPluginRequest) GetCategory() *string {
	return s.Category
}

func (s *UpdateAppPluginRequest) GetConfigItems() *string {
	return s.ConfigItems
}

func (s *UpdateAppPluginRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppPluginRequest) GetExtend() *string {
	return s.Extend
}

func (s *UpdateAppPluginRequest) GetHooks() *string {
	return s.Hooks
}

func (s *UpdateAppPluginRequest) GetIcon() *string {
	return s.Icon
}

func (s *UpdateAppPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *UpdateAppPluginRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *UpdateAppPluginRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *UpdateAppPluginRequest) GetSkillHeader() *string {
	return s.SkillHeader
}

func (s *UpdateAppPluginRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateAppPluginRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateAppPluginRequest) SetCategory(v string) *UpdateAppPluginRequest {
	s.Category = &v
	return s
}

func (s *UpdateAppPluginRequest) SetConfigItems(v string) *UpdateAppPluginRequest {
	s.ConfigItems = &v
	return s
}

func (s *UpdateAppPluginRequest) SetDescription(v string) *UpdateAppPluginRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppPluginRequest) SetExtend(v string) *UpdateAppPluginRequest {
	s.Extend = &v
	return s
}

func (s *UpdateAppPluginRequest) SetHooks(v string) *UpdateAppPluginRequest {
	s.Hooks = &v
	return s
}

func (s *UpdateAppPluginRequest) SetIcon(v string) *UpdateAppPluginRequest {
	s.Icon = &v
	return s
}

func (s *UpdateAppPluginRequest) SetPluginId(v string) *UpdateAppPluginRequest {
	s.PluginId = &v
	return s
}

func (s *UpdateAppPluginRequest) SetPluginName(v string) *UpdateAppPluginRequest {
	s.PluginName = &v
	return s
}

func (s *UpdateAppPluginRequest) SetPluginVersion(v string) *UpdateAppPluginRequest {
	s.PluginVersion = &v
	return s
}

func (s *UpdateAppPluginRequest) SetSkillHeader(v string) *UpdateAppPluginRequest {
	s.SkillHeader = &v
	return s
}

func (s *UpdateAppPluginRequest) SetTags(v string) *UpdateAppPluginRequest {
	s.Tags = &v
	return s
}

func (s *UpdateAppPluginRequest) SetVisibility(v string) *UpdateAppPluginRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateAppPluginRequest) Validate() error {
	return dara.Validate(s)
}
