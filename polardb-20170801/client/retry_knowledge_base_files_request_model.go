// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileIds(v string) *RetryKnowledgeBaseFilesRequest
	GetFileIds() *string
	SetKnowledgeBaseId(v string) *RetryKnowledgeBaseFilesRequest
	GetKnowledgeBaseId() *string
	SetRegionId(v string) *RetryKnowledgeBaseFilesRequest
	GetRegionId() *string
}

type RetryKnowledgeBaseFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 91b97b71-xxxx-xxxx-xxxx-33c6a6341cdc
	FileIds *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
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

func (s RetryKnowledgeBaseFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFilesRequest) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFilesRequest) GetFileIds() *string {
	return s.FileIds
}

func (s *RetryKnowledgeBaseFilesRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetryKnowledgeBaseFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RetryKnowledgeBaseFilesRequest) SetFileIds(v string) *RetryKnowledgeBaseFilesRequest {
	s.FileIds = &v
	return s
}

func (s *RetryKnowledgeBaseFilesRequest) SetKnowledgeBaseId(v string) *RetryKnowledgeBaseFilesRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetryKnowledgeBaseFilesRequest) SetRegionId(v string) *RetryKnowledgeBaseFilesRequest {
	s.RegionId = &v
	return s
}

func (s *RetryKnowledgeBaseFilesRequest) Validate() error {
	return dara.Validate(s)
}
