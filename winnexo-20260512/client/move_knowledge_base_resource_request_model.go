// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveKnowledgeBaseResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeId(v string) *MoveKnowledgeBaseResourceRequest
	GetKnowledgeId() *string
	SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceRequest
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveKnowledgeBaseResourceRequest
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceRequest
	GetTargetDirectoryId() *string
	SetTenantId(v string) *MoveKnowledgeBaseResourceRequest
	GetTenantId() *string
}

type MoveKnowledgeBaseResourceRequest struct {
	// Not supported. This parameter is ignored.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// The source directory ID. This is the enterprise knowledge base directory where the resource currently resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceDirectoryId
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000627
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The target directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1729094555111072
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MoveKnowledgeBaseResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveKnowledgeBaseResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeBaseResourceRequest) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *MoveKnowledgeBaseResourceRequest) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveKnowledgeBaseResourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveKnowledgeBaseResourceRequest) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveKnowledgeBaseResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MoveKnowledgeBaseResourceRequest) SetKnowledgeId(v string) *MoveKnowledgeBaseResourceRequest {
	s.KnowledgeId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceRequest {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetSourceId(v string) *MoveKnowledgeBaseResourceRequest {
	s.SourceId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceRequest {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetTenantId(v string) *MoveKnowledgeBaseResourceRequest {
	s.TenantId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) Validate() error {
	return dara.Validate(s)
}
