// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgePointsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFocusObjectsShrink(v string) *RetrieveKnowledgePointsShrinkRequest
	GetFocusObjectsShrink() *string
	SetKnowledgeBaseIdsShrink(v string) *RetrieveKnowledgePointsShrinkRequest
	GetKnowledgeBaseIdsShrink() *string
	SetMaxResults(v int32) *RetrieveKnowledgePointsShrinkRequest
	GetMaxResults() *int32
	SetQuery(v string) *RetrieveKnowledgePointsShrinkRequest
	GetQuery() *string
}

type RetrieveKnowledgePointsShrinkRequest struct {
	// The graph objects to focus the search on. You can specify up to 2 unique objects. The objects must be within the current authorized search scope and do not expand permissions.
	FocusObjectsShrink *string `json:"FocusObjects,omitempty" xml:"FocusObjects,omitempty"`
	// A list of 1 to 50 unique knowledge base IDs that restricts the search. If you omit this parameter, the operation discovers knowledge bases within the caller’s authorized scope. If you provide IDs, the operation checks access to each knowledge base and the entire request fails if the caller lacks access to any of them.
	KnowledgeBaseIdsShrink *string `json:"KnowledgeBaseIds,omitempty" xml:"KnowledgeBaseIds,omitempty"`
	// The maximum number of hits across all knowledge bases. Default value: 10. Valid values: 1 to 50. This parameter does not specify the quota for each knowledge base or the page size.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The natural language question. The length cannot exceed 256 Unicode code points after leading and trailing whitespaces are removed.
	//
	// This parameter is required.
	//
	// example:
	//
	// How is sales revenue calculated, and which tables does it come from?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s RetrieveKnowledgePointsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsShrinkRequest) GetFocusObjectsShrink() *string {
	return s.FocusObjectsShrink
}

func (s *RetrieveKnowledgePointsShrinkRequest) GetKnowledgeBaseIdsShrink() *string {
	return s.KnowledgeBaseIdsShrink
}

func (s *RetrieveKnowledgePointsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *RetrieveKnowledgePointsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveKnowledgePointsShrinkRequest) SetFocusObjectsShrink(v string) *RetrieveKnowledgePointsShrinkRequest {
	s.FocusObjectsShrink = &v
	return s
}

func (s *RetrieveKnowledgePointsShrinkRequest) SetKnowledgeBaseIdsShrink(v string) *RetrieveKnowledgePointsShrinkRequest {
	s.KnowledgeBaseIdsShrink = &v
	return s
}

func (s *RetrieveKnowledgePointsShrinkRequest) SetMaxResults(v int32) *RetrieveKnowledgePointsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *RetrieveKnowledgePointsShrinkRequest) SetQuery(v string) *RetrieveKnowledgePointsShrinkRequest {
	s.Query = &v
	return s
}

func (s *RetrieveKnowledgePointsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
