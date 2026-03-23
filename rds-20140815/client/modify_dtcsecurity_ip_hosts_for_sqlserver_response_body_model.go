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
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DTCSetResult *string `json:"DTCSetResult,omitempty" xml:"DTCSetResult,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
