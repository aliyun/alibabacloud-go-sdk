// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppVersion interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *NeuronAppVersion
	GetAccountId() *string
	SetAppEntry(v string) *NeuronAppVersion
	GetAppEntry() *string
	SetAppId(v int64) *NeuronAppVersion
	GetAppId() *int64
	SetDescription(v string) *NeuronAppVersion
	GetDescription() *string
	SetEnterpriseId(v int64) *NeuronAppVersion
	GetEnterpriseId() *int64
	SetEnvs(v []*string) *NeuronAppVersion
	GetEnvs() []*string
	SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersion
	GetFeatureDesc() []*NeuronAppInfoContent
	SetGmtCreate(v string) *NeuronAppVersion
	GetGmtCreate() *string
	SetGmtModified(v string) *NeuronAppVersion
	GetGmtModified() *string
	SetId(v int64) *NeuronAppVersion
	GetId() *int64
	SetImageUrls(v []*string) *NeuronAppVersion
	GetImageUrls() []*string
	SetInstructionUrl(v string) *NeuronAppVersion
	GetInstructionUrl() *string
	SetIsLatest(v int32) *NeuronAppVersion
	GetIsLatest() *int32
	SetManageType(v string) *NeuronAppVersion
	GetManageType() *string
	SetMobiCommitId(v string) *NeuronAppVersion
	GetMobiCommitId() *string
	SetMobiId(v int64) *NeuronAppVersion
	GetMobiId() *int64
	SetMobiModuleId(v string) *NeuronAppVersion
	GetMobiModuleId() *string
	SetMobiUrl(v string) *NeuronAppVersion
	GetMobiUrl() *string
	SetPbcs(v []*int64) *NeuronAppVersion
	GetPbcs() []*int64
	SetPrivateInfo(v []*string) *NeuronAppVersion
	GetPrivateInfo() []*string
	SetProductId(v int64) *NeuronAppVersion
	GetProductId() *int64
	SetScopes(v []*string) *NeuronAppVersion
	GetScopes() []*string
	SetSidebar(v string) *NeuronAppVersion
	GetSidebar() *string
	SetStatus(v string) *NeuronAppVersion
	GetStatus() *string
	SetUnbindEffect(v string) *NeuronAppVersion
	GetUnbindEffect() *string
	SetUnbindReasons(v []*string) *NeuronAppVersion
	GetUnbindReasons() []*string
	SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersion
	GetUpdatedFeature() []*NeuronAppInfoContent
	SetVersion(v string) *NeuronAppVersion
	GetVersion() *string
}

type NeuronAppVersion struct {
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
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64                  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Envs         []*string               `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
	FeatureDesc  []*NeuronAppInfoContent `json:"featureDesc,omitempty" xml:"featureDesc,omitempty" type:"Repeated"`
	GmtCreate    *string                 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string                 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
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
	// 1
	IsLatest   *int32  `json:"isLatest,omitempty" xml:"isLatest,omitempty"`
	ManageType *string `json:"manageType,omitempty" xml:"manageType,omitempty"`
	// example:
	//
	// commit_pckesd7d_20230227215101
	MobiCommitId *string `json:"mobiCommitId,omitempty" xml:"mobiCommitId,omitempty"`
	// example:
	//
	// 1
	MobiId *int64 `json:"mobiId,omitempty" xml:"mobiId,omitempty"`
	// example:
	//
	// module_tvtpydeq
	MobiModuleId *string `json:"mobiModuleId,omitempty" xml:"mobiModuleId,omitempty"`
	// example:
	//
	// sda
	MobiUrl     *string   `json:"mobiUrl,omitempty" xml:"mobiUrl,omitempty"`
	Pbcs        []*int64  `json:"pbcs,omitempty" xml:"pbcs,omitempty" type:"Repeated"`
	PrivateInfo []*string `json:"privateInfo,omitempty" xml:"privateInfo,omitempty" type:"Repeated"`
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
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s NeuronAppVersion) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppVersion) GoString() string {
	return s.String()
}

func (s *NeuronAppVersion) GetAccountId() *string {
	return s.AccountId
}

func (s *NeuronAppVersion) GetAppEntry() *string {
	return s.AppEntry
}

func (s *NeuronAppVersion) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronAppVersion) GetDescription() *string {
	return s.Description
}

func (s *NeuronAppVersion) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *NeuronAppVersion) GetEnvs() []*string {
	return s.Envs
}

func (s *NeuronAppVersion) GetFeatureDesc() []*NeuronAppInfoContent {
	return s.FeatureDesc
}

func (s *NeuronAppVersion) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *NeuronAppVersion) GetGmtModified() *string {
	return s.GmtModified
}

func (s *NeuronAppVersion) GetId() *int64 {
	return s.Id
}

func (s *NeuronAppVersion) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *NeuronAppVersion) GetInstructionUrl() *string {
	return s.InstructionUrl
}

func (s *NeuronAppVersion) GetIsLatest() *int32 {
	return s.IsLatest
}

func (s *NeuronAppVersion) GetManageType() *string {
	return s.ManageType
}

func (s *NeuronAppVersion) GetMobiCommitId() *string {
	return s.MobiCommitId
}

func (s *NeuronAppVersion) GetMobiId() *int64 {
	return s.MobiId
}

func (s *NeuronAppVersion) GetMobiModuleId() *string {
	return s.MobiModuleId
}

func (s *NeuronAppVersion) GetMobiUrl() *string {
	return s.MobiUrl
}

func (s *NeuronAppVersion) GetPbcs() []*int64 {
	return s.Pbcs
}

func (s *NeuronAppVersion) GetPrivateInfo() []*string {
	return s.PrivateInfo
}

func (s *NeuronAppVersion) GetProductId() *int64 {
	return s.ProductId
}

func (s *NeuronAppVersion) GetScopes() []*string {
	return s.Scopes
}

func (s *NeuronAppVersion) GetSidebar() *string {
	return s.Sidebar
}

func (s *NeuronAppVersion) GetStatus() *string {
	return s.Status
}

func (s *NeuronAppVersion) GetUnbindEffect() *string {
	return s.UnbindEffect
}

func (s *NeuronAppVersion) GetUnbindReasons() []*string {
	return s.UnbindReasons
}

func (s *NeuronAppVersion) GetUpdatedFeature() []*NeuronAppInfoContent {
	return s.UpdatedFeature
}

func (s *NeuronAppVersion) GetVersion() *string {
	return s.Version
}

func (s *NeuronAppVersion) SetAccountId(v string) *NeuronAppVersion {
	s.AccountId = &v
	return s
}

func (s *NeuronAppVersion) SetAppEntry(v string) *NeuronAppVersion {
	s.AppEntry = &v
	return s
}

func (s *NeuronAppVersion) SetAppId(v int64) *NeuronAppVersion {
	s.AppId = &v
	return s
}

func (s *NeuronAppVersion) SetDescription(v string) *NeuronAppVersion {
	s.Description = &v
	return s
}

func (s *NeuronAppVersion) SetEnterpriseId(v int64) *NeuronAppVersion {
	s.EnterpriseId = &v
	return s
}

func (s *NeuronAppVersion) SetEnvs(v []*string) *NeuronAppVersion {
	s.Envs = v
	return s
}

func (s *NeuronAppVersion) SetFeatureDesc(v []*NeuronAppInfoContent) *NeuronAppVersion {
	s.FeatureDesc = v
	return s
}

func (s *NeuronAppVersion) SetGmtCreate(v string) *NeuronAppVersion {
	s.GmtCreate = &v
	return s
}

func (s *NeuronAppVersion) SetGmtModified(v string) *NeuronAppVersion {
	s.GmtModified = &v
	return s
}

func (s *NeuronAppVersion) SetId(v int64) *NeuronAppVersion {
	s.Id = &v
	return s
}

func (s *NeuronAppVersion) SetImageUrls(v []*string) *NeuronAppVersion {
	s.ImageUrls = v
	return s
}

func (s *NeuronAppVersion) SetInstructionUrl(v string) *NeuronAppVersion {
	s.InstructionUrl = &v
	return s
}

func (s *NeuronAppVersion) SetIsLatest(v int32) *NeuronAppVersion {
	s.IsLatest = &v
	return s
}

func (s *NeuronAppVersion) SetManageType(v string) *NeuronAppVersion {
	s.ManageType = &v
	return s
}

func (s *NeuronAppVersion) SetMobiCommitId(v string) *NeuronAppVersion {
	s.MobiCommitId = &v
	return s
}

func (s *NeuronAppVersion) SetMobiId(v int64) *NeuronAppVersion {
	s.MobiId = &v
	return s
}

func (s *NeuronAppVersion) SetMobiModuleId(v string) *NeuronAppVersion {
	s.MobiModuleId = &v
	return s
}

func (s *NeuronAppVersion) SetMobiUrl(v string) *NeuronAppVersion {
	s.MobiUrl = &v
	return s
}

func (s *NeuronAppVersion) SetPbcs(v []*int64) *NeuronAppVersion {
	s.Pbcs = v
	return s
}

func (s *NeuronAppVersion) SetPrivateInfo(v []*string) *NeuronAppVersion {
	s.PrivateInfo = v
	return s
}

func (s *NeuronAppVersion) SetProductId(v int64) *NeuronAppVersion {
	s.ProductId = &v
	return s
}

func (s *NeuronAppVersion) SetScopes(v []*string) *NeuronAppVersion {
	s.Scopes = v
	return s
}

func (s *NeuronAppVersion) SetSidebar(v string) *NeuronAppVersion {
	s.Sidebar = &v
	return s
}

func (s *NeuronAppVersion) SetStatus(v string) *NeuronAppVersion {
	s.Status = &v
	return s
}

func (s *NeuronAppVersion) SetUnbindEffect(v string) *NeuronAppVersion {
	s.UnbindEffect = &v
	return s
}

func (s *NeuronAppVersion) SetUnbindReasons(v []*string) *NeuronAppVersion {
	s.UnbindReasons = v
	return s
}

func (s *NeuronAppVersion) SetUpdatedFeature(v []*NeuronAppInfoContent) *NeuronAppVersion {
	s.UpdatedFeature = v
	return s
}

func (s *NeuronAppVersion) SetVersion(v string) *NeuronAppVersion {
	s.Version = &v
	return s
}

func (s *NeuronAppVersion) Validate() error {
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
