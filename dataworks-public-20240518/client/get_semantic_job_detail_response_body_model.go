// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSemanticJobDetailResponseBodyData) *GetSemanticJobDetailResponseBody
	GetData() *GetSemanticJobDetailResponseBodyData
	SetRequestId(v string) *GetSemanticJobDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSemanticJobDetailResponseBody
	GetSuccess() *bool
}

type GetSemanticJobDetailResponseBody struct {
	// The job details returned by the executor. Used to determine the run status and view the actual runtime configuration.
	Data *GetSemanticJobDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSemanticJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSemanticJobDetailResponseBody) GetData() *GetSemanticJobDetailResponseBodyData {
	return s.Data
}

func (s *GetSemanticJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSemanticJobDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSemanticJobDetailResponseBody) SetData(v *GetSemanticJobDetailResponseBodyData) *GetSemanticJobDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSemanticJobDetailResponseBody) SetRequestId(v string) *GetSemanticJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSemanticJobDetailResponseBody) SetSuccess(v bool) *GetSemanticJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetSemanticJobDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSemanticJobDetailResponseBodyData struct {
	// The advanced runtime settings returned by the executor.
	AdvanceSettings map[string]interface{} `json:"AdvanceSettings,omitempty" xml:"AdvanceSettings,omitempty"`
	// The code parameter information returned by the executor. Used for troubleshooting the runtime configuration of this run.
	//
	// example:
	//
	// --limit 100
	CodeParameters *string `json:"CodeParameters,omitempty" xml:"CodeParameters,omitempty"`
	// The index of the SQL fragment currently being processed by the executor.
	//
	// example:
	//
	// 0
	CurrentSqlIndex *int32 `json:"CurrentSqlIndex,omitempty" xml:"CurrentSqlIndex,omitempty"`
	// The customer identifier of the executor job.
	//
	// example:
	//
	// meta_semantic
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// The data source identifier used by the executor job.
	//
	// example:
	//
	// maxcompute
	Datasource *string `json:"Datasource,omitempty" xml:"Datasource,omitempty"`
	// The runtime environment identifier returned by the executor.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The list of execution type codes returned by the executor.
	ExecTypes []*int32 `json:"ExecTypes,omitempty" xml:"ExecTypes,omitempty" type:"Repeated"`
	// The executor job ID.
	//
	// example:
	//
	// exec-job-demo
	ExecutorJobId *string `json:"ExecutorJobId,omitempty" xml:"ExecutorJobId,omitempty"`
	// The node type code of the executor. Semantic jobs use Shell node code 6.
	//
	// example:
	//
	// 6
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The DataWorks workspace ID associated with the executor job.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the resource group that actually executed the job.
	//
	// example:
	//
	// rg-demo
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resource URLs associated with the executor job.
	ResourceUrls []map[string]interface{} `json:"ResourceUrls,omitempty" xml:"ResourceUrls,omitempty" type:"Repeated"`
	// The list of status codes returned by the executor. Used to determine the current or final status of the job.
	Statuses []*int32 `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s GetSemanticJobDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSemanticJobDetailResponseBodyData) GetAdvanceSettings() map[string]interface{} {
	return s.AdvanceSettings
}

func (s *GetSemanticJobDetailResponseBodyData) GetCodeParameters() *string {
	return s.CodeParameters
}

func (s *GetSemanticJobDetailResponseBodyData) GetCurrentSqlIndex() *int32 {
	return s.CurrentSqlIndex
}

func (s *GetSemanticJobDetailResponseBodyData) GetCustomerName() *string {
	return s.CustomerName
}

func (s *GetSemanticJobDetailResponseBodyData) GetDatasource() *string {
	return s.Datasource
}

func (s *GetSemanticJobDetailResponseBodyData) GetEnv() *string {
	return s.Env
}

func (s *GetSemanticJobDetailResponseBodyData) GetExecTypes() []*int32 {
	return s.ExecTypes
}

func (s *GetSemanticJobDetailResponseBodyData) GetExecutorJobId() *string {
	return s.ExecutorJobId
}

func (s *GetSemanticJobDetailResponseBodyData) GetFileType() *int32 {
	return s.FileType
}

func (s *GetSemanticJobDetailResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSemanticJobDetailResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetSemanticJobDetailResponseBodyData) GetResourceUrls() []map[string]interface{} {
	return s.ResourceUrls
}

func (s *GetSemanticJobDetailResponseBodyData) GetStatuses() []*int32 {
	return s.Statuses
}

func (s *GetSemanticJobDetailResponseBodyData) SetAdvanceSettings(v map[string]interface{}) *GetSemanticJobDetailResponseBodyData {
	s.AdvanceSettings = v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetCodeParameters(v string) *GetSemanticJobDetailResponseBodyData {
	s.CodeParameters = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetCurrentSqlIndex(v int32) *GetSemanticJobDetailResponseBodyData {
	s.CurrentSqlIndex = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetCustomerName(v string) *GetSemanticJobDetailResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetDatasource(v string) *GetSemanticJobDetailResponseBodyData {
	s.Datasource = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetEnv(v string) *GetSemanticJobDetailResponseBodyData {
	s.Env = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetExecTypes(v []*int32) *GetSemanticJobDetailResponseBodyData {
	s.ExecTypes = v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetExecutorJobId(v string) *GetSemanticJobDetailResponseBodyData {
	s.ExecutorJobId = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetFileType(v int32) *GetSemanticJobDetailResponseBodyData {
	s.FileType = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetProjectId(v int64) *GetSemanticJobDetailResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetResourceGroupId(v string) *GetSemanticJobDetailResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetResourceUrls(v []map[string]interface{}) *GetSemanticJobDetailResponseBodyData {
	s.ResourceUrls = v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) SetStatuses(v []*int32) *GetSemanticJobDetailResponseBodyData {
	s.Statuses = v
	return s
}

func (s *GetSemanticJobDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
