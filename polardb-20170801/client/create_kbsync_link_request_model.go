// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKBSyncLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *CreateKBSyncLinkRequest
	GetClientId() *string
	SetClientSecret(v string) *CreateKBSyncLinkRequest
	GetClientSecret() *string
	SetDescription(v string) *CreateKBSyncLinkRequest
	GetDescription() *string
	SetKnowledgeBaseId(v string) *CreateKBSyncLinkRequest
	GetKnowledgeBaseId() *string
	SetLinkName(v string) *CreateKBSyncLinkRequest
	GetLinkName() *string
	SetRegionId(v string) *CreateKBSyncLinkRequest
	GetRegionId() *string
	SetSourceDir(v string) *CreateKBSyncLinkRequest
	GetSourceDir() *string
	SetSourceType(v string) *CreateKBSyncLinkRequest
	GetSourceType() *string
	SetSyncIntervalMinutes(v int32) *CreateKBSyncLinkRequest
	GetSyncIntervalMinutes() *int32
	SetTenantId(v string) *CreateKBSyncLinkRequest
	GetTenantId() *string
}

type CreateKBSyncLinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cli_xxxxxxbe8
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ******
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// example:
	//
	// testDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testName
	LinkName *string `json:"LinkName,omitempty" xml:"LinkName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.feishu.cn/wiki/space/xxxxxx
	SourceDir *string `json:"SourceDir,omitempty" xml:"SourceDir,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FEISHU
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 30
	SyncIntervalMinutes *int32 `json:"SyncIntervalMinutes,omitempty" xml:"SyncIntervalMinutes,omitempty"`
	// example:
	//
	// 63eexxxx-xxxx-xxxx-xxxx-xxxxxx090f82
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateKBSyncLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKBSyncLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKBSyncLinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CreateKBSyncLinkRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateKBSyncLinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKBSyncLinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKBSyncLinkRequest) GetLinkName() *string {
	return s.LinkName
}

func (s *CreateKBSyncLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateKBSyncLinkRequest) GetSourceDir() *string {
	return s.SourceDir
}

func (s *CreateKBSyncLinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateKBSyncLinkRequest) GetSyncIntervalMinutes() *int32 {
	return s.SyncIntervalMinutes
}

func (s *CreateKBSyncLinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKBSyncLinkRequest) SetClientId(v string) *CreateKBSyncLinkRequest {
	s.ClientId = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetClientSecret(v string) *CreateKBSyncLinkRequest {
	s.ClientSecret = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetDescription(v string) *CreateKBSyncLinkRequest {
	s.Description = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetKnowledgeBaseId(v string) *CreateKBSyncLinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetLinkName(v string) *CreateKBSyncLinkRequest {
	s.LinkName = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetRegionId(v string) *CreateKBSyncLinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetSourceDir(v string) *CreateKBSyncLinkRequest {
	s.SourceDir = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetSourceType(v string) *CreateKBSyncLinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetSyncIntervalMinutes(v int32) *CreateKBSyncLinkRequest {
	s.SyncIntervalMinutes = &v
	return s
}

func (s *CreateKBSyncLinkRequest) SetTenantId(v string) *CreateKBSyncLinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKBSyncLinkRequest) Validate() error {
	return dara.Validate(s)
}
