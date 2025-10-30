// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnalysisExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAnalysisExportTaskResponseBody
	GetRequestId() *string
	SetResultObject(v *CreateAnalysisExportTaskResponseBodyResultObject) *CreateAnalysisExportTaskResponseBody
	GetResultObject() *CreateAnalysisExportTaskResponseBodyResultObject
}

type CreateAnalysisExportTaskResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *CreateAnalysisExportTaskResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s CreateAnalysisExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAnalysisExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnalysisExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAnalysisExportTaskResponseBody) GetResultObject() *CreateAnalysisExportTaskResponseBodyResultObject {
	return s.ResultObject
}

func (s *CreateAnalysisExportTaskResponseBody) SetRequestId(v string) *CreateAnalysisExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBody) SetResultObject(v *CreateAnalysisExportTaskResponseBodyResultObject) *CreateAnalysisExportTaskResponseBody {
	s.ResultObject = v
	return s
}

func (s *CreateAnalysisExportTaskResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAnalysisExportTaskResponseBodyResultObject struct {
	// Export list.
	//
	// example:
	//
	// [
	//
	//                 {
	//
	//                     "fieldName": "requestId",
	//
	//                     "fieldTitle": "RequestId"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "eventTime",
	//
	//                     "fieldTitle": "事件时间"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "accountId",
	//
	//                     "fieldTitle": "账号"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "deviceId",
	//
	//                     "fieldTitle": "设备ID"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "eventCode",
	//
	//                     "fieldTitle": "事件编码"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "ip",
	//
	//                     "fieldTitle": "IP"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "score",
	//
	//                     "fieldTitle": "分值"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "tags",
	//
	//                     "fieldTitle": "标签"
	//
	//                 },
	//
	//                 {
	//
	//                     "fieldName": "DEtest222",
	//
	//                     "fieldTitle": "测试222"
	//
	//                 }
	//
	//             ]
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// Export task conditions.
	//
	// example:
	//
	// {
	//
	//      "fieldName": null,
	//
	//       "fieldValue": null
	//
	//       }
	Conditions *string `json:"conditions,omitempty" xml:"conditions,omitempty"`
	// Event start time.
	//
	// example:
	//
	// 1752076800000
	EventBeginTime *int64 `json:"eventBeginTime,omitempty" xml:"eventBeginTime,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_ahespg8137
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// End time.
	//
	// example:
	//
	// 1753891199000
	EventEndTime *int64 `json:"eventEndTime,omitempty" xml:"eventEndTime,omitempty"`
	// File format.
	//
	// example:
	//
	// CSV
	FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
	// OSS-generated key.
	//
	// example:
	//
	// xxxxx
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
	// Export task scope.
	//
	// example:
	//
	// ALL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// Task status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Export task type.
	//
	// example:
	//
	// BASIC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User UID
	//
	// example:
	//
	// 1519714049632764
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateAnalysisExportTaskResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CreateAnalysisExportTaskResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetColumns() *string {
	return s.Columns
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetConditions() *string {
	return s.Conditions
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetEventBeginTime() *int64 {
	return s.EventBeginTime
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetEventCodes() *string {
	return s.EventCodes
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetEventEndTime() *int64 {
	return s.EventEndTime
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetScope() *string {
	return s.Scope
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetColumns(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.Columns = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetConditions(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.Conditions = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetEventBeginTime(v int64) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.EventBeginTime = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetEventCodes(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.EventCodes = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetEventEndTime(v int64) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.EventEndTime = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetFileFormat(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.FileFormat = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetOssKey(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.OssKey = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetScope(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.Scope = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetStatus(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetType(v string) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) SetUserId(v int64) *CreateAnalysisExportTaskResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *CreateAnalysisExportTaskResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
