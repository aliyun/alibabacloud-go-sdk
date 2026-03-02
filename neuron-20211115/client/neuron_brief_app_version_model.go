// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronBriefAppVersion interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *NeuronBriefAppVersion
	GetAccountId() *string
	SetAppId(v int64) *NeuronBriefAppVersion
	GetAppId() *int64
	SetEnterpriseId(v int64) *NeuronBriefAppVersion
	GetEnterpriseId() *int64
	SetGmtCreate(v string) *NeuronBriefAppVersion
	GetGmtCreate() *string
	SetGmtModified(v string) *NeuronBriefAppVersion
	GetGmtModified() *string
	SetId(v int64) *NeuronBriefAppVersion
	GetId() *int64
	SetManageType(v string) *NeuronBriefAppVersion
	GetManageType() *string
	SetMobiCommitId(v string) *NeuronBriefAppVersion
	GetMobiCommitId() *string
	SetMobiId(v int64) *NeuronBriefAppVersion
	GetMobiId() *int64
	SetMobiModuleId(v string) *NeuronBriefAppVersion
	GetMobiModuleId() *string
	SetMobiUrl(v string) *NeuronBriefAppVersion
	GetMobiUrl() *string
	SetProductId(v int64) *NeuronBriefAppVersion
	GetProductId() *int64
	SetStatus(v string) *NeuronBriefAppVersion
	GetStatus() *string
	SetVersion(v string) *NeuronBriefAppVersion
	GetVersion() *string
}

type NeuronBriefAppVersion struct {
	// example:
	//
	// 1343424
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 1
	EnterpriseId *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	GmtCreate    *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
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
	MobiUrl *string `json:"mobiUrl,omitempty" xml:"mobiUrl,omitempty"`
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s NeuronBriefAppVersion) String() string {
	return dara.Prettify(s)
}

func (s NeuronBriefAppVersion) GoString() string {
	return s.String()
}

func (s *NeuronBriefAppVersion) GetAccountId() *string {
	return s.AccountId
}

func (s *NeuronBriefAppVersion) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronBriefAppVersion) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *NeuronBriefAppVersion) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *NeuronBriefAppVersion) GetGmtModified() *string {
	return s.GmtModified
}

func (s *NeuronBriefAppVersion) GetId() *int64 {
	return s.Id
}

func (s *NeuronBriefAppVersion) GetManageType() *string {
	return s.ManageType
}

func (s *NeuronBriefAppVersion) GetMobiCommitId() *string {
	return s.MobiCommitId
}

func (s *NeuronBriefAppVersion) GetMobiId() *int64 {
	return s.MobiId
}

func (s *NeuronBriefAppVersion) GetMobiModuleId() *string {
	return s.MobiModuleId
}

func (s *NeuronBriefAppVersion) GetMobiUrl() *string {
	return s.MobiUrl
}

func (s *NeuronBriefAppVersion) GetProductId() *int64 {
	return s.ProductId
}

func (s *NeuronBriefAppVersion) GetStatus() *string {
	return s.Status
}

func (s *NeuronBriefAppVersion) GetVersion() *string {
	return s.Version
}

func (s *NeuronBriefAppVersion) SetAccountId(v string) *NeuronBriefAppVersion {
	s.AccountId = &v
	return s
}

func (s *NeuronBriefAppVersion) SetAppId(v int64) *NeuronBriefAppVersion {
	s.AppId = &v
	return s
}

func (s *NeuronBriefAppVersion) SetEnterpriseId(v int64) *NeuronBriefAppVersion {
	s.EnterpriseId = &v
	return s
}

func (s *NeuronBriefAppVersion) SetGmtCreate(v string) *NeuronBriefAppVersion {
	s.GmtCreate = &v
	return s
}

func (s *NeuronBriefAppVersion) SetGmtModified(v string) *NeuronBriefAppVersion {
	s.GmtModified = &v
	return s
}

func (s *NeuronBriefAppVersion) SetId(v int64) *NeuronBriefAppVersion {
	s.Id = &v
	return s
}

func (s *NeuronBriefAppVersion) SetManageType(v string) *NeuronBriefAppVersion {
	s.ManageType = &v
	return s
}

func (s *NeuronBriefAppVersion) SetMobiCommitId(v string) *NeuronBriefAppVersion {
	s.MobiCommitId = &v
	return s
}

func (s *NeuronBriefAppVersion) SetMobiId(v int64) *NeuronBriefAppVersion {
	s.MobiId = &v
	return s
}

func (s *NeuronBriefAppVersion) SetMobiModuleId(v string) *NeuronBriefAppVersion {
	s.MobiModuleId = &v
	return s
}

func (s *NeuronBriefAppVersion) SetMobiUrl(v string) *NeuronBriefAppVersion {
	s.MobiUrl = &v
	return s
}

func (s *NeuronBriefAppVersion) SetProductId(v int64) *NeuronBriefAppVersion {
	s.ProductId = &v
	return s
}

func (s *NeuronBriefAppVersion) SetStatus(v string) *NeuronBriefAppVersion {
	s.Status = &v
	return s
}

func (s *NeuronBriefAppVersion) SetVersion(v string) *NeuronBriefAppVersion {
	s.Version = &v
	return s
}

func (s *NeuronBriefAppVersion) Validate() error {
	return dara.Validate(s)
}
