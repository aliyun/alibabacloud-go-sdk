// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneImportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListInterveneImportTasksRequest
	GetAgentKey() *string
	SetPageIndex(v int32) *ListInterveneImportTasksRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListInterveneImportTasksRequest
	GetPageSize() *int32
}

type ListInterveneImportTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneImportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneImportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListInterveneImportTasksRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListInterveneImportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterveneImportTasksRequest) SetAgentKey(v string) *ListInterveneImportTasksRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneImportTasksRequest) SetPageIndex(v int32) *ListInterveneImportTasksRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneImportTasksRequest) SetPageSize(v int32) *ListInterveneImportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterveneImportTasksRequest) Validate() error {
	return dara.Validate(s)
}
