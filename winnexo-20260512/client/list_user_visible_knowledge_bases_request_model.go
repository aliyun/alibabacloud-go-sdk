// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListUserVisibleKnowledgeBasesRequest
	GetKeyword() *string
	SetTenantId(v string) *ListUserVisibleKnowledgeBasesRequest
	GetTenantId() *string
}

type ListUserVisibleKnowledgeBasesRequest struct {
	// The keyword for fuzzy match on form component data.
	//
	// example:
	//
	// ProductKnowledge
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly in winnexo-cli by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUserVisibleKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetKeyword(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetTenantId(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.TenantId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
