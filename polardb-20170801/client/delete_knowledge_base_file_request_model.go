// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *DeleteKnowledgeBaseFileRequest
	GetFileId() *string
	SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseFileRequest
	GetKnowledgeBaseId() *string
	SetRegionId(v string) *DeleteKnowledgeBaseFileRequest
	GetRegionId() *string
}

type DeleteKnowledgeBaseFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
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
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKnowledgeBaseFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteKnowledgeBaseFileRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeBaseFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKnowledgeBaseFileRequest) SetFileId(v string) *DeleteKnowledgeBaseFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteKnowledgeBaseFileRequest) SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseFileRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeBaseFileRequest) SetRegionId(v string) *DeleteKnowledgeBaseFileRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKnowledgeBaseFileRequest) Validate() error {
	return dara.Validate(s)
}
