// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppVersionDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAppDesc(v string) *NeuronAppVersionDetail
	GetAppDesc() *string
	SetAppEntry(v string) *NeuronAppVersionDetail
	GetAppEntry() *string
	SetAppId(v int64) *NeuronAppVersionDetail
	GetAppId() *int64
	SetAppName(v string) *NeuronAppVersionDetail
	GetAppName() *string
	SetEnterpriseId(v int64) *NeuronAppVersionDetail
	GetEnterpriseId() *int64
	SetEnvs(v []*string) *NeuronAppVersionDetail
	GetEnvs() []*string
	SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersionDetail
	GetFeatureDesc() []*NeuronAppInfoContent
	SetIconUrl(v string) *NeuronAppVersionDetail
	GetIconUrl() *string
	SetImageUrls(v []*string) *NeuronAppVersionDetail
	GetImageUrls() []*string
	SetInstructionUrl(v string) *NeuronAppVersionDetail
	GetInstructionUrl() *string
	SetManageType(v string) *NeuronAppVersionDetail
	GetManageType() *string
	SetMobiId(v int64) *NeuronAppVersionDetail
	GetMobiId() *int64
	SetOwner(v string) *NeuronAppVersionDetail
	GetOwner() *string
	SetPbcs(v []*int64) *NeuronAppVersionDetail
	GetPbcs() []*int64
	SetPluginList(v []*AppPluginInfo) *NeuronAppVersionDetail
	GetPluginList() []*AppPluginInfo
	SetPrivateInfo(v []*string) *NeuronAppVersionDetail
	GetPrivateInfo() []*string
	SetProductId(v int64) *NeuronAppVersionDetail
	GetProductId() *int64
	SetScopes(v []*string) *NeuronAppVersionDetail
	GetScopes() []*string
	SetSidebar(v string) *NeuronAppVersionDetail
	GetSidebar() *string
	SetUnbindEffect(v string) *NeuronAppVersionDetail
	GetUnbindEffect() *string
	SetUnbindReasons(v []*string) *NeuronAppVersionDetail
	GetUnbindReasons() []*string
	SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersionDetail
	GetUpdatedFeature() []*NeuronAppInfoContent
	SetVersion(v string) *NeuronAppVersionDetail
	GetVersion() *string
	SetVersionId(v int64) *NeuronAppVersionDetail
	GetVersionId() *int64
}

type NeuronAppVersionDetail struct {
	// example:
	//
	// 升级订单功能
	AppDesc  *string `json:"appDesc,omitempty" xml:"appDesc,omitempty"`
	AppEntry *string `json:"appEntry,omitempty" xml:"appEntry,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// order
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64                  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Envs         []*string               `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
	FeatureDesc  []*NeuronAppInfoContent `json:"featureDesc,omitempty" xml:"featureDesc,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// http://wwww.ali.com
	IconUrl        *string   `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	ImageUrls      []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	InstructionUrl *string   `json:"instructionUrl,omitempty" xml:"instructionUrl,omitempty"`
	ManageType     *string   `json:"manageType,omitempty" xml:"manageType,omitempty"`
	// example:
	//
	// 1
	MobiId *int64 `json:"mobiId,omitempty" xml:"mobiId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 多端商城
	Owner       *string          `json:"owner,omitempty" xml:"owner,omitempty"`
	Pbcs        []*int64         `json:"pbcs,omitempty" xml:"pbcs,omitempty" type:"Repeated"`
	PluginList  []*AppPluginInfo `json:"pluginList,omitempty" xml:"pluginList,omitempty" type:"Repeated"`
	PrivateInfo []*string        `json:"privateInfo,omitempty" xml:"privateInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ProductId *int64    `json:"productId,omitempty" xml:"productId,omitempty"`
	Scopes    []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
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
	// example:
	//
	// 1
	VersionId *int64 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s NeuronAppVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppVersionDetail) GoString() string {
	return s.String()
}

func (s *NeuronAppVersionDetail) GetAppDesc() *string {
	return s.AppDesc
}

func (s *NeuronAppVersionDetail) GetAppEntry() *string {
	return s.AppEntry
}

func (s *NeuronAppVersionDetail) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronAppVersionDetail) GetAppName() *string {
	return s.AppName
}

func (s *NeuronAppVersionDetail) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *NeuronAppVersionDetail) GetEnvs() []*string {
	return s.Envs
}

func (s *NeuronAppVersionDetail) GetFeatureDesc() []*NeuronAppInfoContent {
	return s.FeatureDesc
}

func (s *NeuronAppVersionDetail) GetIconUrl() *string {
	return s.IconUrl
}

func (s *NeuronAppVersionDetail) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *NeuronAppVersionDetail) GetInstructionUrl() *string {
	return s.InstructionUrl
}

func (s *NeuronAppVersionDetail) GetManageType() *string {
	return s.ManageType
}

func (s *NeuronAppVersionDetail) GetMobiId() *int64 {
	return s.MobiId
}

func (s *NeuronAppVersionDetail) GetOwner() *string {
	return s.Owner
}

func (s *NeuronAppVersionDetail) GetPbcs() []*int64 {
	return s.Pbcs
}

func (s *NeuronAppVersionDetail) GetPluginList() []*AppPluginInfo {
	return s.PluginList
}

func (s *NeuronAppVersionDetail) GetPrivateInfo() []*string {
	return s.PrivateInfo
}

func (s *NeuronAppVersionDetail) GetProductId() *int64 {
	return s.ProductId
}

func (s *NeuronAppVersionDetail) GetScopes() []*string {
	return s.Scopes
}

func (s *NeuronAppVersionDetail) GetSidebar() *string {
	return s.Sidebar
}

func (s *NeuronAppVersionDetail) GetUnbindEffect() *string {
	return s.UnbindEffect
}

func (s *NeuronAppVersionDetail) GetUnbindReasons() []*string {
	return s.UnbindReasons
}

func (s *NeuronAppVersionDetail) GetUpdatedFeature() []*NeuronAppInfoContent {
	return s.UpdatedFeature
}

func (s *NeuronAppVersionDetail) GetVersion() *string {
	return s.Version
}

func (s *NeuronAppVersionDetail) GetVersionId() *int64 {
	return s.VersionId
}

func (s *NeuronAppVersionDetail) SetAppDesc(v string) *NeuronAppVersionDetail {
	s.AppDesc = &v
	return s
}

func (s *NeuronAppVersionDetail) SetAppEntry(v string) *NeuronAppVersionDetail {
	s.AppEntry = &v
	return s
}

func (s *NeuronAppVersionDetail) SetAppId(v int64) *NeuronAppVersionDetail {
	s.AppId = &v
	return s
}

func (s *NeuronAppVersionDetail) SetAppName(v string) *NeuronAppVersionDetail {
	s.AppName = &v
	return s
}

func (s *NeuronAppVersionDetail) SetEnterpriseId(v int64) *NeuronAppVersionDetail {
	s.EnterpriseId = &v
	return s
}

func (s *NeuronAppVersionDetail) SetEnvs(v []*string) *NeuronAppVersionDetail {
	s.Envs = v
	return s
}

func (s *NeuronAppVersionDetail) SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersionDetail {
	s.FeatureDesc = v
	return s
}

func (s *NeuronAppVersionDetail) SetIconUrl(v string) *NeuronAppVersionDetail {
	s.IconUrl = &v
	return s
}

func (s *NeuronAppVersionDetail) SetImageUrls(v []*string) *NeuronAppVersionDetail {
	s.ImageUrls = v
	return s
}

func (s *NeuronAppVersionDetail) SetInstructionUrl(v string) *NeuronAppVersionDetail {
	s.InstructionUrl = &v
	return s
}

func (s *NeuronAppVersionDetail) SetManageType(v string) *NeuronAppVersionDetail {
	s.ManageType = &v
	return s
}

func (s *NeuronAppVersionDetail) SetMobiId(v int64) *NeuronAppVersionDetail {
	s.MobiId = &v
	return s
}

func (s *NeuronAppVersionDetail) SetOwner(v string) *NeuronAppVersionDetail {
	s.Owner = &v
	return s
}

func (s *NeuronAppVersionDetail) SetPbcs(v []*int64) *NeuronAppVersionDetail {
	s.Pbcs = v
	return s
}

func (s *NeuronAppVersionDetail) SetPluginList(v []*AppPluginInfo) *NeuronAppVersionDetail {
	s.PluginList = v
	return s
}

func (s *NeuronAppVersionDetail) SetPrivateInfo(v []*string) *NeuronAppVersionDetail {
	s.PrivateInfo = v
	return s
}

func (s *NeuronAppVersionDetail) SetProductId(v int64) *NeuronAppVersionDetail {
	s.ProductId = &v
	return s
}

func (s *NeuronAppVersionDetail) SetScopes(v []*string) *NeuronAppVersionDetail {
	s.Scopes = v
	return s
}

func (s *NeuronAppVersionDetail) SetSidebar(v string) *NeuronAppVersionDetail {
	s.Sidebar = &v
	return s
}

func (s *NeuronAppVersionDetail) SetUnbindEffect(v string) *NeuronAppVersionDetail {
	s.UnbindEffect = &v
	return s
}

func (s *NeuronAppVersionDetail) SetUnbindReasons(v []*string) *NeuronAppVersionDetail {
	s.UnbindReasons = v
	return s
}

func (s *NeuronAppVersionDetail) SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersionDetail {
	s.UpdatedFeature = v
	return s
}

func (s *NeuronAppVersionDetail) SetVersion(v string) *NeuronAppVersionDetail {
	s.Version = &v
	return s
}

func (s *NeuronAppVersionDetail) SetVersionId(v int64) *NeuronAppVersionDetail {
	s.VersionId = &v
	return s
}

func (s *NeuronAppVersionDetail) Validate() error {
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
