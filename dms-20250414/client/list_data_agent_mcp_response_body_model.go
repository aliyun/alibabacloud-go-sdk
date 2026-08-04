// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataAgentMcpResponseBodyData) *ListDataAgentMcpResponseBody
	GetData() *ListDataAgentMcpResponseBodyData
	SetErrorCode(v string) *ListDataAgentMcpResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentMcpResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataAgentMcpResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentMcpResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataAgentMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentMcpResponseBody
	GetSuccess() *bool
}

type ListDataAgentMcpResponseBody struct {
	// The paging query results of MCP Servers.
	Data *ListDataAgentMcpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return code. The value is success if the request was successful, or an error code if the request failed.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when a system-level request failure occurs.
	//
	// example:
	//
	// Failed to list MCP servers
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The maximum number of records returned in this response.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page. This value is empty when no more results are available.
	//
	// example:
	//
	// page-2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID, which is used to locate this call.
	//
	// example:
	//
	// 550e***544
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataAgentMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMcpResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentMcpResponseBody) GetData() *ListDataAgentMcpResponseBodyData {
	return s.Data
}

func (s *ListDataAgentMcpResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentMcpResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentMcpResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentMcpResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentMcpResponseBody) SetData(v *ListDataAgentMcpResponseBodyData) *ListDataAgentMcpResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentMcpResponseBody) SetErrorCode(v string) *ListDataAgentMcpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentMcpResponseBody) SetErrorMessage(v string) *ListDataAgentMcpResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentMcpResponseBody) SetMaxResults(v int32) *ListDataAgentMcpResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentMcpResponseBody) SetNextToken(v string) *ListDataAgentMcpResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentMcpResponseBody) SetRequestId(v string) *ListDataAgentMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentMcpResponseBody) SetSuccess(v bool) *ListDataAgentMcpResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAgentMcpResponseBodyData struct {
	// The list of MCP Servers on the current page. Each item contains information such as the service identifier, name, workspace, network, connection method, status, and enabled state.
	//
	// example:
	//
	// [{"uuid":"44lg***z65","name":"analytics","workspaceUuid":"	atvx***xmz","netType":"public","transportType":"sse","state":"ready","enable":true}]
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *float32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *float32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the filter conditions.
	//
	// example:
	//
	// 1
	TotalElements *float32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages that match the filter conditions.
	//
	// example:
	//
	// 1
	TotalPages *float32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDataAgentMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentMcpResponseBodyData) GetContent() interface{} {
	return s.Content
}

func (s *ListDataAgentMcpResponseBodyData) GetPageNumber() *float32 {
	return s.PageNumber
}

func (s *ListDataAgentMcpResponseBodyData) GetPageSize() *float32 {
	return s.PageSize
}

func (s *ListDataAgentMcpResponseBodyData) GetTotalElements() *float32 {
	return s.TotalElements
}

func (s *ListDataAgentMcpResponseBodyData) GetTotalPages() *float32 {
	return s.TotalPages
}

func (s *ListDataAgentMcpResponseBodyData) SetContent(v interface{}) *ListDataAgentMcpResponseBodyData {
	s.Content = v
	return s
}

func (s *ListDataAgentMcpResponseBodyData) SetPageNumber(v float32) *ListDataAgentMcpResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentMcpResponseBodyData) SetPageSize(v float32) *ListDataAgentMcpResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentMcpResponseBodyData) SetTotalElements(v float32) *ListDataAgentMcpResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDataAgentMcpResponseBodyData) SetTotalPages(v float32) *ListDataAgentMcpResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDataAgentMcpResponseBodyData) Validate() error {
	return dara.Validate(s)
}
