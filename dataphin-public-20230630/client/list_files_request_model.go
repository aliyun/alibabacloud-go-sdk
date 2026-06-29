// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListFilesRequestListQuery) *ListFilesRequest
	GetListQuery() *ListFilesRequestListQuery
	SetOpTenantId(v int64) *ListFilesRequest
	GetOpTenantId() *int64
}

type ListFilesRequest struct {
	// Query conditions
	//
	// This parameter is required.
	ListQuery *ListFilesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFilesRequest) GetListQuery() *ListFilesRequestListQuery {
	return s.ListQuery
}

func (s *ListFilesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListFilesRequest) SetListQuery(v *ListFilesRequestListQuery) *ListFilesRequest {
	s.ListQuery = v
	return s
}

func (s *ListFilesRequest) SetOpTenantId(v int64) *ListFilesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListFilesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFilesRequestListQuery struct {
	// Menu tree category
	//
	// - Ad hoc query: tempCode
	//
	// - Code task: codeManage
	//
	// - Offline pipeline: offlinePipeline
	//
	// - Pipeline custom component: offlinePipelineCustomPlugin
	//
	// - Sync task: dataX
	//
	// - Real-time meta table: streamMeta
	//
	// - Real-time custom source: streamCustomDataSource
	//
	// - Real-time computing template: streamTemplate
	//
	// - Resource management: resourceManage
	//
	// - Offline function: udfResource
	//
	// - Real-time function: streamFunction
	//
	// This parameter is required.
	//
	// example:
	//
	// tempCode
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// File directory
	//
	// This parameter is required.
	//
	// example:
	//
	// /xx/x
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// Environment identifier
	//
	// - DEV
	//
	// - PROD
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Project ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 11112311
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Whether to recursively traverse and query
	//
	// - true: Yes, returns all nodes under the directory
	//
	// - false: No, returns only the first-level nodes under the directory
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
}

func (s ListFilesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListFilesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListFilesRequestListQuery) GetCategory() *string {
	return s.Category
}

func (s *ListFilesRequestListQuery) GetDirectory() *string {
	return s.Directory
}

func (s *ListFilesRequestListQuery) GetEnv() *string {
	return s.Env
}

func (s *ListFilesRequestListQuery) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFilesRequestListQuery) GetRecursive() *bool {
	return s.Recursive
}

func (s *ListFilesRequestListQuery) SetCategory(v string) *ListFilesRequestListQuery {
	s.Category = &v
	return s
}

func (s *ListFilesRequestListQuery) SetDirectory(v string) *ListFilesRequestListQuery {
	s.Directory = &v
	return s
}

func (s *ListFilesRequestListQuery) SetEnv(v string) *ListFilesRequestListQuery {
	s.Env = &v
	return s
}

func (s *ListFilesRequestListQuery) SetProjectId(v int64) *ListFilesRequestListQuery {
	s.ProjectId = &v
	return s
}

func (s *ListFilesRequestListQuery) SetRecursive(v bool) *ListFilesRequestListQuery {
	s.Recursive = &v
	return s
}

func (s *ListFilesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
