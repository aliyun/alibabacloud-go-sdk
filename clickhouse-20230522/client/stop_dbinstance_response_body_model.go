// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StopDBInstanceResponseBodyData) *StopDBInstanceResponseBody
	GetData() *StopDBInstanceResponseBodyData
	SetRequestId(v string) *StopDBInstanceResponseBody
	GetRequestId() *string
}

type StopDBInstanceResponseBody struct {
	// The data returned.
	Data *StopDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDBInstanceResponseBody) GetData() *StopDBInstanceResponseBodyData {
	return s.Data
}

func (s *StopDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDBInstanceResponseBody) SetData(v *StopDBInstanceResponseBodyData) *StopDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StopDBInstanceResponseBody) SetRequestId(v string) *StopDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDBInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopDBInstanceResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int64 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test1
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 100000785
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopDBInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StopDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopDBInstanceResponseBodyData) GetDBInstanceID() *int64 {
	return s.DBInstanceID
}

func (s *StopDBInstanceResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *StopDBInstanceResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StopDBInstanceResponseBodyData) SetDBInstanceID(v int64) *StopDBInstanceResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *StopDBInstanceResponseBodyData) SetDBInstanceName(v string) *StopDBInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *StopDBInstanceResponseBodyData) SetTaskId(v int64) *StopDBInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *StopDBInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
