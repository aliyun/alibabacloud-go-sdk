// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKBSyncLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *UpdateKBSyncLinkRequest
	GetKnowledgeBaseId() *string
	SetLinkId(v string) *UpdateKBSyncLinkRequest
	GetLinkId() *string
	SetRegionId(v string) *UpdateKBSyncLinkRequest
	GetRegionId() *string
	SetSyncIntervalMinutes(v int32) *UpdateKBSyncLinkRequest
	GetSyncIntervalMinutes() *int32
}

type UpdateKBSyncLinkRequest struct {
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
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 20
	SyncIntervalMinutes *int32 `json:"SyncIntervalMinutes,omitempty" xml:"SyncIntervalMinutes,omitempty"`
}

func (s UpdateKBSyncLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKBSyncLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKBSyncLinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKBSyncLinkRequest) GetLinkId() *string {
	return s.LinkId
}

func (s *UpdateKBSyncLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKBSyncLinkRequest) GetSyncIntervalMinutes() *int32 {
	return s.SyncIntervalMinutes
}

func (s *UpdateKBSyncLinkRequest) SetKnowledgeBaseId(v string) *UpdateKBSyncLinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetLinkId(v string) *UpdateKBSyncLinkRequest {
	s.LinkId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetRegionId(v string) *UpdateKBSyncLinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) SetSyncIntervalMinutes(v int32) *UpdateKBSyncLinkRequest {
	s.SyncIntervalMinutes = &v
	return s
}

func (s *UpdateKBSyncLinkRequest) Validate() error {
	return dara.Validate(s)
}
