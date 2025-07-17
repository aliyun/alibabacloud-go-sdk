// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceLoginAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListInstanceLoginAuditLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListInstanceLoginAuditLogResponseBody
	GetErrorMessage() *string
	SetInstanceLoginAuditLogList(v *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList) *ListInstanceLoginAuditLogResponseBody
	GetInstanceLoginAuditLogList() *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList
	SetRequestId(v string) *ListInstanceLoginAuditLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceLoginAuditLogResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListInstanceLoginAuditLogResponseBody
	GetTotalCount() *int64
}

type ListInstanceLoginAuditLogResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidPageSize
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Specified parameter PageSize is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The logon records of the instance.
	InstanceLoginAuditLogList *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList `json:"InstanceLoginAuditLogList,omitempty" xml:"InstanceLoginAuditLogList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 074CE7C9-4F9C-5B62-89BC-7B4914A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s ListInstanceLoginAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLoginAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceLoginAuditLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInstanceLoginAuditLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInstanceLoginAuditLogResponseBody) GetInstanceLoginAuditLogList() *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList {
	return s.InstanceLoginAuditLogList
}

func (s *ListInstanceLoginAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceLoginAuditLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceLoginAuditLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceLoginAuditLogResponseBody) SetErrorCode(v string) *ListInstanceLoginAuditLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBody) SetErrorMessage(v string) *ListInstanceLoginAuditLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBody) SetInstanceLoginAuditLogList(v *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList) *ListInstanceLoginAuditLogResponseBody {
	s.InstanceLoginAuditLogList = v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBody) SetRequestId(v string) *ListInstanceLoginAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBody) SetSuccess(v bool) *ListInstanceLoginAuditLogResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBody) SetTotalCount(v int64) *ListInstanceLoginAuditLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList struct {
	InstanceLoginAuditLog []*ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog `json:"InstanceLoginAuditLog,omitempty" xml:"InstanceLoginAuditLog,omitempty" type:"Repeated"`
}

func (s ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList) GoString() string {
	return s.String()
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList) GetInstanceLoginAuditLog() []*ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	return s.InstanceLoginAuditLog
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList) SetInstanceLoginAuditLog(v []*ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList {
	s.InstanceLoginAuditLog = v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogList) Validate() error {
	return dara.Validate(s)
}

type ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog struct {
	// The database account that is used to log on to the instance.
	//
	// example:
	//
	// test_User
	DbUser *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 177****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// rm-bp144d5ky4l4rli0417****.mysql.rds.aliyuncs.com:3306[rm-bp144d5ky4l4r****]
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The time when the user performed an operation on the instance.
	//
	// example:
	//
	// 2021-11-18 11:13:26
	OpTime *string `json:"OpTime,omitempty" xml:"OpTime,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 117.36.XX.XX,100.104.XX.XX
	RequestIp *string `json:"RequestIp,omitempty" xml:"RequestIp,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 12****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The alias of the user.
	//
	// example:
	//
	// test_UserName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GoString() string {
	return s.String()
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GetDbUser() *string {
	return s.DbUser
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GetOpTime() *string {
	return s.OpTime
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GetRequestIp() *string {
	return s.RequestIp
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GetUserId() *int64 {
	return s.UserId
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) GetUserName() *string {
	return s.UserName
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) SetDbUser(v string) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	s.DbUser = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) SetInstanceId(v int64) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) SetInstanceName(v string) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) SetOpTime(v string) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	s.OpTime = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) SetRequestIp(v string) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	s.RequestIp = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) SetUserId(v int64) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	s.UserId = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) SetUserName(v string) *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog {
	s.UserName = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponseBodyInstanceLoginAuditLogListInstanceLoginAuditLog) Validate() error {
	return dara.Validate(s)
}
