// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListExecutorGroupResponseBody
	GetCode() *int32
	SetData(v *ListExecutorGroupResponseBodyData) *ListExecutorGroupResponseBody
	GetData() *ListExecutorGroupResponseBodyData
	SetMaxResults(v int32) *ListExecutorGroupResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListExecutorGroupResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListExecutorGroupResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExecutorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExecutorGroupResponseBody
	GetSuccess() *bool
}

type ListExecutorGroupResponseBody struct {
	// The total count of data under the current request conditions. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query results.
	Data *ListExecutorGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries to return for this request. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position from which the current call returns data. An empty value indicates that all data has been read.
	//
	// This parameter is required.
	//
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that specifies the position from which the current call returns data. An empty value indicates that all data has been read. You do not need to specify this parameter for the first query.
	//
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExecutorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListExecutorGroupResponseBody) GetData() *ListExecutorGroupResponseBodyData {
	return s.Data
}

func (s *ListExecutorGroupResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExecutorGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExecutorGroupResponseBody) SetCode(v int32) *ListExecutorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetData(v *ListExecutorGroupResponseBodyData) *ListExecutorGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListExecutorGroupResponseBody) SetMaxResults(v int32) *ListExecutorGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetMessage(v string) *ListExecutorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetNextToken(v string) *ListExecutorGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetRequestId(v string) *ListExecutorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorGroupResponseBody) SetSuccess(v bool) *ListExecutorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListExecutorGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExecutorGroupResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListExecutorGroupResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListExecutorGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorGroupResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorGroupResponseBodyData) GetRecords() []*ListExecutorGroupResponseBodyDataRecords {
	return s.Records
}

func (s *ListExecutorGroupResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListExecutorGroupResponseBodyData) SetPageNumber(v int32) *ListExecutorGroupResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorGroupResponseBodyData) SetPageSize(v int32) *ListExecutorGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListExecutorGroupResponseBodyData) SetRecords(v []*ListExecutorGroupResponseBodyDataRecords) *ListExecutorGroupResponseBodyData {
	s.Records = v
	return s
}

func (s *ListExecutorGroupResponseBodyData) SetTotal(v int32) *ListExecutorGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListExecutorGroupResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExecutorGroupResponseBodyDataRecords struct {
	// The client key.
	//
	// example:
	//
	// f/HElmWgOgmb0mlbRRkNuQ==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The authentication type.
	//
	// example:
	//
	// api_key
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Specifies whether to enable **automatic start and stop**. Valid values:
	//
	// - true: Enabled. If no tasks are scheduled for a period of time, the agent automatically hibernates. If tasks are scheduled for a period of time, the agent wakes up in advance.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The list of referenced applications.
	CiteList []*ListExecutorGroupResponseBodyDataRecordsCiteList `json:"CiteList,omitempty" xml:"CiteList,omitempty" type:"Repeated"`
	// The CloudMonitor workspace.
	//
	// example:
	//
	// scx-workspace-acd09a5d-4133-43cc-889d-c33a122afb6e
	CmsWorkspaceId *string `json:"CmsWorkspaceId,omitempty" xml:"CmsWorkspaceId,omitempty"`
	// The current number of tasks.
	//
	// example:
	//
	// 10
	CurrentJobs *int32 `json:"CurrentJobs,omitempty" xml:"CurrentJobs,omitempty"`
	// The description of the executor group.
	//
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableJobs  *int32  `json:"EnableJobs,omitempty" xml:"EnableJobs,omitempty"`
	// The integration method. Valid values:
	//
	// - manual
	//
	// - acs_sandbox
	//
	// - k8s_service
	//
	// example:
	//
	// acs_sandbox
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	// The maximum number of tasks.
	//
	// example:
	//
	// 1000
	MaxJobs *int32 `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	// The name of the executor group.
	//
	// example:
	//
	// job01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network configuration.
	//
	// example:
	//
	// public
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The protocol.
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The list of resolved workers.
	//
	// example:
	//
	// [{"address":"http://192.168.2.34:8084"}]
	ResolvedWorkers *string `json:"ResolvedWorkers,omitempty" xml:"ResolvedWorkers,omitempty"`
	// The ID of the executor group.
	//
	// example:
	//
	// 12
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// The executor type.
	//
	// example:
	//
	// k8s_service
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// The list of specified machines in JSON format.
	//
	// example:
	//
	// [{"address":"http://47.111.188.191:18789"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
	// The extended parameters.
	//
	// example:
	//
	// {}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
}

func (s ListExecutorGroupResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetAuthType() *string {
	return s.AuthType
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetCiteList() []*ListExecutorGroupResponseBodyDataRecordsCiteList {
	return s.CiteList
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetCmsWorkspaceId() *string {
	return s.CmsWorkspaceId
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetCurrentJobs() *int32 {
	return s.CurrentJobs
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetEnableJobs() *int32 {
	return s.EnableJobs
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetMaxJobs() *int32 {
	return s.MaxJobs
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetNetwork() *string {
	return s.Network
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetProtocol() *string {
	return s.Protocol
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetResolvedWorkers() *string {
	return s.ResolvedWorkers
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetWorkerType() *string {
	return s.WorkerType
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetWorkers() *string {
	return s.Workers
}

func (s *ListExecutorGroupResponseBodyDataRecords) GetXAttrs() *string {
	return s.XAttrs
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetApiKey(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.ApiKey = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetAuthType(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.AuthType = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetAutoScale(v bool) *ListExecutorGroupResponseBodyDataRecords {
	s.AutoScale = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetCiteList(v []*ListExecutorGroupResponseBodyDataRecordsCiteList) *ListExecutorGroupResponseBodyDataRecords {
	s.CiteList = v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetCmsWorkspaceId(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.CmsWorkspaceId = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetCurrentJobs(v int32) *ListExecutorGroupResponseBodyDataRecords {
	s.CurrentJobs = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetDescription(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetEnableJobs(v int32) *ListExecutorGroupResponseBodyDataRecords {
	s.EnableJobs = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetIntegrationType(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.IntegrationType = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetMaxJobs(v int32) *ListExecutorGroupResponseBodyDataRecords {
	s.MaxJobs = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetName(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetNetwork(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Network = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetProtocol(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Protocol = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetResolvedWorkers(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.ResolvedWorkers = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetWorkerId(v int64) *ListExecutorGroupResponseBodyDataRecords {
	s.WorkerId = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetWorkerType(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.WorkerType = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetWorkers(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.Workers = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) SetXAttrs(v string) *ListExecutorGroupResponseBodyDataRecords {
	s.XAttrs = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecords) Validate() error {
	if s.CiteList != nil {
		for _, item := range s.CiteList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExecutorGroupResponseBodyDataRecordsCiteList struct {
	// The name of the application.
	//
	// example:
	//
	// Tornado
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// group1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ListExecutorGroupResponseBodyDataRecordsCiteList) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupResponseBodyDataRecordsCiteList) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) GetAppName() *string {
	return s.AppName
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) GetDescription() *string {
	return s.Description
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) SetAppName(v string) *ListExecutorGroupResponseBodyDataRecordsCiteList {
	s.AppName = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) SetDescription(v string) *ListExecutorGroupResponseBodyDataRecordsCiteList {
	s.Description = &v
	return s
}

func (s *ListExecutorGroupResponseBodyDataRecordsCiteList) Validate() error {
	return dara.Validate(s)
}
