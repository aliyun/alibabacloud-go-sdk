// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMasterNodeShutDownFailOverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MasterNodeShutDownFailOverResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *MasterNodeShutDownFailOverResponseBody
	GetRequestId() *string
	SetTaskID(v string) *MasterNodeShutDownFailOverResponseBody
	GetTaskID() *string
}

type MasterNodeShutDownFailOverResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12123216-4B00-4378-BE4B-08005BFC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 17566
	TaskID *string `json:"TaskID,omitempty" xml:"TaskID,omitempty"`
}

func (s MasterNodeShutDownFailOverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MasterNodeShutDownFailOverResponseBody) GoString() string {
	return s.String()
}

func (s *MasterNodeShutDownFailOverResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MasterNodeShutDownFailOverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MasterNodeShutDownFailOverResponseBody) GetTaskID() *string {
	return s.TaskID
}

func (s *MasterNodeShutDownFailOverResponseBody) SetDBInstanceId(v string) *MasterNodeShutDownFailOverResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *MasterNodeShutDownFailOverResponseBody) SetRequestId(v string) *MasterNodeShutDownFailOverResponseBody {
	s.RequestId = &v
	return s
}

func (s *MasterNodeShutDownFailOverResponseBody) SetTaskID(v string) *MasterNodeShutDownFailOverResponseBody {
	s.TaskID = &v
	return s
}

func (s *MasterNodeShutDownFailOverResponseBody) Validate() error {
	return dara.Validate(s)
}
