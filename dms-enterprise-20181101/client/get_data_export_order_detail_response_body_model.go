// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataExportOrderDetail(v *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) *GetDataExportOrderDetailResponseBody
	GetDataExportOrderDetail() *GetDataExportOrderDetailResponseBodyDataExportOrderDetail
	SetErrorCode(v string) *GetDataExportOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataExportOrderDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataExportOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataExportOrderDetailResponseBody
	GetSuccess() *bool
}

type GetDataExportOrderDetailResponseBody struct {
	// The information about the data export ticket.
	DataExportOrderDetail *GetDataExportOrderDetailResponseBodyDataExportOrderDetail `json:"DataExportOrderDetail,omitempty" xml:"DataExportOrderDetail,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataExportOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBody) GetDataExportOrderDetail() *GetDataExportOrderDetailResponseBodyDataExportOrderDetail {
	return s.DataExportOrderDetail
}

func (s *GetDataExportOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataExportOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataExportOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataExportOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataExportOrderDetailResponseBody) SetDataExportOrderDetail(v *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) *GetDataExportOrderDetailResponseBody {
	s.DataExportOrderDetail = v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetErrorCode(v string) *GetDataExportOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetErrorMessage(v string) *GetDataExportOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetRequestId(v string) *GetDataExportOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetSuccess(v bool) *GetDataExportOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataExportOrderDetailResponseBodyDataExportOrderDetail struct {
	// The status information.
	KeyInfo *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Struct"`
	// The details of the ticket.
	OrderDetail *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail `json:"OrderDetail,omitempty" xml:"OrderDetail,omitempty" type:"Struct"`
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) GetKeyInfo() *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo {
	return s.KeyInfo
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) GetOrderDetail() *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	return s.OrderDetail
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) SetKeyInfo(v *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) *GetDataExportOrderDetailResponseBodyDataExportOrderDetail {
	s.KeyInfo = v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) SetOrderDetail(v *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) *GetDataExportOrderDetailResponseBodyDataExportOrderDetail {
	s.OrderDetail = v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) Validate() error {
	return dara.Validate(s)
}

type GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo struct {
	// Export task ID.
	//
	// example:
	//
	// 1385****
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The state of the data export ticket. Valid values:
	//
	// 	- **PRE_CHECKING**: The ticket was being prechecked.
	//
	// 	- **PRE_CHECK_SUCCESS**: The ticket passed the precheck.
	//
	// 	- **PRE_CHECK_FAIL**: The ticket failed to pass the prechecked.
	//
	// 	- **WAITING_APPLY_AUDIT**: The ticket was to be submitted for approval.
	//
	// 	- **APPLY_AUDIT_SUCCESS**: The ticket was submitted for approval.
	//
	// 	- **ENABLE_EXPORT**: The ticket was approved. Data can be exported.
	//
	// 	- **WAITING_EXPORT**: Data was to be scheduled for export.
	//
	// 	- **DOING_EXPORT**: Data was being exported.
	//
	// 	- **EXPORT_FAIL**: Data failed to be exported.
	//
	// 	- **EXPORT_SUCCESS**: Data was exported.
	//
	// example:
	//
	// EXPORT_SUCCESS
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The precheck ID.
	//
	// example:
	//
	// 123
	PreCheckId *int64 `json:"PreCheckId,omitempty" xml:"PreCheckId,omitempty"`
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) GetJobId() *int64 {
	return s.JobId
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) GetPreCheckId() *int64 {
	return s.PreCheckId
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) SetJobId(v int64) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo {
	s.JobId = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) SetJobStatus(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo {
	s.JobStatus = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) SetPreCheckId(v int64) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo {
	s.PreCheckId = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) Validate() error {
	return dara.Validate(s)
}

type GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail struct {
	// The number of rows that were affected by the SQL statement.
	//
	// example:
	//
	// 1
	ActualAffectRows *int64 `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	// The category of the reason for the data export.
	//
	// example:
	//
	// text
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The name of the database from which data was exported.
	//
	// example:
	//
	// xxx@xxx:3306
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the database from which data was exported.
	//
	// example:
	//
	// 123
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the environment to which the database belongs.
	//
	// example:
	//
	// test
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The SQL statement that was executed to export data.
	//
	// example:
	//
	// select 1
	ExeSQL *string `json:"ExeSQL,omitempty" xml:"ExeSQL,omitempty"`
	// Indicates whether the affected rows are skipped.
	//
	// example:
	//
	// false
	IgnoreAffectRows *bool `json:"IgnoreAffectRows,omitempty" xml:"IgnoreAffectRows,omitempty"`
	// The reason why the affected rows are skipped.
	//
	// example:
	//
	// empty
	IgnoreAffectRowsReason *string `json:"IgnoreAffectRowsReason,omitempty" xml:"IgnoreAffectRowsReason,omitempty"`
	// Indicates whether the database is a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetActualAffectRows() *int64 {
	return s.ActualAffectRows
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetClassify() *string {
	return s.Classify
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetDatabase() *string {
	return s.Database
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetDbId() *int32 {
	return s.DbId
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetExeSQL() *string {
	return s.ExeSQL
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetIgnoreAffectRows() *bool {
	return s.IgnoreAffectRows
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetIgnoreAffectRowsReason() *string {
	return s.IgnoreAffectRowsReason
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GetLogic() *bool {
	return s.Logic
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetActualAffectRows(v int64) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetClassify(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.Classify = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetDatabase(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.Database = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetDbId(v int32) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.DbId = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetEnvType(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.EnvType = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetExeSQL(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.ExeSQL = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetIgnoreAffectRows(v bool) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.IgnoreAffectRows = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetIgnoreAffectRowsReason(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.IgnoreAffectRowsReason = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetLogic(v bool) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.Logic = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) Validate() error {
	return dara.Validate(s)
}
