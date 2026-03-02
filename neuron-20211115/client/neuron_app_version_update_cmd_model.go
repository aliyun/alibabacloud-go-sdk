// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppVersionUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAppEntry(v string) *NeuronAppVersionUpdateCmd
	GetAppEntry() *string
	SetDesc(v string) *NeuronAppVersionUpdateCmd
	GetDesc() *string
	SetEnvs(v []*string) *NeuronAppVersionUpdateCmd
	GetEnvs() []*string
	SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersionUpdateCmd
	GetFeatureDesc() []*NeuronAppInfoContent
	SetId(v int64) *NeuronAppVersionUpdateCmd
	GetId() *int64
	SetImageUrls(v []*string) *NeuronAppVersionUpdateCmd
	GetImageUrls() []*string
	SetInstructionUrl(v string) *NeuronAppVersionUpdateCmd
	GetInstructionUrl() *string
	SetManageType(v string) *NeuronAppVersionUpdateCmd
	GetManageType() *string
	SetMobiId(v int64) *NeuronAppVersionUpdateCmd
	GetMobiId() *int64
	SetPbcs(v []*int64) *NeuronAppVersionUpdateCmd
	GetPbcs() []*int64
	SetPrivateInfo(v []*string) *NeuronAppVersionUpdateCmd
	GetPrivateInfo() []*string
	SetScopes(v []*string) *NeuronAppVersionUpdateCmd
	GetScopes() []*string
	SetSidebar(v string) *NeuronAppVersionUpdateCmd
	GetSidebar() *string
	SetStatus(v string) *NeuronAppVersionUpdateCmd
	GetStatus() *string
	SetUnbindEffect(v string) *NeuronAppVersionUpdateCmd
	GetUnbindEffect() *string
	SetUnbindReasons(v []*string) *NeuronAppVersionUpdateCmd
	GetUnbindReasons() []*string
	SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersionUpdateCmd
	GetUpdatedFeature() []*NeuronAppInfoContent
}

type NeuronAppVersionUpdateCmd struct {
	AppEntry *string `json:"appEntry,omitempty" xml:"appEntry,omitempty"`
	// example:
	//
	// 升级订单功能
	Desc        *string                 `json:"desc,omitempty" xml:"desc,omitempty"`
	Envs        []*string               `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
	FeatureDesc []*NeuronAppInfoContent `json:"featureDesc,omitempty" xml:"featureDesc,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id             *int64    `json:"id,omitempty" xml:"id,omitempty"`
	ImageUrls      []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	InstructionUrl *string   `json:"instructionUrl,omitempty" xml:"instructionUrl,omitempty"`
	// example:
	//
	// NEURON
	ManageType *string `json:"manageType,omitempty" xml:"manageType,omitempty"`
	// example:
	//
	// 1
	MobiId      *int64    `json:"mobiId,omitempty" xml:"mobiId,omitempty"`
	Pbcs        []*int64  `json:"pbcs,omitempty" xml:"pbcs,omitempty" type:"Repeated"`
	PrivateInfo []*string `json:"privateInfo,omitempty" xml:"privateInfo,omitempty" type:"Repeated"`
	Scopes      []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Sidebar *string `json:"sidebar,omitempty" xml:"sidebar,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1.相关代码删除
	UnbindEffect   *string                 `json:"unbindEffect,omitempty" xml:"unbindEffect,omitempty"`
	UnbindReasons  []*string               `json:"unbindReasons,omitempty" xml:"unbindReasons,omitempty" type:"Repeated"`
	UpdatedFeature []*NeuronAppInfoContent `json:"updatedFeature,omitempty" xml:"updatedFeature,omitempty" type:"Repeated"`
}

func (s NeuronAppVersionUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppVersionUpdateCmd) GoString() string {
	return s.String()
}

func (s *NeuronAppVersionUpdateCmd) GetAppEntry() *string {
	return s.AppEntry
}

func (s *NeuronAppVersionUpdateCmd) GetDesc() *string {
	return s.Desc
}

func (s *NeuronAppVersionUpdateCmd) GetEnvs() []*string {
	return s.Envs
}

func (s *NeuronAppVersionUpdateCmd) GetFeatureDesc() []*NeuronAppInfoContent {
	return s.FeatureDesc
}

func (s *NeuronAppVersionUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *NeuronAppVersionUpdateCmd) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *NeuronAppVersionUpdateCmd) GetInstructionUrl() *string {
	return s.InstructionUrl
}

func (s *NeuronAppVersionUpdateCmd) GetManageType() *string {
	return s.ManageType
}

func (s *NeuronAppVersionUpdateCmd) GetMobiId() *int64 {
	return s.MobiId
}

func (s *NeuronAppVersionUpdateCmd) GetPbcs() []*int64 {
	return s.Pbcs
}

func (s *NeuronAppVersionUpdateCmd) GetPrivateInfo() []*string {
	return s.PrivateInfo
}

func (s *NeuronAppVersionUpdateCmd) GetScopes() []*string {
	return s.Scopes
}

func (s *NeuronAppVersionUpdateCmd) GetSidebar() *string {
	return s.Sidebar
}

func (s *NeuronAppVersionUpdateCmd) GetStatus() *string {
	return s.Status
}

func (s *NeuronAppVersionUpdateCmd) GetUnbindEffect() *string {
	return s.UnbindEffect
}

func (s *NeuronAppVersionUpdateCmd) GetUnbindReasons() []*string {
	return s.UnbindReasons
}

func (s *NeuronAppVersionUpdateCmd) GetUpdatedFeature() []*NeuronAppInfoContent {
	return s.UpdatedFeature
}

func (s *NeuronAppVersionUpdateCmd) SetAppEntry(v string) *NeuronAppVersionUpdateCmd {
	s.AppEntry = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetDesc(v string) *NeuronAppVersionUpdateCmd {
	s.Desc = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetEnvs(v []*string) *NeuronAppVersionUpdateCmd {
	s.Envs = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersionUpdateCmd {
	s.FeatureDesc = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetId(v int64) *NeuronAppVersionUpdateCmd {
	s.Id = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetImageUrls(v []*string) *NeuronAppVersionUpdateCmd {
	s.ImageUrls = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetInstructionUrl(v string) *NeuronAppVersionUpdateCmd {
	s.InstructionUrl = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetManageType(v string) *NeuronAppVersionUpdateCmd {
	s.ManageType = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetMobiId(v int64) *NeuronAppVersionUpdateCmd {
	s.MobiId = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetPbcs(v []*int64) *NeuronAppVersionUpdateCmd {
	s.Pbcs = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetPrivateInfo(v []*string) *NeuronAppVersionUpdateCmd {
	s.PrivateInfo = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetScopes(v []*string) *NeuronAppVersionUpdateCmd {
	s.Scopes = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetSidebar(v string) *NeuronAppVersionUpdateCmd {
	s.Sidebar = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetStatus(v string) *NeuronAppVersionUpdateCmd {
	s.Status = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetUnbindEffect(v string) *NeuronAppVersionUpdateCmd {
	s.UnbindEffect = &v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetUnbindReasons(v []*string) *NeuronAppVersionUpdateCmd {
	s.UnbindReasons = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersionUpdateCmd {
	s.UpdatedFeature = v
	return s
}

func (s *NeuronAppVersionUpdateCmd) Validate() error {
	if s.FeatureDesc != nil {
		for _, item := range s.FeatureDesc {
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
