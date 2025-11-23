// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveDataAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSensitiveDataAuditLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSensitiveDataAuditLogResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSensitiveDataAuditLogResponseBody
	GetRequestId() *string
	SetSensitiveDataAuditLogList(v []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) *ListSensitiveDataAuditLogResponseBody
	GetSensitiveDataAuditLogList() []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList
	SetSuccess(v bool) *ListSensitiveDataAuditLogResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSensitiveDataAuditLogResponseBody
	GetTotalCount() *int64
}

type ListSensitiveDataAuditLogResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The audit logs for sensitive data.
	SensitiveDataAuditLogList []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList `json:"SensitiveDataAuditLogList,omitempty" xml:"SensitiveDataAuditLogList,omitempty" type:"Repeated"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSensitiveDataAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveDataAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveDataAuditLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSensitiveDataAuditLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSensitiveDataAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSensitiveDataAuditLogResponseBody) GetSensitiveDataAuditLogList() []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	return s.SensitiveDataAuditLogList
}

func (s *ListSensitiveDataAuditLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSensitiveDataAuditLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSensitiveDataAuditLogResponseBody) SetErrorCode(v string) *ListSensitiveDataAuditLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBody) SetErrorMessage(v string) *ListSensitiveDataAuditLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBody) SetRequestId(v string) *ListSensitiveDataAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBody) SetSensitiveDataAuditLogList(v []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) *ListSensitiveDataAuditLogResponseBody {
	s.SensitiveDataAuditLogList = v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBody) SetSuccess(v bool) *ListSensitiveDataAuditLogResponseBody {
	s.Success = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBody) SetTotalCount(v int64) *ListSensitiveDataAuditLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBody) Validate() error {
	if s.SensitiveDataAuditLogList != nil {
		for _, item := range s.SensitiveDataAuditLogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList struct {
	// The name of the database that stores the sensitive data.
	//
	// example:
	//
	// ExampleDbName@xxx.xxx.xxx.xxx:3306
	DbDisplayName *string `json:"DbDisplayName,omitempty" xml:"DbDisplayName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 12****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the function module whose audit logs were queried.
	//
	// example:
	//
	// SQL_CONSOLE
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The time when the operation was performed. The time is in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-11-18 10:01:00
	OpTime *string `json:"OpTime,omitempty" xml:"OpTime,omitempty"`
	// The logs for sensitive data.
	SensitiveDataLog []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog `json:"SensitiveDataLog,omitempty" xml:"SensitiveDataLog,omitempty" type:"Repeated"`
	// The details of the object on which the operation was performed. The value of this parameter is in one of the following formats:
	//
	// 	- Object name - object ID
	//
	// 	- Object name (object ID)
	//
	// example:
	//
	// Ticket - 1\\*\\*\\*\\*
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The user ID of the requester.
	//
	// example:
	//
	// 1**
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the requester.
	//
	// example:
	//
	// ExampleUserName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GoString() string {
	return s.String()
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetDbDisplayName() *string {
	return s.DbDisplayName
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetOpTime() *string {
	return s.OpTime
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetSensitiveDataLog() []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog {
	return s.SensitiveDataLog
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetUserId() *int64 {
	return s.UserId
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) GetUserName() *string {
	return s.UserName
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetDbDisplayName(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.DbDisplayName = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetInstanceId(v int64) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.InstanceId = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetModuleName(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.ModuleName = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetOpTime(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.OpTime = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetSensitiveDataLog(v []*ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.SensitiveDataLog = v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetTargetName(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.TargetName = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetUserId(v int64) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.UserId = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) SetUserName(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList {
	s.UserName = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogList) Validate() error {
	if s.SensitiveDataLog != nil {
		for _, item := range s.SensitiveDataLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog struct {
	// The name of the column that contains sensitive data.
	//
	// example:
	//
	// ExampleColumnName
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The permission that the user has on the column. Valid values:
	//
	// 	- **No permission**
	//
	// 	- **Partial redaction**
	//
	// 	- **Plaintext**
	//
	// 	- **Change**
	//
	// 	- **Enable data masking**
	//
	// 	- **Disable data masking**
	//
	// example:
	//
	// Change
	ColumnPermissionType *string `json:"ColumnPermissionType,omitempty" xml:"ColumnPermissionType,omitempty"`
	// The algorithm used for data masking.
	//
	// example:
	//
	// Default - Full redaction
	DesensitizationRule *string `json:"DesensitizationRule,omitempty" xml:"DesensitizationRule,omitempty"`
	// The sensitivity level of the data. Valid values:
	//
	// 	- **Low**
	//
	// 	- **Medium**
	//
	// 	- **High**
	//
	// example:
	//
	// Low
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The name of the table that stores the sensitive data.
	//
	// example:
	//
	// ExampleTableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) GoString() string {
	return s.String()
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) GetColumnPermissionType() *string {
	return s.ColumnPermissionType
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) GetDesensitizationRule() *string {
	return s.DesensitizationRule
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) SetColumnName(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) SetColumnPermissionType(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog {
	s.ColumnPermissionType = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) SetDesensitizationRule(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog {
	s.DesensitizationRule = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) SetSecurityLevel(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog {
	s.SecurityLevel = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) SetTableName(v string) *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog {
	s.TableName = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponseBodySensitiveDataAuditLogListSensitiveDataLog) Validate() error {
	return dara.Validate(s)
}
