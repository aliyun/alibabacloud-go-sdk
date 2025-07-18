// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleActiveSQLRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *HandleActiveSQLRecordRequest
	GetDBInstanceId() *string
	SetOperateType(v int32) *HandleActiveSQLRecordRequest
	GetOperateType() *int32
	SetPids(v string) *HandleActiveSQLRecordRequest
	GetPids() *string
}

type HandleActiveSQLRecordRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the operation on the process ID. Valid values:
	//
	// 	- **0**: cancel.
	//
	// 	- **1**: terminate.
	//
	// 	- **2**: forcefully terminate.
	//
	// example:
	//
	// 0
	OperateType *int32 `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The process IDs. A process ID is a unique identifier of a query.
	//
	// This parameter is required.
	//
	// example:
	//
	// "3003925,3003928"
	Pids *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
}

func (s HandleActiveSQLRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleActiveSQLRecordRequest) GoString() string {
	return s.String()
}

func (s *HandleActiveSQLRecordRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *HandleActiveSQLRecordRequest) GetOperateType() *int32 {
	return s.OperateType
}

func (s *HandleActiveSQLRecordRequest) GetPids() *string {
	return s.Pids
}

func (s *HandleActiveSQLRecordRequest) SetDBInstanceId(v string) *HandleActiveSQLRecordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *HandleActiveSQLRecordRequest) SetOperateType(v int32) *HandleActiveSQLRecordRequest {
	s.OperateType = &v
	return s
}

func (s *HandleActiveSQLRecordRequest) SetPids(v string) *HandleActiveSQLRecordRequest {
	s.Pids = &v
	return s
}

func (s *HandleActiveSQLRecordRequest) Validate() error {
	return dara.Validate(s)
}
