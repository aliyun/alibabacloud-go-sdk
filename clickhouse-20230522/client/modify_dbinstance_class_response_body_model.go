// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDBInstanceClassResponseBodyData) *ModifyDBInstanceClassResponseBody
	GetData() *ModifyDBInstanceClassResponseBodyData
	SetRequestId(v string) *ModifyDBInstanceClassResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceClassResponseBody struct {
	// The returned result.
	Data *ModifyDBInstanceClassResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponseBody) GetData() *ModifyDBInstanceClassResponseBodyData {
	return s.Data
}

func (s *ModifyDBInstanceClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceClassResponseBody) SetData(v *ModifyDBInstanceClassResponseBodyData) *ModifyDBInstanceClassResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceClassResponseBody) SetRequestId(v string) *ModifyDBInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceClassResponseBodyData struct {
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceID *int64 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 2
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10000****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBInstanceClassResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponseBodyData) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *ModifyDBInstanceClassResponseBodyData) GetDBInstanceID() *int64 {
	return s.DBInstanceID
}

func (s *ModifyDBInstanceClassResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceClassResponseBodyData) GetScaleMax() *int64 {
	return s.ScaleMax
}

func (s *ModifyDBInstanceClassResponseBodyData) GetScaleMin() *int64 {
	return s.ScaleMin
}

func (s *ModifyDBInstanceClassResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifyDBInstanceClassResponseBodyData) SetComputingGroupId(v string) *ModifyDBInstanceClassResponseBodyData {
	s.ComputingGroupId = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetDBInstanceID(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetDBInstanceName(v string) *ModifyDBInstanceClassResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetScaleMax(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.ScaleMax = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetScaleMin(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.ScaleMin = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) SetTaskId(v int64) *ModifyDBInstanceClassResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBodyData) Validate() error {
	return dara.Validate(s)
}
