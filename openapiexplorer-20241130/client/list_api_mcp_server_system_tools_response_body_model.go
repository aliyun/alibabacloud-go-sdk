// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiMcpServerSystemToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListApiMcpServerSystemToolsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiMcpServerSystemToolsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApiMcpServerSystemToolsResponseBody
	GetRequestId() *string
	SetSystemTools(v []*ListApiMcpServerSystemToolsResponseBodySystemTools) *ListApiMcpServerSystemToolsResponseBody
	GetSystemTools() []*ListApiMcpServerSystemToolsResponseBodySystemTools
	SetTotalCount(v int32) *ListApiMcpServerSystemToolsResponseBody
	GetTotalCount() *int32
}

type ListApiMcpServerSystemToolsResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId   *string                                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SystemTools []*ListApiMcpServerSystemToolsResponseBodySystemTools `json:"systemTools,omitempty" xml:"systemTools,omitempty" type:"Repeated"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApiMcpServerSystemToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServerSystemToolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiMcpServerSystemToolsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiMcpServerSystemToolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiMcpServerSystemToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiMcpServerSystemToolsResponseBody) GetSystemTools() []*ListApiMcpServerSystemToolsResponseBodySystemTools {
	return s.SystemTools
}

func (s *ListApiMcpServerSystemToolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApiMcpServerSystemToolsResponseBody) SetMaxResults(v int32) *ListApiMcpServerSystemToolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApiMcpServerSystemToolsResponseBody) SetNextToken(v string) *ListApiMcpServerSystemToolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApiMcpServerSystemToolsResponseBody) SetRequestId(v string) *ListApiMcpServerSystemToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiMcpServerSystemToolsResponseBody) SetSystemTools(v []*ListApiMcpServerSystemToolsResponseBodySystemTools) *ListApiMcpServerSystemToolsResponseBody {
	s.SystemTools = v
	return s
}

func (s *ListApiMcpServerSystemToolsResponseBody) SetTotalCount(v int32) *ListApiMcpServerSystemToolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApiMcpServerSystemToolsResponseBody) Validate() error {
	if s.SystemTools != nil {
		for _, item := range s.SystemTools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiMcpServerSystemToolsResponseBodySystemTools struct {
	// example:
	//
	// system tool description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// mcp-system
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListApiMcpServerSystemToolsResponseBodySystemTools) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServerSystemToolsResponseBodySystemTools) GoString() string {
	return s.String()
}

func (s *ListApiMcpServerSystemToolsResponseBodySystemTools) GetDescription() *string {
	return s.Description
}

func (s *ListApiMcpServerSystemToolsResponseBodySystemTools) GetName() *string {
	return s.Name
}

func (s *ListApiMcpServerSystemToolsResponseBodySystemTools) SetDescription(v string) *ListApiMcpServerSystemToolsResponseBodySystemTools {
	s.Description = &v
	return s
}

func (s *ListApiMcpServerSystemToolsResponseBodySystemTools) SetName(v string) *ListApiMcpServerSystemToolsResponseBodySystemTools {
	s.Name = &v
	return s
}

func (s *ListApiMcpServerSystemToolsResponseBodySystemTools) Validate() error {
	return dara.Validate(s)
}
