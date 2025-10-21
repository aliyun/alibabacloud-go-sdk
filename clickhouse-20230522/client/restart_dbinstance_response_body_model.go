// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RestartDBInstanceResponseBodyData) *RestartDBInstanceResponseBody
	GetData() *RestartDBInstanceResponseBodyData
	SetRequestId(v string) *RestartDBInstanceResponseBody
	GetRequestId() *string
}

type RestartDBInstanceResponseBody struct {
	// The data returned.
	Data *RestartDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBody) GetData() *RestartDBInstanceResponseBodyData {
	return s.Data
}

func (s *RestartDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBInstanceResponseBody) SetData(v *RestartDBInstanceResponseBodyData) *RestartDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RestartDBInstanceResponseBody) SetRequestId(v string) *RestartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RestartDBInstanceResponseBodyData struct {
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
	// 100001080
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartDBInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBodyData) GetDBInstanceID() *int64 {
	return s.DBInstanceID
}

func (s *RestartDBInstanceResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestartDBInstanceResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *RestartDBInstanceResponseBodyData) SetDBInstanceID(v int64) *RestartDBInstanceResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *RestartDBInstanceResponseBodyData) SetDBInstanceName(v string) *RestartDBInstanceResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *RestartDBInstanceResponseBodyData) SetTaskId(v int64) *RestartDBInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RestartDBInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
