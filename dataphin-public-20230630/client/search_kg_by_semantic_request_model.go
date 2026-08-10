// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKgBySemanticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SearchKgBySemanticRequest
	GetOpTenantId() *int64
	SetSearchCommand(v *SearchKgBySemanticRequestSearchCommand) *SearchKgBySemanticRequest
	GetSearchCommand() *SearchKgBySemanticRequestSearchCommand
	SetWorkspaceId(v string) *SearchKgBySemanticRequest
	GetWorkspaceId() *string
}

type SearchKgBySemanticRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The search command.
	//
	// This parameter is required.
	SearchCommand *SearchKgBySemanticRequestSearchCommand `json:"SearchCommand,omitempty" xml:"SearchCommand,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchKgBySemanticRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchKgBySemanticRequest) GoString() string {
	return s.String()
}

func (s *SearchKgBySemanticRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SearchKgBySemanticRequest) GetSearchCommand() *SearchKgBySemanticRequestSearchCommand {
	return s.SearchCommand
}

func (s *SearchKgBySemanticRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SearchKgBySemanticRequest) SetOpTenantId(v int64) *SearchKgBySemanticRequest {
	s.OpTenantId = &v
	return s
}

func (s *SearchKgBySemanticRequest) SetSearchCommand(v *SearchKgBySemanticRequestSearchCommand) *SearchKgBySemanticRequest {
	s.SearchCommand = v
	return s
}

func (s *SearchKgBySemanticRequest) SetWorkspaceId(v string) *SearchKgBySemanticRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchKgBySemanticRequest) Validate() error {
	if s.SearchCommand != nil {
		if err := s.SearchCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchKgBySemanticRequestSearchCommand struct {
	// The entity type codes used for filtering. If this parameter is not specified, all entity types are searched.
	EntityTypeCodes []*string `json:"EntityTypeCodes,omitempty" xml:"EntityTypeCodes,omitempty" type:"Repeated"`
	// The minimum similarity threshold. Valid values: 0.0 to 1.0. Default value: 0.0 (no filtering). This parameter takes effect only for the semantic search path.
	//
	// example:
	//
	// 0.5
	MinSimilarity *float32 `json:"MinSimilarity,omitempty" xml:"MinSimilarity,omitempty"`
	// The property code for semantic search. If this parameter is not specified, all properties with semantic search enabled are searched.
	//
	// example:
	//
	// name
	PropertyCode *string `json:"PropertyCode,omitempty" xml:"PropertyCode,omitempty"`
	// The natural language query text. The value can be 0 to 500 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Query students in Beijing
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The maximum number of results to return. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s SearchKgBySemanticRequestSearchCommand) String() string {
	return dara.Prettify(s)
}

func (s SearchKgBySemanticRequestSearchCommand) GoString() string {
	return s.String()
}

func (s *SearchKgBySemanticRequestSearchCommand) GetEntityTypeCodes() []*string {
	return s.EntityTypeCodes
}

func (s *SearchKgBySemanticRequestSearchCommand) GetMinSimilarity() *float32 {
	return s.MinSimilarity
}

func (s *SearchKgBySemanticRequestSearchCommand) GetPropertyCode() *string {
	return s.PropertyCode
}

func (s *SearchKgBySemanticRequestSearchCommand) GetQueryText() *string {
	return s.QueryText
}

func (s *SearchKgBySemanticRequestSearchCommand) GetTopK() *int32 {
	return s.TopK
}

func (s *SearchKgBySemanticRequestSearchCommand) SetEntityTypeCodes(v []*string) *SearchKgBySemanticRequestSearchCommand {
	s.EntityTypeCodes = v
	return s
}

func (s *SearchKgBySemanticRequestSearchCommand) SetMinSimilarity(v float32) *SearchKgBySemanticRequestSearchCommand {
	s.MinSimilarity = &v
	return s
}

func (s *SearchKgBySemanticRequestSearchCommand) SetPropertyCode(v string) *SearchKgBySemanticRequestSearchCommand {
	s.PropertyCode = &v
	return s
}

func (s *SearchKgBySemanticRequestSearchCommand) SetQueryText(v string) *SearchKgBySemanticRequestSearchCommand {
	s.QueryText = &v
	return s
}

func (s *SearchKgBySemanticRequestSearchCommand) SetTopK(v int32) *SearchKgBySemanticRequestSearchCommand {
	s.TopK = &v
	return s
}

func (s *SearchKgBySemanticRequestSearchCommand) Validate() error {
	return dara.Validate(s)
}
