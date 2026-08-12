// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchExportTaskResponseBody
	GetCode() *string
	SetData(v *GetBatchExportTaskResponseBodyData) *GetBatchExportTaskResponseBody
	GetData() *GetBatchExportTaskResponseBodyData
	SetMessage(v string) *GetBatchExportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchExportTaskResponseBody
	GetRequestId() *string
}

type GetBatchExportTaskResponseBody struct {
	// example:
	//
	// Ok
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetBatchExportTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// CE534E1D-FCE4-5930-B784-E055EC1AEE6F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetBatchExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchExportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchExportTaskResponseBody) GetData() *GetBatchExportTaskResponseBodyData {
	return s.Data
}

func (s *GetBatchExportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchExportTaskResponseBody) SetCode(v string) *GetBatchExportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchExportTaskResponseBody) SetData(v *GetBatchExportTaskResponseBodyData) *GetBatchExportTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchExportTaskResponseBody) SetMessage(v string) *GetBatchExportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchExportTaskResponseBody) SetRequestId(v string) *GetBatchExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchExportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchExportTaskResponseBodyData struct {
	// example:
	//
	// 2026-05-15T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// some apis export failed
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 5
	ProcessedCount *int32                                    `json:"processedCount,omitempty" xml:"processedCount,omitempty"`
	Result         *GetBatchExportTaskResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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
	// BatchExport
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetBatchExportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBatchExportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBatchExportTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetBatchExportTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchExportTaskResponseBodyData) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *GetBatchExportTaskResponseBodyData) GetResult() *GetBatchExportTaskResponseBodyDataResult {
	return s.Result
}

func (s *GetBatchExportTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetBatchExportTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetBatchExportTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetBatchExportTaskResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetBatchExportTaskResponseBodyData) SetCreateTime(v string) *GetBatchExportTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) SetErrorMessage(v string) *GetBatchExportTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) SetProcessedCount(v int32) *GetBatchExportTaskResponseBodyData {
	s.ProcessedCount = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) SetResult(v *GetBatchExportTaskResponseBodyDataResult) *GetBatchExportTaskResponseBodyData {
	s.Result = v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) SetStatus(v string) *GetBatchExportTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) SetTaskId(v string) *GetBatchExportTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) SetTaskType(v string) *GetBatchExportTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) SetTotalCount(v int32) *GetBatchExportTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchExportTaskResponseBodyDataResult struct {
	// example:
	//
	// Http
	ApiType      *string                                                 `json:"apiType,omitempty" xml:"apiType,omitempty"`
	FailureItems []*GetBatchExportTaskResponseBodyDataResultFailureItems `json:"failureItems,omitempty" xml:"failureItems,omitempty" type:"Repeated"`
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// UEsDBBQAAAA...
	SpecContentBase64 *string                                                 `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
	SuccessItems      []*GetBatchExportTaskResponseBodyDataResultSuccessItems `json:"successItems,omitempty" xml:"successItems,omitempty" type:"Repeated"`
}

func (s GetBatchExportTaskResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetBatchExportTaskResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetBatchExportTaskResponseBodyDataResult) GetApiType() *string {
	return s.ApiType
}

func (s *GetBatchExportTaskResponseBodyDataResult) GetFailureItems() []*GetBatchExportTaskResponseBodyDataResultFailureItems {
	return s.FailureItems
}

func (s *GetBatchExportTaskResponseBodyDataResult) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetBatchExportTaskResponseBodyDataResult) GetSpecContentBase64() *string {
	return s.SpecContentBase64
}

func (s *GetBatchExportTaskResponseBodyDataResult) GetSuccessItems() []*GetBatchExportTaskResponseBodyDataResultSuccessItems {
	return s.SuccessItems
}

func (s *GetBatchExportTaskResponseBodyDataResult) SetApiType(v string) *GetBatchExportTaskResponseBodyDataResult {
	s.ApiType = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResult) SetFailureItems(v []*GetBatchExportTaskResponseBodyDataResultFailureItems) *GetBatchExportTaskResponseBodyDataResult {
	s.FailureItems = v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResult) SetGatewayId(v string) *GetBatchExportTaskResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResult) SetSpecContentBase64(v string) *GetBatchExportTaskResponseBodyDataResult {
	s.SpecContentBase64 = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResult) SetSuccessItems(v []*GetBatchExportTaskResponseBodyDataResultSuccessItems) *GetBatchExportTaskResponseBodyDataResult {
	s.SuccessItems = v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResult) Validate() error {
	if s.FailureItems != nil {
		for _, item := range s.FailureItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type GetBatchExportTaskResponseBodyDataResultFailureItems struct {
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
	// api definition is invalid
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

func (s GetBatchExportTaskResponseBodyDataResultFailureItems) String() string {
	return dara.Prettify(s)
}

func (s GetBatchExportTaskResponseBodyDataResultFailureItems) GoString() string {
	return s.String()
}

func (s *GetBatchExportTaskResponseBodyDataResultFailureItems) GetApiId() *string {
	return s.ApiId
}

func (s *GetBatchExportTaskResponseBodyDataResultFailureItems) GetApiName() *string {
	return s.ApiName
}

func (s *GetBatchExportTaskResponseBodyDataResultFailureItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchExportTaskResponseBodyDataResultFailureItems) SetApiId(v string) *GetBatchExportTaskResponseBodyDataResultFailureItems {
	s.ApiId = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResultFailureItems) SetApiName(v string) *GetBatchExportTaskResponseBodyDataResultFailureItems {
	s.ApiName = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResultFailureItems) SetErrorMessage(v string) *GetBatchExportTaskResponseBodyDataResultFailureItems {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResultFailureItems) Validate() error {
	return dara.Validate(s)
}

type GetBatchExportTaskResponseBodyDataResultSuccessItems struct {
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
	// api definition is invalid
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

func (s GetBatchExportTaskResponseBodyDataResultSuccessItems) String() string {
	return dara.Prettify(s)
}

func (s GetBatchExportTaskResponseBodyDataResultSuccessItems) GoString() string {
	return s.String()
}

func (s *GetBatchExportTaskResponseBodyDataResultSuccessItems) GetApiId() *string {
	return s.ApiId
}

func (s *GetBatchExportTaskResponseBodyDataResultSuccessItems) GetApiName() *string {
	return s.ApiName
}

func (s *GetBatchExportTaskResponseBodyDataResultSuccessItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchExportTaskResponseBodyDataResultSuccessItems) SetApiId(v string) *GetBatchExportTaskResponseBodyDataResultSuccessItems {
	s.ApiId = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResultSuccessItems) SetApiName(v string) *GetBatchExportTaskResponseBodyDataResultSuccessItems {
	s.ApiName = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResultSuccessItems) SetErrorMessage(v string) *GetBatchExportTaskResponseBodyDataResultSuccessItems {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchExportTaskResponseBodyDataResultSuccessItems) Validate() error {
	return dara.Validate(s)
}
