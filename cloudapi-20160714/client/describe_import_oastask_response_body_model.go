// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportOASTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiResults(v *DescribeImportOASTaskResponseBodyApiResults) *DescribeImportOASTaskResponseBody
	GetApiResults() *DescribeImportOASTaskResponseBodyApiResults
	SetModelResults(v *DescribeImportOASTaskResponseBodyModelResults) *DescribeImportOASTaskResponseBody
	GetModelResults() *DescribeImportOASTaskResponseBodyModelResults
	SetRequestId(v string) *DescribeImportOASTaskResponseBody
	GetRequestId() *string
	SetTaskStatus(v string) *DescribeImportOASTaskResponseBody
	GetTaskStatus() *string
}

type DescribeImportOASTaskResponseBody struct {
	// The execution status of the subtask. Valid values:
	//
	// 	- RUNNING
	//
	// 	- WAIT
	//
	// 	- OVER
	//
	// 	- FAIL
	//
	// 	- CANCEL
	ApiResults *DescribeImportOASTaskResponseBodyApiResults `json:"ApiResults,omitempty" xml:"ApiResults,omitempty" type:"Struct"`
	// The execution status of the subtask. Valid values:
	//
	// 	- RUNNING
	//
	// 	- WAIT
	//
	// 	- OVER
	//
	// 	- FAIL
	//
	// 	- CANCEL
	ModelResults *DescribeImportOASTaskResponseBodyModelResults `json:"ModelResults,omitempty" xml:"ModelResults,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CE5722A6-AE78-4741-A9B0-6C817D360510
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the import task. Valid values:
	//
	// 	- Running
	//
	// 	- Finished
	//
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeImportOASTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportOASTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportOASTaskResponseBody) GetApiResults() *DescribeImportOASTaskResponseBodyApiResults {
	return s.ApiResults
}

func (s *DescribeImportOASTaskResponseBody) GetModelResults() *DescribeImportOASTaskResponseBodyModelResults {
	return s.ModelResults
}

func (s *DescribeImportOASTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImportOASTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeImportOASTaskResponseBody) SetApiResults(v *DescribeImportOASTaskResponseBodyApiResults) *DescribeImportOASTaskResponseBody {
	s.ApiResults = v
	return s
}

func (s *DescribeImportOASTaskResponseBody) SetModelResults(v *DescribeImportOASTaskResponseBodyModelResults) *DescribeImportOASTaskResponseBody {
	s.ModelResults = v
	return s
}

func (s *DescribeImportOASTaskResponseBody) SetRequestId(v string) *DescribeImportOASTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImportOASTaskResponseBody) SetTaskStatus(v string) *DescribeImportOASTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeImportOASTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImportOASTaskResponseBodyApiResults struct {
	ApiResult []*DescribeImportOASTaskResponseBodyApiResultsApiResult `json:"ApiResult,omitempty" xml:"ApiResult,omitempty" type:"Repeated"`
}

func (s DescribeImportOASTaskResponseBodyApiResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportOASTaskResponseBodyApiResults) GoString() string {
	return s.String()
}

func (s *DescribeImportOASTaskResponseBodyApiResults) GetApiResult() []*DescribeImportOASTaskResponseBodyApiResultsApiResult {
	return s.ApiResult
}

func (s *DescribeImportOASTaskResponseBodyApiResults) SetApiResult(v []*DescribeImportOASTaskResponseBodyApiResultsApiResult) *DescribeImportOASTaskResponseBodyApiResults {
	s.ApiResult = v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResults) Validate() error {
	return dara.Validate(s)
}

type DescribeImportOASTaskResponseBodyApiResultsApiResult struct {
	// The API ID.
	//
	// example:
	//
	// c5a0c2900ff746b789c007545be22fb8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// GetByCreatorIdUsingGET
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The API description.
	//
	// example:
	//
	// release data api 411055691505041
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cause of the failure if the API fails to be imported.
	//
	// example:
	//
	// Internal Error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The API group ID.
	//
	// example:
	//
	// 736508d885074167ba8fbce3bc95ea0b
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The HTTP request HTTP method of the API.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /creator/getByCreatorId
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The execution status of the subtask. Valid values:
	//
	// 	- RUNNING
	//
	// 	- WAIT
	//
	// 	- OVER
	//
	// 	- FAIL
	//
	// 	- CANCEL
	//
	// example:
	//
	// WAIT
	UpdateStatus *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
}

func (s DescribeImportOASTaskResponseBodyApiResultsApiResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportOASTaskResponseBodyApiResultsApiResult) GoString() string {
	return s.String()
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetMethod() *string {
	return s.Method
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetPath() *string {
	return s.Path
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetApiId(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.ApiId = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetApiName(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.ApiName = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetDescription(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.Description = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetErrorMessage(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetGroupId(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.GroupId = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetMethod(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.Method = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetPath(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.Path = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) SetUpdateStatus(v string) *DescribeImportOASTaskResponseBodyApiResultsApiResult {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyApiResultsApiResult) Validate() error {
	return dara.Validate(s)
}

type DescribeImportOASTaskResponseBodyModelResults struct {
	ModelResult []*DescribeImportOASTaskResponseBodyModelResultsModelResult `json:"ModelResult,omitempty" xml:"ModelResult,omitempty" type:"Repeated"`
}

func (s DescribeImportOASTaskResponseBodyModelResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportOASTaskResponseBodyModelResults) GoString() string {
	return s.String()
}

func (s *DescribeImportOASTaskResponseBodyModelResults) GetModelResult() []*DescribeImportOASTaskResponseBodyModelResultsModelResult {
	return s.ModelResult
}

func (s *DescribeImportOASTaskResponseBodyModelResults) SetModelResult(v []*DescribeImportOASTaskResponseBodyModelResultsModelResult) *DescribeImportOASTaskResponseBodyModelResults {
	s.ModelResult = v
	return s
}

func (s *DescribeImportOASTaskResponseBodyModelResults) Validate() error {
	return dara.Validate(s)
}

type DescribeImportOASTaskResponseBodyModelResultsModelResult struct {
	// The cause of the failure if the model fails to be imported.
	//
	// example:
	//
	// Internal Error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The API group ID.
	//
	// example:
	//
	// 736508d885074167ba8fbce3bc95ea0b
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the imported model.
	//
	// example:
	//
	// 6b48d724c921415486e190c494dd6bf8
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The model name.
	//
	// example:
	//
	// Pet
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The execution status of the subtask. Valid values:
	//
	// 	- RUNNING
	//
	// 	- WAIT
	//
	// 	- OVER
	//
	// 	- FAIL
	//
	// 	- CANCEL
	//
	// example:
	//
	// FAIL
	UpdateStatus *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
}

func (s DescribeImportOASTaskResponseBodyModelResultsModelResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportOASTaskResponseBodyModelResultsModelResult) GoString() string {
	return s.String()
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) GetModelId() *string {
	return s.ModelId
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) SetErrorMessage(v string) *DescribeImportOASTaskResponseBodyModelResultsModelResult {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) SetGroupId(v string) *DescribeImportOASTaskResponseBodyModelResultsModelResult {
	s.GroupId = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) SetModelId(v string) *DescribeImportOASTaskResponseBodyModelResultsModelResult {
	s.ModelId = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) SetModelName(v string) *DescribeImportOASTaskResponseBodyModelResultsModelResult {
	s.ModelName = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) SetUpdateStatus(v string) *DescribeImportOASTaskResponseBodyModelResultsModelResult {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeImportOASTaskResponseBodyModelResultsModelResult) Validate() error {
	return dara.Validate(s)
}
