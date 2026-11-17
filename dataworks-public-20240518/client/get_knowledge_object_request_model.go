// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *GetKnowledgeObjectRequest
	GetKnowledgeBaseId() *string
	SetKnowledgeObjectId(v string) *GetKnowledgeObjectRequest
	GetKnowledgeObjectId() *string
	SetVersionNo(v int64) *GetKnowledgeObjectRequest
	GetVersionNo() *int64
}

type GetKnowledgeObjectRequest struct {
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the long-term business object.
	//
	// This parameter is required.
	//
	// example:
	//
	// obj-001
	KnowledgeObjectId *string `json:"KnowledgeObjectId,omitempty" xml:"KnowledgeObjectId,omitempty"`
	// Applies only to semantic models. Specifies the version to read. If omitted, the operation returns the published version if available, or the latest version otherwise.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s GetKnowledgeObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeObjectRequest) GetKnowledgeObjectId() *string {
	return s.KnowledgeObjectId
}

func (s *GetKnowledgeObjectRequest) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *GetKnowledgeObjectRequest) SetKnowledgeBaseId(v string) *GetKnowledgeObjectRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeObjectRequest) SetKnowledgeObjectId(v string) *GetKnowledgeObjectRequest {
	s.KnowledgeObjectId = &v
	return s
}

func (s *GetKnowledgeObjectRequest) SetVersionNo(v int64) *GetKnowledgeObjectRequest {
	s.VersionNo = &v
	return s
}

func (s *GetKnowledgeObjectRequest) Validate() error {
	return dara.Validate(s)
}
