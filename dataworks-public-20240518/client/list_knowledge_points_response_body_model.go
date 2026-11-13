// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgePointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgePointsResponseBody
	GetCode() *string
	SetData(v *ListKnowledgePointsResponseBodyData) *ListKnowledgePointsResponseBody
	GetData() *ListKnowledgePointsResponseBodyData
	SetMaxResults(v int32) *ListKnowledgePointsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListKnowledgePointsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgePointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgePointsResponseBody
	GetSuccess() *bool
}

type ListKnowledgePointsResponseBody struct {
	// The business error code. This parameter is omitted when the request is successful. If the request fails, the common error response protocol is followed.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The summaries of the knowledge points on the current page and the cursor for the next page.
	Data *ListKnowledgePointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of knowledge bases to return per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The readable description of the business error. This parameter is omitted when the request is successful. When troubleshooting a failure, provide the RequestId at the same time.
	//
	// example:
	//
	// The request parameters are invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKnowledgePointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgePointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgePointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgePointsResponseBody) GetData() *ListKnowledgePointsResponseBodyData {
	return s.Data
}

func (s *ListKnowledgePointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgePointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgePointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgePointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgePointsResponseBody) SetCode(v string) *ListKnowledgePointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgePointsResponseBody) SetData(v *ListKnowledgePointsResponseBodyData) *ListKnowledgePointsResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgePointsResponseBody) SetMaxResults(v int32) *ListKnowledgePointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgePointsResponseBody) SetMessage(v string) *ListKnowledgePointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgePointsResponseBody) SetRequestId(v string) *ListKnowledgePointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgePointsResponseBody) SetSuccess(v bool) *ListKnowledgePointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgePointsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgePointsResponseBodyData struct {
	// The summaries of the knowledge points that the current caller has permission to read on the current page. The KnowledgeBaseId, Type, and Id of each item can be directly used in the GetKnowledgePoint operation.
	KnowledgePoints []*ListKnowledgePointsResponseBodyDataKnowledgePoints `json:"KnowledgePoints,omitempty" xml:"KnowledgePoints,omitempty" type:"Repeated"`
	// The pagination cursor. Do not specify this parameter for the first page. For subsequent pages, pass the NextToken value returned in the previous response as is, and keep the knowledge base, type, and page size unchanged. An empty cursor indicates that no more data is available. If the current page is empty but the cursor is not empty, continue querying.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgePointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgePointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgePointsResponseBodyData) GetKnowledgePoints() []*ListKnowledgePointsResponseBodyDataKnowledgePoints {
	return s.KnowledgePoints
}

func (s *ListKnowledgePointsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgePointsResponseBodyData) SetKnowledgePoints(v []*ListKnowledgePointsResponseBodyDataKnowledgePoints) *ListKnowledgePointsResponseBodyData {
	s.KnowledgePoints = v
	return s
}

func (s *ListKnowledgePointsResponseBodyData) SetNextToken(v string) *ListKnowledgePointsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgePointsResponseBodyData) Validate() error {
	if s.KnowledgePoints != nil {
		for _, item := range s.KnowledgePoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgePointsResponseBodyDataKnowledgePoints struct {
	// The read identifier of the current knowledge point. For GRAPH, it is the ID of the local graph entry node. For WIKI, it is the page ID. For KNOWLEDGE_CONTENT, it is the content ID. Use this value as is in the GetKnowledgePoint operation without parsing or concatenation.
	//
	// example:
	//
	// node-sales
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the knowledge base to which the knowledge point belongs. This ID is used together with Type and Id to locate the knowledge point.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The readable summary of the knowledge point. The summary is not guaranteed to contain the complete content. To obtain the details, call the GetKnowledgePoint operation.
	//
	// example:
	//
	// Sales metrics and their associated order tables
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The readable title of the knowledge point. For GRAPH, the name of the entry node is used.
	//
	// example:
	//
	// Sales
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The read category of the knowledge point. Valid values:
	//
	// - GRAPH: local graph.
	//
	// - WIKI: topic page.
	//
	// - KNOWLEDGE_CONTENT: knowledge content.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgePointsResponseBodyDataKnowledgePoints) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgePointsResponseBodyDataKnowledgePoints) GoString() string {
	return s.String()
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) GetId() *string {
	return s.Id
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) GetSummary() *string {
	return s.Summary
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) GetTitle() *string {
	return s.Title
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) GetType() *string {
	return s.Type
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) SetId(v string) *ListKnowledgePointsResponseBodyDataKnowledgePoints {
	s.Id = &v
	return s
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) SetKnowledgeBaseId(v string) *ListKnowledgePointsResponseBodyDataKnowledgePoints {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) SetSummary(v string) *ListKnowledgePointsResponseBodyDataKnowledgePoints {
	s.Summary = &v
	return s
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) SetTitle(v string) *ListKnowledgePointsResponseBodyDataKnowledgePoints {
	s.Title = &v
	return s
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) SetType(v string) *ListKnowledgePointsResponseBodyDataKnowledgePoints {
	s.Type = &v
	return s
}

func (s *ListKnowledgePointsResponseBodyDataKnowledgePoints) Validate() error {
	return dara.Validate(s)
}
