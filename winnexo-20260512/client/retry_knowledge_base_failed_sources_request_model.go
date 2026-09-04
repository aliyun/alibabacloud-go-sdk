// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseFailedSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *RetryKnowledgeBaseFailedSourcesRequest
	GetDirectoryId() *string
	SetTenantId(v string) *RetryKnowledgeBaseFailedSourcesRequest
	GetTenantId() *string
}

type RetryKnowledgeBaseFailedSourcesRequest struct {
	// The enterprise knowledge base directory ID (recursively includes failed resources in subdirectories).
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RetryKnowledgeBaseFailedSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFailedSourcesRequest) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) SetDirectoryId(v string) *RetryKnowledgeBaseFailedSourcesRequest {
	s.DirectoryId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) SetTenantId(v string) *RetryKnowledgeBaseFailedSourcesRequest {
	s.TenantId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) Validate() error {
	return dara.Validate(s)
}
