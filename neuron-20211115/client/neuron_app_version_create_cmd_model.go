// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppVersionCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *NeuronAppVersionCreateCmd
	GetAccountId() *string
	SetAppEntry(v string) *NeuronAppVersionCreateCmd
	GetAppEntry() *string
	SetAppId(v int64) *NeuronAppVersionCreateCmd
	GetAppId() *int64
	SetDescription(v string) *NeuronAppVersionCreateCmd
	GetDescription() *string
	SetEnvs(v []*string) *NeuronAppVersionCreateCmd
	GetEnvs() []*string
	SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersionCreateCmd
	GetFeatureDesc() []*NeuronAppInfoContent
	SetImageUrls(v []*string) *NeuronAppVersionCreateCmd
	GetImageUrls() []*string
	SetInstructionUrl(v string) *NeuronAppVersionCreateCmd
	GetInstructionUrl() *string
	SetManageType(v string) *NeuronAppVersionCreateCmd
	GetManageType() *string
	SetMobiId(v int64) *NeuronAppVersionCreateCmd
	GetMobiId() *int64
	SetPbcs(v []*int64) *NeuronAppVersionCreateCmd
	GetPbcs() []*int64
	SetPluginList(v []*AppPluginInfo) *NeuronAppVersionCreateCmd
	GetPluginList() []*AppPluginInfo
	SetPrivateInfo(v []*string) *NeuronAppVersionCreateCmd
	GetPrivateInfo() []*string
	SetScopes(v []*string) *NeuronAppVersionCreateCmd
	GetScopes() []*string
	SetSidebar(v string) *NeuronAppVersionCreateCmd
	GetSidebar() *string
	SetUnbindEffect(v string) *NeuronAppVersionCreateCmd
	GetUnbindEffect() *string
	SetUnbindReasons(v []*string) *NeuronAppVersionCreateCmd
	GetUnbindReasons() []*string
	SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersionCreateCmd
	GetUpdatedFeature() []*NeuronAppInfoContent
	SetVersion(v string) *NeuronAppVersionCreateCmd
	GetVersion() *string
}

type NeuronAppVersionCreateCmd struct {
	// example:
	//
	// 1343424
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AppEntry  *string `json:"appEntry,omitempty" xml:"appEntry,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 升级订单功能
	Description    *string                 `json:"description,omitempty" xml:"description,omitempty"`
	Envs           []*string               `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
	FeatureDesc    []*NeuronAppInfoContent `json:"featureDesc,omitempty" xml:"featureDesc,omitempty" type:"Repeated"`
	ImageUrls      []*string               `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	InstructionUrl *string                 `json:"instructionUrl,omitempty" xml:"instructionUrl,omitempty"`
	ManageType     *string                 `json:"manageType,omitempty" xml:"manageType,omitempty"`
	// example:
	//
	// 1
	MobiId      *int64           `json:"mobiId,omitempty" xml:"mobiId,omitempty"`
	Pbcs        []*int64         `json:"pbcs,omitempty" xml:"pbcs,omitempty" type:"Repeated"`
	PluginList  []*AppPluginInfo `json:"pluginList,omitempty" xml:"pluginList,omitempty" type:"Repeated"`
	PrivateInfo []*string        `json:"privateInfo,omitempty" xml:"privateInfo,omitempty" type:"Repeated"`
	Scopes      []*string        `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Sidebar *string `json:"sidebar,omitempty" xml:"sidebar,omitempty"`
	// example:
	//
	// sda
	UnbindEffect   *string                 `json:"unbindEffect,omitempty" xml:"unbindEffect,omitempty"`
	UnbindReasons  []*string               `json:"unbindReasons,omitempty" xml:"unbindReasons,omitempty" type:"Repeated"`
	UpdatedFeature []*NeuronAppInfoContent `json:"updatedFeature,omitempty" xml:"updatedFeature,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s NeuronAppVersionCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppVersionCreateCmd) GoString() string {
	return s.String()
}

func (s *NeuronAppVersionCreateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *NeuronAppVersionCreateCmd) GetAppEntry() *string {
	return s.AppEntry
}

func (s *NeuronAppVersionCreateCmd) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronAppVersionCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *NeuronAppVersionCreateCmd) GetEnvs() []*string {
	return s.Envs
}

func (s *NeuronAppVersionCreateCmd) GetFeatureDesc() []*NeuronAppInfoContent {
	return s.FeatureDesc
}

func (s *NeuronAppVersionCreateCmd) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *NeuronAppVersionCreateCmd) GetInstructionUrl() *string {
	return s.InstructionUrl
}

func (s *NeuronAppVersionCreateCmd) GetManageType() *string {
	return s.ManageType
}

func (s *NeuronAppVersionCreateCmd) GetMobiId() *int64 {
	return s.MobiId
}

func (s *NeuronAppVersionCreateCmd) GetPbcs() []*int64 {
	return s.Pbcs
}

func (s *NeuronAppVersionCreateCmd) GetPluginList() []*AppPluginInfo {
	return s.PluginList
}

func (s *NeuronAppVersionCreateCmd) GetPrivateInfo() []*string {
	return s.PrivateInfo
}

func (s *NeuronAppVersionCreateCmd) GetScopes() []*string {
	return s.Scopes
}

func (s *NeuronAppVersionCreateCmd) GetSidebar() *string {
	return s.Sidebar
}

func (s *NeuronAppVersionCreateCmd) GetUnbindEffect() *string {
	return s.UnbindEffect
}

func (s *NeuronAppVersionCreateCmd) GetUnbindReasons() []*string {
	return s.UnbindReasons
}

func (s *NeuronAppVersionCreateCmd) GetUpdatedFeature() []*NeuronAppInfoContent {
	return s.UpdatedFeature
}

func (s *NeuronAppVersionCreateCmd) GetVersion() *string {
	return s.Version
}

func (s *NeuronAppVersionCreateCmd) SetAccountId(v string) *NeuronAppVersionCreateCmd {
	s.AccountId = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetAppEntry(v string) *NeuronAppVersionCreateCmd {
	s.AppEntry = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetAppId(v int64) *NeuronAppVersionCreateCmd {
	s.AppId = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetDescription(v string) *NeuronAppVersionCreateCmd {
	s.Description = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetEnvs(v []*string) *NeuronAppVersionCreateCmd {
	s.Envs = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersionCreateCmd {
	s.FeatureDesc = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetImageUrls(v []*string) *NeuronAppVersionCreateCmd {
	s.ImageUrls = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetInstructionUrl(v string) *NeuronAppVersionCreateCmd {
	s.InstructionUrl = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetManageType(v string) *NeuronAppVersionCreateCmd {
	s.ManageType = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetMobiId(v int64) *NeuronAppVersionCreateCmd {
	s.MobiId = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetPbcs(v []*int64) *NeuronAppVersionCreateCmd {
	s.Pbcs = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetPluginList(v []*AppPluginInfo) *NeuronAppVersionCreateCmd {
	s.PluginList = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetPrivateInfo(v []*string) *NeuronAppVersionCreateCmd {
	s.PrivateInfo = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetScopes(v []*string) *NeuronAppVersionCreateCmd {
	s.Scopes = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetSidebar(v string) *NeuronAppVersionCreateCmd {
	s.Sidebar = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetUnbindEffect(v string) *NeuronAppVersionCreateCmd {
	s.UnbindEffect = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetUnbindReasons(v []*string) *NeuronAppVersionCreateCmd {
	s.UnbindReasons = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersionCreateCmd {
	s.UpdatedFeature = v
	return s
}

func (s *NeuronAppVersionCreateCmd) SetVersion(v string) *NeuronAppVersionCreateCmd {
	s.Version = &v
	return s
}

func (s *NeuronAppVersionCreateCmd) Validate() error {
	if s.FeatureDesc != nil {
		for _, item := range s.FeatureDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PluginList != nil {
		for _, item := range s.PluginList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpdatedFeature != nil {
		for _, item := range s.UpdatedFeature {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
