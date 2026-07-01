// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOperationRecordResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListOperationRecordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListOperationRecordResponseBody
	GetMessage() *string
	SetOperationLogListResponse(v *ListOperationRecordResponseBodyOperationLogListResponse) *ListOperationRecordResponseBody
	GetOperationLogListResponse() *ListOperationRecordResponseBodyOperationLogListResponse
	SetRequestId(v string) *ListOperationRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOperationRecordResponseBody
	GetSuccess() *bool
}

type ListOperationRecordResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The operation log list response.
	OperationLogListResponse *ListOperationRecordResponseBodyOperationLogListResponse `json:"OperationLogListResponse,omitempty" xml:"OperationLogListResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOperationRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOperationRecordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOperationRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOperationRecordResponseBody) GetOperationLogListResponse() *ListOperationRecordResponseBodyOperationLogListResponse {
	return s.OperationLogListResponse
}

func (s *ListOperationRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOperationRecordResponseBody) SetCode(v string) *ListOperationRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListOperationRecordResponseBody) SetHttpStatusCode(v int32) *ListOperationRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOperationRecordResponseBody) SetMessage(v string) *ListOperationRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ListOperationRecordResponseBody) SetOperationLogListResponse(v *ListOperationRecordResponseBodyOperationLogListResponse) *ListOperationRecordResponseBody {
	s.OperationLogListResponse = v
	return s
}

func (s *ListOperationRecordResponseBody) SetRequestId(v string) *ListOperationRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationRecordResponseBody) SetSuccess(v bool) *ListOperationRecordResponseBody {
	s.Success = &v
	return s
}

func (s *ListOperationRecordResponseBody) Validate() error {
	if s.OperationLogListResponse != nil {
		if err := s.OperationLogListResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOperationRecordResponseBodyOperationLogListResponse struct {
	// The total number of records.
	//
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of operation logs.
	ResultData []*ListOperationRecordResponseBodyOperationLogListResponseResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Repeated"`
}

func (s ListOperationRecordResponseBodyOperationLogListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationRecordResponseBodyOperationLogListResponse) GoString() string {
	return s.String()
}

func (s *ListOperationRecordResponseBodyOperationLogListResponse) GetCount() *int32 {
	return s.Count
}

func (s *ListOperationRecordResponseBodyOperationLogListResponse) GetResultData() []*ListOperationRecordResponseBodyOperationLogListResponseResultData {
	return s.ResultData
}

func (s *ListOperationRecordResponseBodyOperationLogListResponse) SetCount(v int32) *ListOperationRecordResponseBodyOperationLogListResponse {
	s.Count = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponse) SetResultData(v []*ListOperationRecordResponseBodyOperationLogListResponseResultData) *ListOperationRecordResponseBodyOperationLogListResponse {
	s.ResultData = v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponse) Validate() error {
	if s.ResultData != nil {
		for _, item := range s.ResultData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationRecordResponseBodyOperationLogListResponseResultData struct {
	// The start time.
	//
	// example:
	//
	// 2025-01-15 10:30:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The code type.
	//
	// example:
	//
	// 0
	CodeType *int32 `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// The execution duration, in milliseconds.
	//
	// example:
	//
	// 120
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// 测试任务
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type.
	//
	// example:
	//
	// onedata-ide
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The operation record ID.
	//
	// example:
	//
	// 987654321
	OperationId *int64 `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of related tables.
	RelationTables []*string `json:"RelationTables,omitempty" xml:"RelationTables,omitempty" type:"Repeated"`
	// The ID of the executor.
	//
	// example:
	//
	// 30231123
	Runner *string `json:"Runner,omitempty" xml:"Runner,omitempty"`
	// The name of the executor.
	//
	// example:
	//
	// 张三
	RunnerName *string `json:"RunnerName,omitempty" xml:"RunnerName,omitempty"`
	// The task status.
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListOperationRecordResponseBodyOperationLogListResponseResultData) String() string {
	return dara.Prettify(s)
}

func (s ListOperationRecordResponseBodyOperationLogListResponseResultData) GoString() string {
	return s.String()
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetCodeType() *int32 {
	return s.CodeType
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetDuration() *int64 {
	return s.Duration
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetId() *int64 {
	return s.Id
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetName() *string {
	return s.Name
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetOperationId() *int64 {
	return s.OperationId
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetRelationTables() []*string {
	return s.RelationTables
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetRunner() *string {
	return s.Runner
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetRunnerName() *string {
	return s.RunnerName
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetStatus() *int32 {
	return s.Status
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetBeginTime(v string) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.BeginTime = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetCodeType(v int32) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.CodeType = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetDuration(v int64) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.Duration = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetId(v int64) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.Id = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetName(v string) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.Name = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetObjectType(v string) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.ObjectType = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetOperationId(v int64) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.OperationId = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetProjectId(v int64) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.ProjectId = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetRelationTables(v []*string) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.RelationTables = v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetRunner(v string) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.Runner = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetRunnerName(v string) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.RunnerName = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetStatus(v int32) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.Status = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) SetTenantId(v int64) *ListOperationRecordResponseBodyOperationLogListResponseResultData {
	s.TenantId = &v
	return s
}

func (s *ListOperationRecordResponseBodyOperationLogListResponseResultData) Validate() error {
	return dara.Validate(s)
}
