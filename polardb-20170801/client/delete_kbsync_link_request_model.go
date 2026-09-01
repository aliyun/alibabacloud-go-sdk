// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKBSyncLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DeleteKBSyncLinkRequest
	GetKnowledgeBaseId() *string
	SetLinkId(v string) *DeleteKBSyncLinkRequest
	GetLinkId() *string
	SetRegionId(v string) *DeleteKBSyncLinkRequest
	GetRegionId() *string
}

type DeleteKBSyncLinkRequest struct {
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
}

func (s DeleteKBSyncLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKBSyncLinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteKBSyncLinkRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKBSyncLinkRequest) GetLinkId() *string {
	return s.LinkId
}

func (s *DeleteKBSyncLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKBSyncLinkRequest) SetKnowledgeBaseId(v string) *DeleteKBSyncLinkRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKBSyncLinkRequest) SetLinkId(v string) *DeleteKBSyncLinkRequest {
	s.LinkId = &v
	return s
}

func (s *DeleteKBSyncLinkRequest) SetRegionId(v string) *DeleteKBSyncLinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKBSyncLinkRequest) Validate() error {
	return dara.Validate(s)
}
