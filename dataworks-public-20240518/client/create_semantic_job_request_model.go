// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSemanticJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSemanticJobRequest
	GetName() *string
	SetProjectId(v int64) *CreateSemanticJobRequest
	GetProjectId() *int64
	SetReferenceFileIds(v []*string) *CreateSemanticJobRequest
	GetReferenceFileIds() []*string
	SetReferenceFileUris(v []*string) *CreateSemanticJobRequest
	GetReferenceFileUris() []*string
	SetResourceGroupId(v string) *CreateSemanticJobRequest
	GetResourceGroupId() *string
	SetSource(v map[string]interface{}) *CreateSemanticJobRequest
	GetSource() map[string]interface{}
}

type CreateSemanticJobRequest struct {
	// The semantic job name, which also serves as the job identifier for subsequent calls to RunSemanticJob, DeleteSemanticJob, ListSemanticJobRuns, and DownloadSemanticResults. The name must be unique within the current tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// semantic-job-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. This parameter is required for all Source.type values except singleTableFile. The Data.ProjectId in the creation result can be reused for GetSemanticJobDetail, GetSemanticJobLog, and KillSemanticJob.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of uploaded reference file IDs. When Source.type=singleTableFile, use either this parameter or ReferenceFileUris, and the selected array must contain exactly one non-empty element. The ID must come from Data.FileId returned by UploadSemanticFile, and only CSV or XLSX files are supported. For other Source.type values, you can pass multiple IDs. The service validates each ID during creation, and you can also pass ReferenceFileUris at the same time.
	ReferenceFileIds []*string `json:"ReferenceFileIds,omitempty" xml:"ReferenceFileIds,omitempty" type:"Repeated"`
	// The list of reference file URIs accessible by the caller. When Source.type=singleTableFile, use either this parameter or ReferenceFileIds, and the selected array must contain exactly one non-empty URI. For other Source.type values, you can pass multiple URIs and also pass ReferenceFileIds at the same time. When using the upload path from UploadSemanticFile, pass Data.FileId after the PUT upload is complete instead of the short-lived UploadUrl.
	ReferenceFileUris []*string `json:"ReferenceFileUris,omitempty" xml:"ReferenceFileUris,omitempty" type:"Repeated"`
	// The ID of the resource group used to run the semantic job. RunSemanticJob does not accept this parameter and instead uses the resource group saved during creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-demo
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The input datasource config for the semantic node. The type field is required. This parameter specifies the data to be analyzed and is not the semantic_model YAML output. The domain field is a character string that serves as the identity of the business domain and focus of the node, such as sales. Supported types: 1) maxcompute: Use pinnedScopeInfo to specify the scope. Array elements contain type and name. When type=project, name is the MaxCompute project name. When type=schema, project is the project name and name is the schema name. For table-level scope, project is the project name, schema is optional, and name is the table name. 2) holo or starrocks: In addition to type, you must specify dataSourceName and dataSourceEnv, and pass ProjectId at the top level of the request. You can use pinnedScopeInfo to limit the scope to schemas or tables. The name element is the schema or table name, and the schema element for table-level scope is the database or schema. 3) singleTableFile: ProjectId is not required. For file reference rules, see ReferenceFileIds and ReferenceFileUris. After the node runs successfully, use DownloadSemanticResults to retrieve the semantic_model YAML and other result files. The example shows a MaxCompute project-level scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"type":"maxcompute","domain":"sales","pinnedScopeInfo":[{"type":"project","name":"mc_project"}]}
	Source map[string]interface{} `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateSemanticJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSemanticJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateSemanticJobRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateSemanticJobRequest) GetReferenceFileIds() []*string {
	return s.ReferenceFileIds
}

func (s *CreateSemanticJobRequest) GetReferenceFileUris() []*string {
	return s.ReferenceFileUris
}

func (s *CreateSemanticJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSemanticJobRequest) GetSource() map[string]interface{} {
	return s.Source
}

func (s *CreateSemanticJobRequest) SetName(v string) *CreateSemanticJobRequest {
	s.Name = &v
	return s
}

func (s *CreateSemanticJobRequest) SetProjectId(v int64) *CreateSemanticJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateSemanticJobRequest) SetReferenceFileIds(v []*string) *CreateSemanticJobRequest {
	s.ReferenceFileIds = v
	return s
}

func (s *CreateSemanticJobRequest) SetReferenceFileUris(v []*string) *CreateSemanticJobRequest {
	s.ReferenceFileUris = v
	return s
}

func (s *CreateSemanticJobRequest) SetResourceGroupId(v string) *CreateSemanticJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSemanticJobRequest) SetSource(v map[string]interface{}) *CreateSemanticJobRequest {
	s.Source = v
	return s
}

func (s *CreateSemanticJobRequest) Validate() error {
	return dara.Validate(s)
}
