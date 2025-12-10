// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMCTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *SearchMCTablesRequest
	GetKeyword() *string
	SetWorkspaceId(v string) *SearchMCTablesRequest
	GetWorkspaceId() *string
}

type SearchMCTablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project.table_name_prefix
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchMCTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMCTablesRequest) GoString() string {
	return s.String()
}

func (s *SearchMCTablesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchMCTablesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SearchMCTablesRequest) SetKeyword(v string) *SearchMCTablesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchMCTablesRequest) SetWorkspaceId(v string) *SearchMCTablesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchMCTablesRequest) Validate() error {
	return dara.Validate(s)
}
