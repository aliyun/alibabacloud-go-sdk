// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetDirectoryTreeRequest
	GetCategory() *string
	SetOpTenantId(v int64) *GetDirectoryTreeRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetDirectoryTreeRequest
	GetProjectId() *int64
}

type GetDirectoryTreeRequest struct {
	// Directory type.
	//
	// - Ad-hoc Query: tempCode
	//
	// - Code Task: codeManage
	//
	// - Offline Pipeline: offlinePipeline
	//
	// - Pipeline Custom Component: offlinePipelineCustomPlugin
	//
	// - Sync Task: dataX
	//
	// - Real-time Meta Table: streamMeta
	//
	// - Real-time Custom Source: streamCustomDataSource
	//
	// - Real-time Compute Template: streamTemplate
	//
	// - Resource Management: resourceManage
	//
	// - Offline Function: udfResource
	//
	// - Real-time Function: streamFunction
	//
	// This parameter is required.
	//
	// example:
	//
	// codeManage
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// Project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7081229106458752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDirectoryTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryTreeRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryTreeRequest) GetCategory() *string {
	return s.Category
}

func (s *GetDirectoryTreeRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDirectoryTreeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDirectoryTreeRequest) SetCategory(v string) *GetDirectoryTreeRequest {
	s.Category = &v
	return s
}

func (s *GetDirectoryTreeRequest) SetOpTenantId(v int64) *GetDirectoryTreeRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDirectoryTreeRequest) SetProjectId(v int64) *GetDirectoryTreeRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDirectoryTreeRequest) Validate() error {
	return dara.Validate(s)
}
