// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKBSyncLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *UpdateKBSyncLinkRequest
	GetClientId() *string
	SetClientSecret(v string) *UpdateKBSyncLinkRequest
	GetClientSecret() *string
	SetKnowledgeBaseId(v string) *UpdateKBSyncLinkRequest
	GetKnowledgeBaseId() *string
	SetLinkId(v string) *UpdateKBSyncLinkRequest
	GetLinkId() *string
	SetMcpEndpoint(v string) *UpdateKBSyncLinkRequest
	GetMcpEndpoint() *string
	SetRegionId(v string) *UpdateKBSyncLinkRequest
	GetRegionId() *string
	SetSheetMcpEndpoint(v string) *UpdateKBSyncLinkRequest
	GetSheetMcpEndpoint() *string
	SetSyncEnabled(v bool) *UpdateKBSyncLinkRequest
	GetSyncEnabled() *bool
	SetSyncIntervalMinutes(v int32) *UpdateKBSyncLinkRequest
	GetSyncIntervalMinutes() *int32
	SetUserId(v string) *UpdateKBSyncLinkRequest
	GetUserId() *string
}

type UpdateKBSyncLinkRequest struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
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
	// pkbl-xxxxx
	LinkId      *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	McpEndpoint *string `json:"McpEndpoint,omitempty" xml:"McpEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SheetMcpEndpoint *string `json:"SheetMcpEndpoint,omitempty" xml:"SheetMcpEndpoint,omitempty"`
	SyncEnabled      *bool   `json:"SyncEnabled,omitempty" xml:"SyncEnabled,omitempty"`
	// example:
	//
	// 20
	SyncIntervalMinutes *int32  `json:"SyncIntervalMinutes,omitempty" xml:"SyncIntervalMinutes,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateKBSyncLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKBSyncLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKBSyncLinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateKBSyncLinkRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateKBSyncLinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKBSyncLinkRequest) GetLinkId() *string {
	return s.LinkId
}

func (s *UpdateKBSyncLinkRequest) GetMcpEndpoint() *string {
	return s.McpEndpoint
}

func (s *UpdateKBSyncLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKBSyncLinkRequest) GetSheetMcpEndpoint() *string {
	return s.SheetMcpEndpoint
}

func (s *UpdateKBSyncLinkRequest) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateKBSyncLinkRequest) GetSyncIntervalMinutes() *int32 {
	return s.SyncIntervalMinutes
}

func (s *UpdateKBSyncLinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateKBSyncLinkRequest) SetClientId(v string) *UpdateKBSyncLinkRequest {
	s.ClientId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetClientSecret(v string) *UpdateKBSyncLinkRequest {
	s.ClientSecret = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetKnowledgeBaseId(v string) *UpdateKBSyncLinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetLinkId(v string) *UpdateKBSyncLinkRequest {
	s.LinkId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetMcpEndpoint(v string) *UpdateKBSyncLinkRequest {
	s.McpEndpoint = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetRegionId(v string) *UpdateKBSyncLinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetSheetMcpEndpoint(v string) *UpdateKBSyncLinkRequest {
	s.SheetMcpEndpoint = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetSyncEnabled(v bool) *UpdateKBSyncLinkRequest {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetSyncIntervalMinutes(v int32) *UpdateKBSyncLinkRequest {
	s.SyncIntervalMinutes = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetUserId(v string) *UpdateKBSyncLinkRequest {
	s.UserId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) Validate() error {
	return dara.Validate(s)
}
