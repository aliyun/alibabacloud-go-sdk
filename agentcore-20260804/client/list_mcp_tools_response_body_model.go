// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpToolsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMcpToolsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMcpToolsResponseBodyItems) *ListMcpToolsResponseBody
	GetItems() []*ListMcpToolsResponseBodyItems
	SetMaxResults(v int32) *ListMcpToolsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpToolsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpToolsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpToolsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpToolsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMcpToolsResponseBody
	GetTotalCount() *int64
}

type ListMcpToolsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list data.
	Items []*ListMcpToolsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of results per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// Request processed successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMcpToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpToolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpToolsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpToolsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMcpToolsResponseBody) GetItems() []*ListMcpToolsResponseBodyItems {
	return s.Items
}

func (s *ListMcpToolsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpToolsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpToolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpToolsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpToolsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcpToolsResponseBody) SetCode(v string) *ListMcpToolsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetHttpStatusCode(v int32) *ListMcpToolsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetItems(v []*ListMcpToolsResponseBodyItems) *ListMcpToolsResponseBody {
	s.Items = v
	return s
}

func (s *ListMcpToolsResponseBody) SetMaxResults(v int32) *ListMcpToolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetMessage(v string) *ListMcpToolsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetNextToken(v string) *ListMcpToolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetRequestId(v string) *ListMcpToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetSuccess(v bool) *ListMcpToolsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetTotalCount(v int64) *ListMcpToolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpToolsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpToolsResponseBodyItems struct {
	// The input parameter schema of the tool.
	//
	// example:
	//
	// {"type":"object"}
	InputSchema *string `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// The output parameter schema of the tool.
	//
	// example:
	//
	// {"type":"object"}
	OutputSchema *string `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
	// The MCP tool description.
	//
	// example:
	//
	// Get the current time
	ToolDescription *string `json:"toolDescription,omitempty" xml:"toolDescription,omitempty"`
	// The MCP tool name.
	//
	// example:
	//
	// get-current-time
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s ListMcpToolsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMcpToolsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMcpToolsResponseBodyItems) GetInputSchema() *string {
	return s.InputSchema
}

func (s *ListMcpToolsResponseBodyItems) GetOutputSchema() *string {
	return s.OutputSchema
}

func (s *ListMcpToolsResponseBodyItems) GetToolDescription() *string {
	return s.ToolDescription
}

func (s *ListMcpToolsResponseBodyItems) GetToolName() *string {
	return s.ToolName
}

func (s *ListMcpToolsResponseBodyItems) SetInputSchema(v string) *ListMcpToolsResponseBodyItems {
	s.InputSchema = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) SetOutputSchema(v string) *ListMcpToolsResponseBodyItems {
	s.OutputSchema = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) SetToolDescription(v string) *ListMcpToolsResponseBodyItems {
	s.ToolDescription = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) SetToolName(v string) *ListMcpToolsResponseBodyItems {
	s.ToolName = &v
	return s
}

func (s *ListMcpToolsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
