// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentAccuracyTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccuracyTestInsId(v string) *UpdateDataAgentAccuracyTestRequest
	GetAccuracyTestInsId() *string
	SetCustomerAgentId(v string) *UpdateDataAgentAccuracyTestRequest
	GetCustomerAgentId() *string
	SetDataset(v string) *UpdateDataAgentAccuracyTestRequest
	GetDataset() *string
	SetDesc(v string) *UpdateDataAgentAccuracyTestRequest
	GetDesc() *string
	SetDmsUnit(v string) *UpdateDataAgentAccuracyTestRequest
	GetDmsUnit() *string
	SetEvaluationPrompt(v string) *UpdateDataAgentAccuracyTestRequest
	GetEvaluationPrompt() *string
	SetFileId(v string) *UpdateDataAgentAccuracyTestRequest
	GetFileId() *string
	SetMaxConcurrent(v int32) *UpdateDataAgentAccuracyTestRequest
	GetMaxConcurrent() *int32
	SetMode(v int32) *UpdateDataAgentAccuracyTestRequest
	GetMode() *int32
	SetName(v string) *UpdateDataAgentAccuracyTestRequest
	GetName() *string
	SetNeedDelete(v bool) *UpdateDataAgentAccuracyTestRequest
	GetNeedDelete() *bool
	SetRegionId(v string) *UpdateDataAgentAccuracyTestRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *UpdateDataAgentAccuracyTestRequest
	GetWorkspaceId() *string
}

type UpdateDataAgentAccuracyTestRequest struct {
	// The accuracy test instance ID.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
	// The new custom agent ID.
	//
	// example:
	//
	// ca-xxxxxxxxxxxxxxxxxxxx
	CustomerAgentId *string `json:"CustomerAgentId,omitempty" xml:"CustomerAgentId,omitempty"`
	// The data source.
	//
	// example:
	//
	// [{\\"DataSourceType\\":\\"database\\",\\"RegionId\\":\\"cn-hangzhou\\",\\"DmsInstanceId\\":\\"27xxxxx\\",\\"DmsDatabaseId\\":\\"752xxxxx\\",\\"Database\\":\\"employees\\",\\"Tables\\":[\\"employees\\",\\"salaries\\",\\"departments\\"]}]
	Dataset *string `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// The description.
	//
	// example:
	//
	// null
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The current DMS unit.
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
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The maximum number of concurrent sessions during the test.
	//
	// example:
	//
	// 5
	MaxConcurrent *int32 `json:"MaxConcurrent,omitempty" xml:"MaxConcurrent,omitempty"`
	// The analysis mode to be tested.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the test item.
	//
	// example:
	//
	// test123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether sessions are displayed after analysis. This parameter is not supported.
	//
	// example:
	//
	// no use
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
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDataAgentAccuracyTestRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentAccuracyTestRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentAccuracyTestRequest) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *UpdateDataAgentAccuracyTestRequest) GetCustomerAgentId() *string {
	return s.CustomerAgentId
}

func (s *UpdateDataAgentAccuracyTestRequest) GetDataset() *string {
	return s.Dataset
}

func (s *UpdateDataAgentAccuracyTestRequest) GetDesc() *string {
	return s.Desc
}

func (s *UpdateDataAgentAccuracyTestRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *UpdateDataAgentAccuracyTestRequest) GetEvaluationPrompt() *string {
	return s.EvaluationPrompt
}

func (s *UpdateDataAgentAccuracyTestRequest) GetFileId() *string {
	return s.FileId
}

func (s *UpdateDataAgentAccuracyTestRequest) GetMaxConcurrent() *int32 {
	return s.MaxConcurrent
}

func (s *UpdateDataAgentAccuracyTestRequest) GetMode() *int32 {
	return s.Mode
}

func (s *UpdateDataAgentAccuracyTestRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataAgentAccuracyTestRequest) GetNeedDelete() *bool {
	return s.NeedDelete
}

func (s *UpdateDataAgentAccuracyTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataAgentAccuracyTestRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDataAgentAccuracyTestRequest) SetAccuracyTestInsId(v string) *UpdateDataAgentAccuracyTestRequest {
	s.AccuracyTestInsId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetCustomerAgentId(v string) *UpdateDataAgentAccuracyTestRequest {
	s.CustomerAgentId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetDataset(v string) *UpdateDataAgentAccuracyTestRequest {
	s.Dataset = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetDesc(v string) *UpdateDataAgentAccuracyTestRequest {
	s.Desc = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetDmsUnit(v string) *UpdateDataAgentAccuracyTestRequest {
	s.DmsUnit = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetEvaluationPrompt(v string) *UpdateDataAgentAccuracyTestRequest {
	s.EvaluationPrompt = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetFileId(v string) *UpdateDataAgentAccuracyTestRequest {
	s.FileId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetMaxConcurrent(v int32) *UpdateDataAgentAccuracyTestRequest {
	s.MaxConcurrent = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetMode(v int32) *UpdateDataAgentAccuracyTestRequest {
	s.Mode = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetName(v string) *UpdateDataAgentAccuracyTestRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetNeedDelete(v bool) *UpdateDataAgentAccuracyTestRequest {
	s.NeedDelete = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetRegionId(v string) *UpdateDataAgentAccuracyTestRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) SetWorkspaceId(v string) *UpdateDataAgentAccuracyTestRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataAgentAccuracyTestRequest) Validate() error {
	return dara.Validate(s)
}
