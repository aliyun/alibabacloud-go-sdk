// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterKnowledgeBaseFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *RegisterKnowledgeBaseFileRequest
	GetFilePath() *string
	SetKnowledgeBaseId(v string) *RegisterKnowledgeBaseFileRequest
	GetKnowledgeBaseId() *string
	SetRegionId(v string) *RegisterKnowledgeBaseFileRequest
	GetRegionId() *string
}

type RegisterKnowledgeBaseFileRequest struct {
	// The OSS object key of the uploaded file, excluding the oss://BucketName/ prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// pks-2ze123456789abcd/pkb-2ze123456789abcd/example.pdf
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-2ze123456789abcd
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the region where the knowledge base resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegisterKnowledgeBaseFileRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterKnowledgeBaseFileRequest) GoString() string {
	return s.String()
}

func (s *RegisterKnowledgeBaseFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *RegisterKnowledgeBaseFileRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RegisterKnowledgeBaseFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RegisterKnowledgeBaseFileRequest) SetFilePath(v string) *RegisterKnowledgeBaseFileRequest {
	s.FilePath = &v
	return s
}

func (s *RegisterKnowledgeBaseFileRequest) SetKnowledgeBaseId(v string) *RegisterKnowledgeBaseFileRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RegisterKnowledgeBaseFileRequest) SetRegionId(v string) *RegisterKnowledgeBaseFileRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterKnowledgeBaseFileRequest) Validate() error {
	return dara.Validate(s)
}
