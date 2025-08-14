// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnalysisExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateAnalysisExportTaskRequest
	GetLang() *string
	SetColumns(v string) *CreateAnalysisExportTaskRequest
	GetColumns() *string
	SetConditions(v string) *CreateAnalysisExportTaskRequest
	GetConditions() *string
	SetEventBeginTime(v int64) *CreateAnalysisExportTaskRequest
	GetEventBeginTime() *int64
	SetEventCodes(v string) *CreateAnalysisExportTaskRequest
	GetEventCodes() *string
	SetEventEndTime(v int64) *CreateAnalysisExportTaskRequest
	GetEventEndTime() *int64
	SetFieldName(v string) *CreateAnalysisExportTaskRequest
	GetFieldName() *string
	SetFieldValue(v string) *CreateAnalysisExportTaskRequest
	GetFieldValue() *string
	SetFileFormat(v string) *CreateAnalysisExportTaskRequest
	GetFileFormat() *string
	SetRegId(v string) *CreateAnalysisExportTaskRequest
	GetRegId() *string
	SetScope(v string) *CreateAnalysisExportTaskRequest
	GetScope() *string
	SetType(v string) *CreateAnalysisExportTaskRequest
	GetType() *string
}

type CreateAnalysisExportTaskRequest struct {
	// Sets the language type for the request and response messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Custom columns
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
	// Query expression
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
	// Start time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1752076800000
	EventBeginTime *int64 `json:"eventBeginTime,omitempty" xml:"eventBeginTime,omitempty"`
	// Event code.
	//
	// This parameter is required.
	//
	// example:
	//
	// de_afghcf6411
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1753891199000
	EventEndTime *int64 `json:"eventEndTime,omitempty" xml:"eventEndTime,omitempty"`
	// Field name
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Field value
	//
	// example:
	//
	// 20
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// File format, Excel, CSV
	//
	// This parameter is required.
	//
	// example:
	//
	// CSV
	FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Export scope: ALL: All, SELECT: Selected rows
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// Type, BASIC: Basic query, ADVANCE: Advanced query, BATCH: Batch query
	//
	// This parameter is required.
	//
	// example:
	//
	// BASIC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAnalysisExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnalysisExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAnalysisExportTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateAnalysisExportTaskRequest) GetColumns() *string {
	return s.Columns
}

func (s *CreateAnalysisExportTaskRequest) GetConditions() *string {
	return s.Conditions
}

func (s *CreateAnalysisExportTaskRequest) GetEventBeginTime() *int64 {
	return s.EventBeginTime
}

func (s *CreateAnalysisExportTaskRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *CreateAnalysisExportTaskRequest) GetEventEndTime() *int64 {
	return s.EventEndTime
}

func (s *CreateAnalysisExportTaskRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateAnalysisExportTaskRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateAnalysisExportTaskRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateAnalysisExportTaskRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateAnalysisExportTaskRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateAnalysisExportTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateAnalysisExportTaskRequest) SetLang(v string) *CreateAnalysisExportTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetColumns(v string) *CreateAnalysisExportTaskRequest {
	s.Columns = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetConditions(v string) *CreateAnalysisExportTaskRequest {
	s.Conditions = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetEventBeginTime(v int64) *CreateAnalysisExportTaskRequest {
	s.EventBeginTime = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetEventCodes(v string) *CreateAnalysisExportTaskRequest {
	s.EventCodes = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetEventEndTime(v int64) *CreateAnalysisExportTaskRequest {
	s.EventEndTime = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetFieldName(v string) *CreateAnalysisExportTaskRequest {
	s.FieldName = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetFieldValue(v string) *CreateAnalysisExportTaskRequest {
	s.FieldValue = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetFileFormat(v string) *CreateAnalysisExportTaskRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetRegId(v string) *CreateAnalysisExportTaskRequest {
	s.RegId = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetScope(v string) *CreateAnalysisExportTaskRequest {
	s.Scope = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) SetType(v string) *CreateAnalysisExportTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateAnalysisExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
