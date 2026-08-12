// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchImportTaskResponseBody
	GetCode() *string
	SetData(v *GetBatchImportTaskResponseBodyData) *GetBatchImportTaskResponseBody
	GetData() *GetBatchImportTaskResponseBodyData
	SetMessage(v string) *GetBatchImportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchImportTaskResponseBody
	GetRequestId() *string
}

type GetBatchImportTaskResponseBody struct {
	// example:
	//
	// Ok
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetBatchImportTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// CE534E1D-FCE4-5930-B784-E055EC1AEE6F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetBatchImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchImportTaskResponseBody) GetData() *GetBatchImportTaskResponseBodyData {
	return s.Data
}

func (s *GetBatchImportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchImportTaskResponseBody) SetCode(v string) *GetBatchImportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchImportTaskResponseBody) SetData(v *GetBatchImportTaskResponseBodyData) *GetBatchImportTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchImportTaskResponseBody) SetMessage(v string) *GetBatchImportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchImportTaskResponseBody) SetRequestId(v string) *GetBatchImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchImportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchImportTaskResponseBodyData struct {
	// example:
	//
	// 2026-05-15T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// some apis import failed
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 5
	ProcessedCount *int32                                    `json:"processedCount,omitempty" xml:"processedCount,omitempty"`
	Result         *GetBatchImportTaskResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// async-task-xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// BatchImport
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetBatchImportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetBatchImportTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchImportTaskResponseBodyData) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *GetBatchImportTaskResponseBodyData) GetResult() *GetBatchImportTaskResponseBodyDataResult {
	return s.Result
}

func (s *GetBatchImportTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetBatchImportTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetBatchImportTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetBatchImportTaskResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetBatchImportTaskResponseBodyData) SetCreateTime(v string) *GetBatchImportTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) SetErrorMessage(v string) *GetBatchImportTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) SetProcessedCount(v int32) *GetBatchImportTaskResponseBodyData {
	s.ProcessedCount = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) SetResult(v *GetBatchImportTaskResponseBodyDataResult) *GetBatchImportTaskResponseBodyData {
	s.Result = v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) SetStatus(v string) *GetBatchImportTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) SetTaskId(v string) *GetBatchImportTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) SetTaskType(v string) *GetBatchImportTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) SetTotalCount(v int32) *GetBatchImportTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchImportTaskResponseBodyDataResult struct {
	// example:
	//
	// Http
	ApiType *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	// example:
	//
	// true
	DryRun        *bool                                                    `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	DryRunResults []*GetBatchImportTaskResponseBodyDataResultDryRunResults `json:"dryRunResults,omitempty" xml:"dryRunResults,omitempty" type:"Repeated"`
	FailureItems  []*GetBatchImportTaskResponseBodyDataResultFailureItems  `json:"failureItems,omitempty" xml:"failureItems,omitempty" type:"Repeated"`
	// example:
	//
	// gw-xxx
	GatewayId     *string                                                 `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ImportRequest *GetBatchImportTaskResponseBodyDataResultImportRequest  `json:"importRequest,omitempty" xml:"importRequest,omitempty" type:"Struct"`
	SuccessItems  []*GetBatchImportTaskResponseBodyDataResultSuccessItems `json:"successItems,omitempty" xml:"successItems,omitempty" type:"Repeated"`
}

func (s GetBatchImportTaskResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResult) GetApiType() *string {
	return s.ApiType
}

func (s *GetBatchImportTaskResponseBodyDataResult) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetBatchImportTaskResponseBodyDataResult) GetDryRunResults() []*GetBatchImportTaskResponseBodyDataResultDryRunResults {
	return s.DryRunResults
}

func (s *GetBatchImportTaskResponseBodyDataResult) GetFailureItems() []*GetBatchImportTaskResponseBodyDataResultFailureItems {
	return s.FailureItems
}

func (s *GetBatchImportTaskResponseBodyDataResult) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetBatchImportTaskResponseBodyDataResult) GetImportRequest() *GetBatchImportTaskResponseBodyDataResultImportRequest {
	return s.ImportRequest
}

func (s *GetBatchImportTaskResponseBodyDataResult) GetSuccessItems() []*GetBatchImportTaskResponseBodyDataResultSuccessItems {
	return s.SuccessItems
}

func (s *GetBatchImportTaskResponseBodyDataResult) SetApiType(v string) *GetBatchImportTaskResponseBodyDataResult {
	s.ApiType = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResult) SetDryRun(v bool) *GetBatchImportTaskResponseBodyDataResult {
	s.DryRun = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResult) SetDryRunResults(v []*GetBatchImportTaskResponseBodyDataResultDryRunResults) *GetBatchImportTaskResponseBodyDataResult {
	s.DryRunResults = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResult) SetFailureItems(v []*GetBatchImportTaskResponseBodyDataResultFailureItems) *GetBatchImportTaskResponseBodyDataResult {
	s.FailureItems = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResult) SetGatewayId(v string) *GetBatchImportTaskResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResult) SetImportRequest(v *GetBatchImportTaskResponseBodyDataResultImportRequest) *GetBatchImportTaskResponseBodyDataResult {
	s.ImportRequest = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResult) SetSuccessItems(v []*GetBatchImportTaskResponseBodyDataResultSuccessItems) *GetBatchImportTaskResponseBodyDataResult {
	s.SuccessItems = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResult) Validate() error {
	if s.DryRunResults != nil {
		for _, item := range s.DryRunResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FailureItems != nil {
		for _, item := range s.FailureItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImportRequest != nil {
		if err := s.ImportRequest.Validate(); err != nil {
			return err
		}
	}
	if s.SuccessItems != nil {
		for _, item := range s.SuccessItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchImportTaskResponseBodyDataResultDryRunResults struct {
	// example:
	//
	// petstore
	ApiName    *string                                                          `json:"apiName,omitempty" xml:"apiName,omitempty"`
	DryRunInfo *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo `json:"dryRunInfo,omitempty" xml:"dryRunInfo,omitempty" type:"Struct"`
	// example:
	//
	// unsupported oas version
	Error *string `json:"error,omitempty" xml:"error,omitempty"`
	// example:
	//
	// petstore.yaml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResults) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResults) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) GetApiName() *string {
	return s.ApiName
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) GetDryRunInfo() *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	return s.DryRunInfo
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) GetError() *string {
	return s.Error
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) GetFileName() *string {
	return s.FileName
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) SetApiName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResults {
	s.ApiName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) SetDryRunInfo(v *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) *GetBatchImportTaskResponseBodyDataResultDryRunResults {
	s.DryRunInfo = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) SetError(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResults {
	s.Error = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) SetFileName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResults {
	s.FileName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResults) Validate() error {
	if s.DryRunInfo != nil {
		if err := s.DryRunInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo struct {
	ErrorMessages     []*string                                                                           `json:"errorMessages,omitempty" xml:"errorMessages,omitempty" type:"Repeated"`
	ExistHttpApiInfo  *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo    `json:"existHttpApiInfo,omitempty" xml:"existHttpApiInfo,omitempty" type:"Struct"`
	FailureComponents []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents `json:"failureComponents,omitempty" xml:"failureComponents,omitempty" type:"Repeated"`
	FailureOperations []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations `json:"failureOperations,omitempty" xml:"failureOperations,omitempty" type:"Repeated"`
	FailureRoutes     []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes     `json:"failureRoutes,omitempty" xml:"failureRoutes,omitempty" type:"Repeated"`
	SuccessComponents []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents `json:"successComponents,omitempty" xml:"successComponents,omitempty" type:"Repeated"`
	SuccessOperations []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations `json:"successOperations,omitempty" xml:"successOperations,omitempty" type:"Repeated"`
	SuccessRoutes     []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes     `json:"successRoutes,omitempty" xml:"successRoutes,omitempty" type:"Repeated"`
	WarningMessages   []*string                                                                           `json:"warningMessages,omitempty" xml:"warningMessages,omitempty" type:"Repeated"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetErrorMessages() []*string {
	return s.ErrorMessages
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetExistHttpApiInfo() *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo {
	return s.ExistHttpApiInfo
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetFailureComponents() []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents {
	return s.FailureComponents
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetFailureOperations() []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations {
	return s.FailureOperations
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetFailureRoutes() []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes {
	return s.FailureRoutes
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetSuccessComponents() []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents {
	return s.SuccessComponents
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetSuccessOperations() []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations {
	return s.SuccessOperations
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetSuccessRoutes() []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes {
	return s.SuccessRoutes
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) GetWarningMessages() []*string {
	return s.WarningMessages
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetErrorMessages(v []*string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.ErrorMessages = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetExistHttpApiInfo(v *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.ExistHttpApiInfo = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetFailureComponents(v []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.FailureComponents = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetFailureOperations(v []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.FailureOperations = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetFailureRoutes(v []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.FailureRoutes = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetSuccessComponents(v []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.SuccessComponents = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetSuccessOperations(v []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.SuccessOperations = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetSuccessRoutes(v []*GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.SuccessRoutes = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) SetWarningMessages(v []*string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo {
	s.WarningMessages = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfo) Validate() error {
	if s.ExistHttpApiInfo != nil {
		if err := s.ExistHttpApiInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FailureComponents != nil {
		for _, item := range s.FailureComponents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FailureOperations != nil {
		for _, item := range s.FailureOperations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FailureRoutes != nil {
		for _, item := range s.FailureRoutes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessComponents != nil {
		for _, item := range s.SuccessComponents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessOperations != nil {
		for _, item := range s.SuccessOperations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessRoutes != nil {
		for _, item := range s.SuccessRoutes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo struct {
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// example:
	//
	// petstore
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Rest
	Type        *string                                                                                     `json:"type,omitempty" xml:"type,omitempty"`
	VersionInfo *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo `json:"versionInfo,omitempty" xml:"versionInfo,omitempty" type:"Struct"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) GetBasePath() *string {
	return s.BasePath
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) GetName() *string {
	return s.Name
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) GetType() *string {
	return s.Type
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) GetVersionInfo() *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo {
	return s.VersionInfo
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) SetBasePath(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo {
	s.BasePath = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) SetGatewayId(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo {
	s.GatewayId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) SetHttpApiId(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo {
	s.HttpApiId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) SetName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo {
	s.Name = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) SetType(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo {
	s.Type = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) SetVersionInfo(v *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo {
	s.VersionInfo = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfo) Validate() error {
	if s.VersionInfo != nil {
		if err := s.VersionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// my-version
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// example:
	//
	// myVersion
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// example:
	//
	// Query
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) GetEnable() *bool {
	return s.Enable
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) GetHeaderName() *string {
	return s.HeaderName
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) GetQueryName() *string {
	return s.QueryName
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) GetScheme() *string {
	return s.Scheme
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) GetVersion() *string {
	return s.Version
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) SetEnable(v bool) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo {
	s.Enable = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) SetHeaderName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo {
	s.HeaderName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) SetQueryName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo {
	s.QueryName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) SetScheme(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo {
	s.Scheme = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) SetVersion(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo {
	s.Version = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoExistHttpApiInfoVersionInfo) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents struct {
	// example:
	//
	// invalid schema
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// Pet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) GetName() *string {
	return s.Name
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) SetErrorMessage(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) SetName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents {
	s.Name = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureComponents) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations struct {
	// example:
	//
	// unsupported operation
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// /pets/{petId}
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) GetMethod() *string {
	return s.Method
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) GetPath() *string {
	return s.Path
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) SetErrorMessage(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) SetMethod(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations {
	s.Method = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) SetPath(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations {
	s.Path = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureOperations) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes struct {
	// example:
	//
	// domain not found
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// route-pets
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) GetName() *string {
	return s.Name
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) SetErrorMessage(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) SetName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes {
	s.Name = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoFailureRoutes) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents struct {
	// example:
	//
	// Create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// Pet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) GetAction() *string {
	return s.Action
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) GetName() *string {
	return s.Name
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) SetAction(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents {
	s.Action = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) SetName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents {
	s.Name = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessComponents) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations struct {
	// example:
	//
	// Create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// getPetById
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// /pets/{petId}
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) GetAction() *string {
	return s.Action
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) GetMethod() *string {
	return s.Method
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) GetName() *string {
	return s.Name
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) GetPath() *string {
	return s.Path
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) SetAction(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations {
	s.Action = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) SetMethod(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations {
	s.Method = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) SetName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations {
	s.Name = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) SetPath(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations {
	s.Path = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessOperations) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes struct {
	// example:
	//
	// Create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// route-pets
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) GetAction() *string {
	return s.Action
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) GetName() *string {
	return s.Name
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) SetAction(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes {
	s.Action = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) SetName(v string) *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes {
	s.Name = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultDryRunResultsDryRunInfoSuccessRoutes) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultFailureItems struct {
	// example:
	//
	// api-xxx
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// example:
	//
	// petstore
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// example:
	//
	// invalid oas format
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// petstore.yaml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultFailureItems) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultFailureItems) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) GetApiId() *string {
	return s.ApiId
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) GetApiName() *string {
	return s.ApiName
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) GetFileName() *string {
	return s.FileName
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) SetApiId(v string) *GetBatchImportTaskResponseBodyDataResultFailureItems {
	s.ApiId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) SetApiName(v string) *GetBatchImportTaskResponseBodyDataResultFailureItems {
	s.ApiName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) SetErrorMessage(v string) *GetBatchImportTaskResponseBodyDataResultFailureItems {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) SetFileName(v string) *GetBatchImportTaskResponseBodyDataResultFailureItems {
	s.FileName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultFailureItems) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultImportRequest struct {
	AllowUpdate *bool `json:"allowUpdate,omitempty" xml:"allowUpdate,omitempty"`
	// example:
	//
	// Http
	ApiType *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com/my-bucket/imports/batch.zip
	SpecFileUrl   *string                                                             `json:"specFileUrl,omitempty" xml:"specFileUrl,omitempty"`
	SpecOssConfig *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig `json:"specOssConfig,omitempty" xml:"specOssConfig,omitempty" type:"Struct"`
	// example:
	//
	// ExistFirst
	Strategy             *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	WithGatewayExtension *bool   `json:"withGatewayExtension,omitempty" xml:"withGatewayExtension,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultImportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultImportRequest) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetAllowUpdate() *bool {
	return s.AllowUpdate
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetSpecFileUrl() *string {
	return s.SpecFileUrl
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetSpecOssConfig() *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig {
	return s.SpecOssConfig
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) GetWithGatewayExtension() *bool {
	return s.WithGatewayExtension
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetAllowUpdate(v bool) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.AllowUpdate = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetApiType(v string) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.ApiType = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetDryRun(v bool) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.DryRun = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetGatewayId(v string) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.GatewayId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetResourceGroupId(v string) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetSpecFileUrl(v string) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.SpecFileUrl = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetSpecOssConfig(v *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.SpecOssConfig = v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetStrategy(v string) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.Strategy = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) SetWithGatewayExtension(v bool) *GetBatchImportTaskResponseBodyDataResultImportRequest {
	s.WithGatewayExtension = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequest) Validate() error {
	if s.SpecOssConfig != nil {
		if err := s.SpecOssConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig struct {
	// example:
	//
	// my-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// example:
	//
	// imports/batch.zip
	ObjectKey *string `json:"objectKey,omitempty" xml:"objectKey,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) SetBucketName(v string) *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig {
	s.BucketName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) SetObjectKey(v string) *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig {
	s.ObjectKey = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) SetRegionId(v string) *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig {
	s.RegionId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultImportRequestSpecOssConfig) Validate() error {
	return dara.Validate(s)
}

type GetBatchImportTaskResponseBodyDataResultSuccessItems struct {
	// example:
	//
	// api-xxx
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// example:
	//
	// petstore
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// example:
	//
	// invalid oas format
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// petstore.yaml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetBatchImportTaskResponseBodyDataResultSuccessItems) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponseBodyDataResultSuccessItems) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) GetApiId() *string {
	return s.ApiId
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) GetApiName() *string {
	return s.ApiName
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) GetFileName() *string {
	return s.FileName
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) SetApiId(v string) *GetBatchImportTaskResponseBodyDataResultSuccessItems {
	s.ApiId = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) SetApiName(v string) *GetBatchImportTaskResponseBodyDataResultSuccessItems {
	s.ApiName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) SetErrorMessage(v string) *GetBatchImportTaskResponseBodyDataResultSuccessItems {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) SetFileName(v string) *GetBatchImportTaskResponseBodyDataResultSuccessItems {
	s.FileName = &v
	return s
}

func (s *GetBatchImportTaskResponseBodyDataResultSuccessItems) Validate() error {
	return dara.Validate(s)
}
