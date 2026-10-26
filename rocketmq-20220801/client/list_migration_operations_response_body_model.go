// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMigrationOperationsResponseBody
	GetCode() *string
	SetData(v *ListMigrationOperationsResponseBodyData) *ListMigrationOperationsResponseBody
	GetData() *ListMigrationOperationsResponseBodyData
	SetDynamicCode(v string) *ListMigrationOperationsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListMigrationOperationsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListMigrationOperationsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMigrationOperationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMigrationOperationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMigrationOperationsResponseBody
	GetSuccess() *bool
}

type ListMigrationOperationsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListMigrationOperationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// consumerGroupId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The topic already exists.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7E8AE3A-219B-52EE-BE32-4036F5F88833
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMigrationOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMigrationOperationsResponseBody) GetData() *ListMigrationOperationsResponseBodyData {
	return s.Data
}

func (s *ListMigrationOperationsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListMigrationOperationsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListMigrationOperationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMigrationOperationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMigrationOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMigrationOperationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMigrationOperationsResponseBody) SetCode(v string) *ListMigrationOperationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMigrationOperationsResponseBody) SetData(v *ListMigrationOperationsResponseBodyData) *ListMigrationOperationsResponseBody {
	s.Data = v
	return s
}

func (s *ListMigrationOperationsResponseBody) SetDynamicCode(v string) *ListMigrationOperationsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListMigrationOperationsResponseBody) SetDynamicMessage(v string) *ListMigrationOperationsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListMigrationOperationsResponseBody) SetHttpStatusCode(v int32) *ListMigrationOperationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMigrationOperationsResponseBody) SetMessage(v string) *ListMigrationOperationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMigrationOperationsResponseBody) SetRequestId(v string) *ListMigrationOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMigrationOperationsResponseBody) SetSuccess(v bool) *ListMigrationOperationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMigrationOperationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMigrationOperationsResponseBodyData struct {
	// The paginated data.
	List []*ListMigrationOperationsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMigrationOperationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsResponseBodyData) GetList() []*ListMigrationOperationsResponseBodyDataList {
	return s.List
}

func (s *ListMigrationOperationsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListMigrationOperationsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMigrationOperationsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMigrationOperationsResponseBodyData) SetList(v []*ListMigrationOperationsResponseBodyDataList) *ListMigrationOperationsResponseBodyData {
	s.List = v
	return s
}

func (s *ListMigrationOperationsResponseBodyData) SetPageNumber(v int64) *ListMigrationOperationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyData) SetPageSize(v int64) *ListMigrationOperationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyData) SetTotalCount(v int64) *ListMigrationOperationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMigrationOperationsResponseBodyDataList struct {
	// The operation business status.
	//
	// example:
	//
	// DOUBLE_READ
	BusinessStatus *string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
	// The creation time, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The migration ID.
	//
	// example:
	//
	// 1
	MigrationId *int64 `json:"migrationId,omitempty" xml:"migrationId,omitempty"`
	// The migration operation ID.
	//
	// example:
	//
	// 183
	OperationId *int64 `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The primary key of the operation object.
	//
	//
	// >
	//
	// > - During the message migration stage, the operation object is the topic name.
	//
	// example:
	//
	// normal-topic-1
	OperationKey *string `json:"operationKey,omitempty" xml:"operationKey,omitempty"`
	// The operation parameter.
	OperationParam *ListMigrationOperationsResponseBodyDataListOperationParam `json:"operationParam,omitempty" xml:"operationParam,omitempty" type:"Struct"`
	// The operation result.
	OperationResult *ListMigrationOperationsResponseBodyDataListOperationResult `json:"operationResult,omitempty" xml:"operationResult,omitempty" type:"Struct"`
	// The operation status.
	//
	// example:
	//
	// DOING
	OperationStatus *string `json:"operationStatus,omitempty" xml:"operationStatus,omitempty"`
	// The operation type.
	//
	// example:
	//
	// MIGRATE_MESSAGE
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// The migration stage type.
	//
	// example:
	//
	// MIGRATE_MESSAGE
	StageType *string `json:"stageType,omitempty" xml:"stageType,omitempty"`
	// The update time, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListMigrationOperationsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsResponseBodyDataList) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListMigrationOperationsResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMigrationOperationsResponseBodyDataList) GetMigrationId() *int64 {
	return s.MigrationId
}

func (s *ListMigrationOperationsResponseBodyDataList) GetOperationId() *int64 {
	return s.OperationId
}

func (s *ListMigrationOperationsResponseBodyDataList) GetOperationKey() *string {
	return s.OperationKey
}

func (s *ListMigrationOperationsResponseBodyDataList) GetOperationParam() *ListMigrationOperationsResponseBodyDataListOperationParam {
	return s.OperationParam
}

func (s *ListMigrationOperationsResponseBodyDataList) GetOperationResult() *ListMigrationOperationsResponseBodyDataListOperationResult {
	return s.OperationResult
}

func (s *ListMigrationOperationsResponseBodyDataList) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *ListMigrationOperationsResponseBodyDataList) GetOperationType() *string {
	return s.OperationType
}

func (s *ListMigrationOperationsResponseBodyDataList) GetStageType() *string {
	return s.StageType
}

func (s *ListMigrationOperationsResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMigrationOperationsResponseBodyDataList) SetBusinessStatus(v string) *ListMigrationOperationsResponseBodyDataList {
	s.BusinessStatus = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetCreateTime(v string) *ListMigrationOperationsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetMigrationId(v int64) *ListMigrationOperationsResponseBodyDataList {
	s.MigrationId = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetOperationId(v int64) *ListMigrationOperationsResponseBodyDataList {
	s.OperationId = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetOperationKey(v string) *ListMigrationOperationsResponseBodyDataList {
	s.OperationKey = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetOperationParam(v *ListMigrationOperationsResponseBodyDataListOperationParam) *ListMigrationOperationsResponseBodyDataList {
	s.OperationParam = v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetOperationResult(v *ListMigrationOperationsResponseBodyDataListOperationResult) *ListMigrationOperationsResponseBodyDataList {
	s.OperationResult = v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetOperationStatus(v string) *ListMigrationOperationsResponseBodyDataList {
	s.OperationStatus = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetOperationType(v string) *ListMigrationOperationsResponseBodyDataList {
	s.OperationType = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetStageType(v string) *ListMigrationOperationsResponseBodyDataList {
	s.StageType = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) SetUpdateTime(v string) *ListMigrationOperationsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataList) Validate() error {
	if s.OperationParam != nil {
		if err := s.OperationParam.Validate(); err != nil {
			return err
		}
	}
	if s.OperationResult != nil {
		if err := s.OperationResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMigrationOperationsResponseBodyDataListOperationParam struct {
	// The parameter information.
	//
	// example:
	//
	// {
	//
	//       "action": "MIGRATE_FORWARD",
	//
	//       "targetStage": "MIGRATE_FINISH"
	//
	// }
	ParamData interface{} `json:"paramData,omitempty" xml:"paramData,omitempty"`
}

func (s ListMigrationOperationsResponseBodyDataListOperationParam) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsResponseBodyDataListOperationParam) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsResponseBodyDataListOperationParam) GetParamData() interface{} {
	return s.ParamData
}

func (s *ListMigrationOperationsResponseBodyDataListOperationParam) SetParamData(v interface{}) *ListMigrationOperationsResponseBodyDataListOperationParam {
	s.ParamData = v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataListOperationParam) Validate() error {
	return dara.Validate(s)
}

type ListMigrationOperationsResponseBodyDataListOperationResult struct {
	// The operation result.
	//
	// example:
	//
	// {
	//
	//       "checkStatus": "WAIT_CHECK",
	//
	//       "sourceTopicStatus": "PERM_6",
	//
	//       "targetTopicStatus": "PERM_4",
	//
	//       "checkResult": {},
	//
	//       "migrateStage": "DOUBLE_READ"
	//
	// }
	ResultData interface{} `json:"resultData,omitempty" xml:"resultData,omitempty"`
}

func (s ListMigrationOperationsResponseBodyDataListOperationResult) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsResponseBodyDataListOperationResult) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsResponseBodyDataListOperationResult) GetResultData() interface{} {
	return s.ResultData
}

func (s *ListMigrationOperationsResponseBodyDataListOperationResult) SetResultData(v interface{}) *ListMigrationOperationsResponseBodyDataListOperationResult {
	s.ResultData = v
	return s
}

func (s *ListMigrationOperationsResponseBodyDataListOperationResult) Validate() error {
	return dara.Validate(s)
}
