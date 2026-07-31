// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSemanticJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSemanticJobShrinkRequest
	GetName() *string
	SetProjectId(v int64) *CreateSemanticJobShrinkRequest
	GetProjectId() *int64
	SetReferenceFileIdsShrink(v string) *CreateSemanticJobShrinkRequest
	GetReferenceFileIdsShrink() *string
	SetReferenceFileUrisShrink(v string) *CreateSemanticJobShrinkRequest
	GetReferenceFileUrisShrink() *string
	SetResourceGroupId(v string) *CreateSemanticJobShrinkRequest
	GetResourceGroupId() *string
	SetSourceShrink(v string) *CreateSemanticJobShrinkRequest
	GetSourceShrink() *string
}

type CreateSemanticJobShrinkRequest struct {
	// The semantic task name, which also serves as the task identifier for subsequent calls to RunSemanticJob, DeleteSemanticJob, ListSemanticJobRuns, and DownloadSemanticResults. The name must be unique within the current tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// semantic-job-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. This parameter is required for all Source.type values except singleTableFile. The Data.ProjectId in the response can be reused for GetSemanticJobDetail, GetSemanticJobLog, and KillSemanticJob.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of uploaded reference file IDs. When Source.type=singleTableFile, use either this parameter or ReferenceFileUris, and the selected array must contain exactly one non-empty element. The ID must come from Data.FileId returned by UploadSemanticFile, and only CSV or XLSX files are supported. For other Source.type values, you can pass multiple IDs. The service validates each ID during creation, and you can also pass ReferenceFileUris at the same time.
	ReferenceFileIdsShrink *string `json:"ReferenceFileIds,omitempty" xml:"ReferenceFileIds,omitempty"`
	// The list of reference file URIs accessible by the caller. When Source.type=singleTableFile, use either this parameter or ReferenceFileIds, and the selected array must contain exactly one non-empty URI. For other Source.type values, you can pass multiple URIs, and you can also pass ReferenceFileIds at the same time. When using the upload path from UploadSemanticFile, pass Data.FileId after the PUT upload is complete instead of the short-lived UploadUrl.
	ReferenceFileUrisShrink *string `json:"ReferenceFileUris,omitempty" xml:"ReferenceFileUris,omitempty"`
	// The ID of the resource group used to run the semantic task. RunSemanticJob does not accept this parameter and instead uses the resource group saved during creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-demo
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The input datasource config for the semantic node. The type field is required and specifies the data to be analyzed. This is not the semantic_model YAML produced by the node. The domain field is a character string that identifies the business domain and focus of the node, such as sales. Supported types: 1) maxcompute: Use pinnedScopeInfo to specify the scope. Array elements contain type and name. When type=project, name is the MaxCompute project name. When type=schema, project is the project name and name is the schema name. For table-level scope, project is the project name, schema is optional, and name is the table name. 2) holo or starrocks: In addition to type, provide dataSourceName and dataSourceEnv, and pass ProjectId at the top level of the request. Use pinnedScopeInfo to limit the scope to schemas or tables. The name element is the schema or table name. For table-level scope, schema is the database or schema that contains the table. 3) singleTableFile: ProjectId is not required. Refer to ReferenceFileIds and ReferenceFileUris for file reference rules. After the node runs successfully, call DownloadSemanticResults to obtain the semantic_model YAML and other result files. The example shows a MaxCompute project-level scope. Active pinnedScopeInfo elements define the scope boundaries.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"type":"maxcompute","domain":"sales","pinnedScopeInfo":[{"type":"project","name":"mc_project"}]}
	SourceShrink *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateSemanticJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSemanticJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateSemanticJobShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateSemanticJobShrinkRequest) GetReferenceFileIdsShrink() *string {
	return s.ReferenceFileIdsShrink
}

func (s *CreateSemanticJobShrinkRequest) GetReferenceFileUrisShrink() *string {
	return s.ReferenceFileUrisShrink
}

func (s *CreateSemanticJobShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSemanticJobShrinkRequest) GetSourceShrink() *string {
	return s.SourceShrink
}

func (s *CreateSemanticJobShrinkRequest) SetName(v string) *CreateSemanticJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSemanticJobShrinkRequest) SetProjectId(v int64) *CreateSemanticJobShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateSemanticJobShrinkRequest) SetReferenceFileIdsShrink(v string) *CreateSemanticJobShrinkRequest {
	s.ReferenceFileIdsShrink = &v
	return s
}

func (s *CreateSemanticJobShrinkRequest) SetReferenceFileUrisShrink(v string) *CreateSemanticJobShrinkRequest {
	s.ReferenceFileUrisShrink = &v
	return s
}

func (s *CreateSemanticJobShrinkRequest) SetResourceGroupId(v string) *CreateSemanticJobShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSemanticJobShrinkRequest) SetSourceShrink(v string) *CreateSemanticJobShrinkRequest {
	s.SourceShrink = &v
	return s
}

func (s *CreateSemanticJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
