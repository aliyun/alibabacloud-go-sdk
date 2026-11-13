// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgePointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFocusObjects(v []*RetrieveKnowledgePointsRequestFocusObjects) *RetrieveKnowledgePointsRequest
	GetFocusObjects() []*RetrieveKnowledgePointsRequestFocusObjects
	SetKnowledgeBaseIds(v []*string) *RetrieveKnowledgePointsRequest
	GetKnowledgeBaseIds() []*string
	SetMaxResults(v int32) *RetrieveKnowledgePointsRequest
	GetMaxResults() *int32
	SetQuery(v string) *RetrieveKnowledgePointsRequest
	GetQuery() *string
}

type RetrieveKnowledgePointsRequest struct {
	// The graph objects to focus the search on. You can specify up to 2 unique objects. The objects must be within the current authorized search scope and do not expand permissions.
	FocusObjects []*RetrieveKnowledgePointsRequestFocusObjects `json:"FocusObjects,omitempty" xml:"FocusObjects,omitempty" type:"Repeated"`
	// A list of 1 to 50 unique knowledge base IDs that restricts the search. If you omit this parameter, the operation discovers knowledge bases within the caller’s authorized scope. If you provide IDs, the operation checks access to each knowledge base and the entire request fails if the caller lacks access to any of them.
	KnowledgeBaseIds []*string `json:"KnowledgeBaseIds,omitempty" xml:"KnowledgeBaseIds,omitempty" type:"Repeated"`
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

func (s RetrieveKnowledgePointsRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsRequest) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsRequest) GetFocusObjects() []*RetrieveKnowledgePointsRequestFocusObjects {
	return s.FocusObjects
}

func (s *RetrieveKnowledgePointsRequest) GetKnowledgeBaseIds() []*string {
	return s.KnowledgeBaseIds
}

func (s *RetrieveKnowledgePointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *RetrieveKnowledgePointsRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveKnowledgePointsRequest) SetFocusObjects(v []*RetrieveKnowledgePointsRequestFocusObjects) *RetrieveKnowledgePointsRequest {
	s.FocusObjects = v
	return s
}

func (s *RetrieveKnowledgePointsRequest) SetKnowledgeBaseIds(v []*string) *RetrieveKnowledgePointsRequest {
	s.KnowledgeBaseIds = v
	return s
}

func (s *RetrieveKnowledgePointsRequest) SetMaxResults(v int32) *RetrieveKnowledgePointsRequest {
	s.MaxResults = &v
	return s
}

func (s *RetrieveKnowledgePointsRequest) SetQuery(v string) *RetrieveKnowledgePointsRequest {
	s.Query = &v
	return s
}

func (s *RetrieveKnowledgePointsRequest) Validate() error {
	if s.FocusObjects != nil {
		for _, item := range s.FocusObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RetrieveKnowledgePointsRequestFocusObjects struct {
	// The location ID of the knowledge point. Copy this value exactly as it appears in the search results and use it together with the corresponding KnowledgeBaseId and Type. Do not parse or convert this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// node-sales
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the knowledge base to which the knowledge point belongs. This parameter is used together with Type and Id to locate the knowledge point.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The read category of the knowledge point. Use GRAPH for the current focus objects and graph nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RetrieveKnowledgePointsRequestFocusObjects) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsRequestFocusObjects) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsRequestFocusObjects) GetId() *string {
	return s.Id
}

func (s *RetrieveKnowledgePointsRequestFocusObjects) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetrieveKnowledgePointsRequestFocusObjects) GetType() *string {
	return s.Type
}

func (s *RetrieveKnowledgePointsRequestFocusObjects) SetId(v string) *RetrieveKnowledgePointsRequestFocusObjects {
	s.Id = &v
	return s
}

func (s *RetrieveKnowledgePointsRequestFocusObjects) SetKnowledgeBaseId(v string) *RetrieveKnowledgePointsRequestFocusObjects {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetrieveKnowledgePointsRequestFocusObjects) SetType(v string) *RetrieveKnowledgePointsRequestFocusObjects {
	s.Type = &v
	return s
}

func (s *RetrieveKnowledgePointsRequestFocusObjects) Validate() error {
	return dara.Validate(s)
}
