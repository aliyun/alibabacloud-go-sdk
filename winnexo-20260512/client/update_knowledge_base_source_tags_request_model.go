// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetSourceId() *string
	SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetSourceTags() *string
	SetTenantId(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetTenantId() *string
}

type UpdateKnowledgeBaseSourceTagsRequest struct {
	// The unique identifier on the business system side, that is, the business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The resource tags. This is an optional parameter that accepts a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateKnowledgeBaseSourceTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceTagsRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.SourceTags = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetTenantId(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) Validate() error {
	return dara.Validate(s)
}
