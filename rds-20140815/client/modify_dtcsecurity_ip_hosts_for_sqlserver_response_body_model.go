// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDTCSecurityIpHostsForSQLServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody
	GetDBInstanceId() *string
	SetDTCSetResult(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody
	GetDTCSetResult() *string
	SetRequestId(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody
	GetTaskId() *string
}

type ModifyDTCSecurityIpHostsForSQLServerResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The result of the IP address whitelist configuration. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Fail**
	//
	// example:
	//
	// Success
	DTCSetResult *string `json:"DTCSetResult,omitempty" xml:"DTCSetResult,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 671B6D32-B907-4EFF-A3B7-94D2EAD5E3A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 178968983
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDTCSecurityIpHostsForSQLServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDTCSecurityIpHostsForSQLServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) GetDTCSetResult() *string {
	return s.DTCSetResult
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) SetDBInstanceId(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) SetDTCSetResult(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody {
	s.DTCSetResult = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) SetRequestId(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) SetTaskId(v string) *ModifyDTCSecurityIpHostsForSQLServerResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponseBody) Validate() error {
	return dara.Validate(s)
}
