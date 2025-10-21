// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartDBInstanceResponseBodyData) *StartDBInstanceResponseBody
	GetData() *StartDBInstanceResponseBodyData
	SetRequestId(v string) *StartDBInstanceResponseBody
	GetRequestId() *string
}

type StartDBInstanceResponseBody struct {
	// The data returned.
	Data *StartDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDBInstanceResponseBody) GetData() *StartDBInstanceResponseBodyData {
	return s.Data
}

func (s *StartDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDBInstanceResponseBody) SetData(v *StartDBInstanceResponseBodyData) *StartDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StartDBInstanceResponseBody) SetRequestId(v string) *StartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDBInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartDBInstanceResponseBodyData struct {
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
	// 100000837
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartDBInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartDBInstanceResponseBodyData) GetDBInstanceID() *int64 {
	return s.DBInstanceID
}

func (s *StartDBInstanceResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *StartDBInstanceResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StartDBInstanceResponseBodyData) SetDBInstanceID(v int64) *StartDBInstanceResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *StartDBInstanceResponseBodyData) SetDBInstanceName(v string) *StartDBInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *StartDBInstanceResponseBodyData) SetTaskId(v int64) *StartDBInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *StartDBInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
