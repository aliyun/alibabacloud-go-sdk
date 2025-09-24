// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionIds(v []*string) *ListConnectionsRequest
	GetConnectionIds() []*string
	SetConnectionName(v string) *ListConnectionsRequest
	GetConnectionName() *string
	SetConnectionTypes(v []*string) *ListConnectionsRequest
	GetConnectionTypes() []*string
	SetCreator(v string) *ListConnectionsRequest
	GetCreator() *string
	SetEncryptOption(v string) *ListConnectionsRequest
	GetEncryptOption() *string
	SetMaxResults(v int32) *ListConnectionsRequest
	GetMaxResults() *int32
	SetModel(v string) *ListConnectionsRequest
	GetModel() *string
	SetModelTypes(v []*string) *ListConnectionsRequest
	GetModelTypes() []*string
	SetNextToken(v string) *ListConnectionsRequest
	GetNextToken() *string
	SetOrder(v string) *ListConnectionsRequest
	GetOrder() *string
	SetSortBy(v string) *ListConnectionsRequest
	GetSortBy() *string
	SetToolCall(v bool) *ListConnectionsRequest
	GetToolCall() *bool
	SetWorkspaceId(v string) *ListConnectionsRequest
	GetWorkspaceId() *string
}

type ListConnectionsRequest struct {
	// The list of connection IDs.
	ConnectionIds []*string `json:"ConnectionIds,omitempty" xml:"ConnectionIds,omitempty" type:"Repeated"`
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The list of connection types.
	ConnectionTypes []*string `json:"ConnectionTypes,omitempty" xml:"ConnectionTypes,omitempty" type:"Repeated"`
	Creator         *string   `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The encryption settings. Valid values:
	//
	// 	- PlainText
	//
	// 	- Secret
	//
	// example:
	//
	// PlainText
	EncryptOption *string `json:"EncryptOption,omitempty" xml:"EncryptOption,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The model identifier.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The list of model types.
	ModelTypes []*string `json:"ModelTypes,omitempty" xml:"ModelTypes,omitempty" type:"Repeated"`
	// The pagination token that indicates the start position from which to retrieve data on the next page.
	//
	// example:
	//
	// 15
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The order in which the entries are sorted by the specific field on the returned page. This parameter must be used together with SortBy.
	//
	// 	- ASC: ascending order.
	//
	// 	- DESC: descending order. This is the default value.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field used to sort the results in queries by page. Default value: GmtCreateTime. Valid value:
	//
	// 	- GmtCreateTime: The results are sorted by creation time. This is the default value.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether a tool can be called by using ToolCall. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ToolCall *bool `json:"ToolCall,omitempty" xml:"ToolCall,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsRequest) GetConnectionIds() []*string {
	return s.ConnectionIds
}

func (s *ListConnectionsRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *ListConnectionsRequest) GetConnectionTypes() []*string {
	return s.ConnectionTypes
}

func (s *ListConnectionsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListConnectionsRequest) GetEncryptOption() *string {
	return s.EncryptOption
}

func (s *ListConnectionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectionsRequest) GetModel() *string {
	return s.Model
}

func (s *ListConnectionsRequest) GetModelTypes() []*string {
	return s.ModelTypes
}

func (s *ListConnectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListConnectionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListConnectionsRequest) GetToolCall() *bool {
	return s.ToolCall
}

func (s *ListConnectionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListConnectionsRequest) SetConnectionIds(v []*string) *ListConnectionsRequest {
	s.ConnectionIds = v
	return s
}

func (s *ListConnectionsRequest) SetConnectionName(v string) *ListConnectionsRequest {
	s.ConnectionName = &v
	return s
}

func (s *ListConnectionsRequest) SetConnectionTypes(v []*string) *ListConnectionsRequest {
	s.ConnectionTypes = v
	return s
}

func (s *ListConnectionsRequest) SetCreator(v string) *ListConnectionsRequest {
	s.Creator = &v
	return s
}

func (s *ListConnectionsRequest) SetEncryptOption(v string) *ListConnectionsRequest {
	s.EncryptOption = &v
	return s
}

func (s *ListConnectionsRequest) SetMaxResults(v int32) *ListConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionsRequest) SetModel(v string) *ListConnectionsRequest {
	s.Model = &v
	return s
}

func (s *ListConnectionsRequest) SetModelTypes(v []*string) *ListConnectionsRequest {
	s.ModelTypes = v
	return s
}

func (s *ListConnectionsRequest) SetNextToken(v string) *ListConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectionsRequest) SetOrder(v string) *ListConnectionsRequest {
	s.Order = &v
	return s
}

func (s *ListConnectionsRequest) SetSortBy(v string) *ListConnectionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListConnectionsRequest) SetToolCall(v bool) *ListConnectionsRequest {
	s.ToolCall = &v
	return s
}

func (s *ListConnectionsRequest) SetWorkspaceId(v string) *ListConnectionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
