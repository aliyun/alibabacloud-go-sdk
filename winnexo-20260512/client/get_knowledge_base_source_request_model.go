// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *GetKnowledgeBaseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *GetKnowledgeBaseSourceRequest
	GetTenantId() *string
}

type GetKnowledgeBaseSourceRequest struct {
	// The unique identifier on the business system side, that is, the business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 676577544219585
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetKnowledgeBaseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSourceRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetKnowledgeBaseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetKnowledgeBaseSourceRequest) SetSourceId(v string) *GetKnowledgeBaseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetKnowledgeBaseSourceRequest) SetTenantId(v string) *GetKnowledgeBaseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *GetKnowledgeBaseSourceRequest) Validate() error {
	return dara.Validate(s)
}
