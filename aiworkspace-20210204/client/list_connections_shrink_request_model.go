// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListConnectionsShrinkRequest
	GetAccessibility() *string
	SetConnectionIdsShrink(v string) *ListConnectionsShrinkRequest
	GetConnectionIdsShrink() *string
	SetConnectionName(v string) *ListConnectionsShrinkRequest
	GetConnectionName() *string
	SetConnectionTypesShrink(v string) *ListConnectionsShrinkRequest
	GetConnectionTypesShrink() *string
	SetCreator(v string) *ListConnectionsShrinkRequest
	GetCreator() *string
	SetEncryptOption(v string) *ListConnectionsShrinkRequest
	GetEncryptOption() *string
	SetMaxResults(v int32) *ListConnectionsShrinkRequest
	GetMaxResults() *int32
	SetModel(v string) *ListConnectionsShrinkRequest
	GetModel() *string
	SetModelTypesShrink(v string) *ListConnectionsShrinkRequest
	GetModelTypesShrink() *string
	SetNextToken(v string) *ListConnectionsShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListConnectionsShrinkRequest
	GetOrder() *string
	SetSortBy(v string) *ListConnectionsShrinkRequest
	GetSortBy() *string
	SetToolCall(v bool) *ListConnectionsShrinkRequest
	GetToolCall() *bool
	SetWorkspaceId(v string) *ListConnectionsShrinkRequest
	GetWorkspaceId() *string
}

type ListConnectionsShrinkRequest struct {
	// Visibility of the connection. Valid values:
	//
	// - PUBLIC: visible to all workspace members.
	//
	// - PRIVATE: visible only to the creator.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// List of connection IDs to filter by.
	ConnectionIdsShrink *string `json:"ConnectionIds,omitempty" xml:"ConnectionIds,omitempty"`
	// Connection name. Supports fuzzy matching.
	//
	// example:
	//
	// Database connection
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// List of connection types to filter by.
	ConnectionTypesShrink *string `json:"ConnectionTypes,omitempty" xml:"ConnectionTypes,omitempty"`
	// Alibaba Cloud account ID of the creator.
	//
	// example:
	//
	// 12908*******3242
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Encryption option for sensitive fields in the response. Valid values:
	//
	// - PlainText: returns values in plaintext.
	//
	// - Secret: returns values in ciphertext.
	//
	// example:
	//
	// PlainText
	EncryptOption *string `json:"EncryptOption,omitempty" xml:"EncryptOption,omitempty"`
	// Maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Model identifier. Filters connections associated with this model.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// List of model types to filter by.
	ModelTypesShrink *string `json:"ModelTypes,omitempty" xml:"ModelTypes,omitempty"`
	// The token that marks the starting position for the next page of results.
	//
	// example:
	//
	// 15
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Sort order. Use with the SortBy parameter. Valid values:
	//
	// - ASC: ascending order.
	//
	// - DESC (default): descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Field by which to sort results. Default value: GmtCreateTime. Valid values:
	//
	// - GmtCreateTime (default): sorts by creation time.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Whether tool calling is supported. Valid values:
	//
	// - true: supported.
	//
	// - false: not supported.
	//
	// example:
	//
	// true
	ToolCall *bool `json:"ToolCall,omitempty" xml:"ToolCall,omitempty"`
	// Workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListConnectionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsShrinkRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListConnectionsShrinkRequest) GetConnectionIdsShrink() *string {
	return s.ConnectionIdsShrink
}

func (s *ListConnectionsShrinkRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *ListConnectionsShrinkRequest) GetConnectionTypesShrink() *string {
	return s.ConnectionTypesShrink
}

func (s *ListConnectionsShrinkRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListConnectionsShrinkRequest) GetEncryptOption() *string {
	return s.EncryptOption
}

func (s *ListConnectionsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectionsShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ListConnectionsShrinkRequest) GetModelTypesShrink() *string {
	return s.ModelTypesShrink
}

func (s *ListConnectionsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectionsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListConnectionsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListConnectionsShrinkRequest) GetToolCall() *bool {
	return s.ToolCall
}

func (s *ListConnectionsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListConnectionsShrinkRequest) SetAccessibility(v string) *ListConnectionsShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetConnectionIdsShrink(v string) *ListConnectionsShrinkRequest {
	s.ConnectionIdsShrink = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetConnectionName(v string) *ListConnectionsShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetConnectionTypesShrink(v string) *ListConnectionsShrinkRequest {
	s.ConnectionTypesShrink = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetCreator(v string) *ListConnectionsShrinkRequest {
	s.Creator = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetEncryptOption(v string) *ListConnectionsShrinkRequest {
	s.EncryptOption = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetMaxResults(v int32) *ListConnectionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetModel(v string) *ListConnectionsShrinkRequest {
	s.Model = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetModelTypesShrink(v string) *ListConnectionsShrinkRequest {
	s.ModelTypesShrink = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetNextToken(v string) *ListConnectionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetOrder(v string) *ListConnectionsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetSortBy(v string) *ListConnectionsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetToolCall(v bool) *ListConnectionsShrinkRequest {
	s.ToolCall = &v
	return s
}

func (s *ListConnectionsShrinkRequest) SetWorkspaceId(v string) *ListConnectionsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListConnectionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
