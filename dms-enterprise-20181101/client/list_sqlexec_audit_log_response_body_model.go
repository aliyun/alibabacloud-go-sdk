// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLExecAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSQLExecAuditLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSQLExecAuditLogResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSQLExecAuditLogResponseBody
	GetRequestId() *string
	SetSQLExecAuditLogList(v *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) *ListSQLExecAuditLogResponseBody
	GetSQLExecAuditLogList() *ListSQLExecAuditLogResponseBodySQLExecAuditLogList
	SetSuccess(v bool) *ListSQLExecAuditLogResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSQLExecAuditLogResponseBody
	GetTotalCount() *int64
}

type ListSQLExecAuditLogResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// MissingStartTime
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// StartTime is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 39BC9C86-95AE-58F2-9862-A7C3D896****
	RequestId           *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLExecAuditLogList *ListSQLExecAuditLogResponseBodySQLExecAuditLogList `json:"SQLExecAuditLogList,omitempty" xml:"SQLExecAuditLogList,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSQLExecAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSQLExecAuditLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSQLExecAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSQLExecAuditLogResponseBody) GetSQLExecAuditLogList() *ListSQLExecAuditLogResponseBodySQLExecAuditLogList {
	return s.SQLExecAuditLogList
}

func (s *ListSQLExecAuditLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSQLExecAuditLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSQLExecAuditLogResponseBody) SetErrorCode(v string) *ListSQLExecAuditLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetErrorMessage(v string) *ListSQLExecAuditLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetRequestId(v string) *ListSQLExecAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetSQLExecAuditLogList(v *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) *ListSQLExecAuditLogResponseBody {
	s.SQLExecAuditLogList = v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetSuccess(v bool) *ListSQLExecAuditLogResponseBody {
	s.Success = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetTotalCount(v int64) *ListSQLExecAuditLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) Validate() error {
	if s.SQLExecAuditLogList != nil {
		if err := s.SQLExecAuditLogList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSQLExecAuditLogResponseBodySQLExecAuditLogList struct {
	SQLExecAuditLog []*ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog `json:"SQLExecAuditLog,omitempty" xml:"SQLExecAuditLog,omitempty" type:"Repeated"`
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogList) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogList) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) GetSQLExecAuditLog() []*ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	return s.SQLExecAuditLog
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) SetSQLExecAuditLog(v []*ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) *ListSQLExecAuditLogResponseBodySQLExecAuditLogList {
	s.SQLExecAuditLog = v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) Validate() error {
	if s.SQLExecAuditLog != nil {
		for _, item := range s.SQLExecAuditLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog struct {
	AffectRows   *int64  `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	DbId         *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	ElapsedTime  *int64  `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	ExecState    *string `json:"ExecState,omitempty" xml:"ExecState,omitempty"`
	InstanceId   *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Logic        *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	OpTime       *string `json:"OpTime,omitempty" xml:"OpTime,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SQL          *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	SQLType      *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	SchemaName   *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetDbId() *int64 {
	return s.DbId
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetExecState() *string {
	return s.ExecState
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetLogic() *bool {
	return s.Logic
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetOpTime() *string {
	return s.OpTime
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetRemark() *string {
	return s.Remark
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetSQL() *string {
	return s.SQL
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetSQLType() *string {
	return s.SQLType
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetUserId() *int64 {
	return s.UserId
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetUserName() *string {
	return s.UserName
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetAffectRows(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.AffectRows = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetDbId(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.DbId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetElapsedTime(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.ElapsedTime = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetExecState(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.ExecState = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetInstanceId(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.InstanceId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetInstanceName(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.InstanceName = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetLogic(v bool) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.Logic = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetOpTime(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.OpTime = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetRemark(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.Remark = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetSQL(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.SQL = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetSQLType(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.SQLType = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetSchemaName(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.SchemaName = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetUserId(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.UserId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetUserName(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.UserName = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) Validate() error {
	return dara.Validate(s)
}
