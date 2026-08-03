// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentAccuracyTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *CreateDataAgentAccuracyTestRequest
	GetCustomAgentId() *string
	SetDataset(v string) *CreateDataAgentAccuracyTestRequest
	GetDataset() *string
	SetDatasource(v string) *CreateDataAgentAccuracyTestRequest
	GetDatasource() *string
	SetDesc(v string) *CreateDataAgentAccuracyTestRequest
	GetDesc() *string
	SetDmsUnit(v string) *CreateDataAgentAccuracyTestRequest
	GetDmsUnit() *string
	SetEvaluationPrompt(v string) *CreateDataAgentAccuracyTestRequest
	GetEvaluationPrompt() *string
	SetFileId(v string) *CreateDataAgentAccuracyTestRequest
	GetFileId() *string
	SetLanguage(v string) *CreateDataAgentAccuracyTestRequest
	GetLanguage() *string
	SetMaxConcurrent(v int32) *CreateDataAgentAccuracyTestRequest
	GetMaxConcurrent() *int32
	SetMode(v int32) *CreateDataAgentAccuracyTestRequest
	GetMode() *int32
	SetName(v string) *CreateDataAgentAccuracyTestRequest
	GetName() *string
	SetNeedDelete(v bool) *CreateDataAgentAccuracyTestRequest
	GetNeedDelete() *bool
	SetRegionId(v string) *CreateDataAgentAccuracyTestRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CreateDataAgentAccuracyTestRequest
	GetWorkspaceId() *string
}

type CreateDataAgentAccuracyTestRequest struct {
	// The ID of the custom agent to be tested for accuracy.
	//
	// example:
	//
	// ca-xxxxxxxxxxxxxxxxxxxx
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// null
	Dataset *string `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// The data source. We recommend that you configure this in the custom agent.
	//
	// example:
	//
	// [{"isInternal":"N","TableIds":["51***70","51***71"],"DataSourceType":"database","Database":"internal_data_employees","DmsInstanceId":"27***5","DmsDatabaseId":"71***04","Tables":["employees","salaries"],"FileId":"rm-
	//
	// ***","DbName":"internal_data_employees","CatalogName":"def","RegionId":"cn-hangzhou","Engine":"mysql"}]
	Datasource *string `json:"Datasource,omitempty" xml:"Datasource,omitempty"`
	// The description.
	//
	// example:
	//
	// null
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The DMS unit used to create the resource.
	//
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// The accuracy evaluation criteria. An empty value indicates the default criteria.
	//
	// example:
	//
	// null
	EvaluationPrompt *string `json:"EvaluationPrompt,omitempty" xml:"EvaluationPrompt,omitempty"`
	// The file ID in the data center.
	//
	// example:
	//
	// f-xxxxxxxxxxxxxxxxxxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The language used for the analysis task.
	//
	// example:
	//
	// ENGLISH
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The maximum number of concurrent sessions during the test.
	//
	// example:
	//
	// 5
	MaxConcurrent *int32 `json:"MaxConcurrent,omitempty" xml:"MaxConcurrent,omitempty"`
	// The analysis mode.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the test item.
	//
	// example:
	//
	// Test01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether sessions are displayed after analysis. This parameter is not supported.
	//
	// example:
	//
	// null
	NeedDelete *bool `json:"NeedDelete,omitempty" xml:"NeedDelete,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentAccuracyTestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentAccuracyTestRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentAccuracyTestRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *CreateDataAgentAccuracyTestRequest) GetDataset() *string {
	return s.Dataset
}

func (s *CreateDataAgentAccuracyTestRequest) GetDatasource() *string {
	return s.Datasource
}

func (s *CreateDataAgentAccuracyTestRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreateDataAgentAccuracyTestRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *CreateDataAgentAccuracyTestRequest) GetEvaluationPrompt() *string {
	return s.EvaluationPrompt
}

func (s *CreateDataAgentAccuracyTestRequest) GetFileId() *string {
	return s.FileId
}

func (s *CreateDataAgentAccuracyTestRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateDataAgentAccuracyTestRequest) GetMaxConcurrent() *int32 {
	return s.MaxConcurrent
}

func (s *CreateDataAgentAccuracyTestRequest) GetMode() *int32 {
	return s.Mode
}

func (s *CreateDataAgentAccuracyTestRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataAgentAccuracyTestRequest) GetNeedDelete() *bool {
	return s.NeedDelete
}

func (s *CreateDataAgentAccuracyTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataAgentAccuracyTestRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentAccuracyTestRequest) SetCustomAgentId(v string) *CreateDataAgentAccuracyTestRequest {
	s.CustomAgentId = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetDataset(v string) *CreateDataAgentAccuracyTestRequest {
	s.Dataset = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetDatasource(v string) *CreateDataAgentAccuracyTestRequest {
	s.Datasource = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetDesc(v string) *CreateDataAgentAccuracyTestRequest {
	s.Desc = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetDmsUnit(v string) *CreateDataAgentAccuracyTestRequest {
	s.DmsUnit = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetEvaluationPrompt(v string) *CreateDataAgentAccuracyTestRequest {
	s.EvaluationPrompt = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetFileId(v string) *CreateDataAgentAccuracyTestRequest {
	s.FileId = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetLanguage(v string) *CreateDataAgentAccuracyTestRequest {
	s.Language = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetMaxConcurrent(v int32) *CreateDataAgentAccuracyTestRequest {
	s.MaxConcurrent = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetMode(v int32) *CreateDataAgentAccuracyTestRequest {
	s.Mode = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetName(v string) *CreateDataAgentAccuracyTestRequest {
	s.Name = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetNeedDelete(v bool) *CreateDataAgentAccuracyTestRequest {
	s.NeedDelete = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetRegionId(v string) *CreateDataAgentAccuracyTestRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) SetWorkspaceId(v string) *CreateDataAgentAccuracyTestRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentAccuracyTestRequest) Validate() error {
	return dara.Validate(s)
}
