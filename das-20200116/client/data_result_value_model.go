// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataResultValue interface {
	dara.Model
	String() string
	GoString() string
	SetSqlId(v string) *DataResultValue
	GetSqlId() *string
	SetInstanceId(v string) *DataResultValue
	GetInstanceId() *string
	SetCount(v int32) *DataResultValue
	GetCount() *int32
}

type DataResultValue struct {
	// The SQL template ID.
	//
	// example:
	//
	// 61f45ee381b2fa4e8a6545e3bee9
	SqlId *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The number of errors.
	//
	// example:
	//
	// 10
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s DataResultValue) String() string {
	return dara.Prettify(s)
}

func (s DataResultValue) GoString() string {
	return s.String()
}

func (s *DataResultValue) GetSqlId() *string {
	return s.SqlId
}

func (s *DataResultValue) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DataResultValue) GetCount() *int32 {
	return s.Count
}

func (s *DataResultValue) SetSqlId(v string) *DataResultValue {
	s.SqlId = &v
	return s
}

func (s *DataResultValue) SetInstanceId(v string) *DataResultValue {
	s.InstanceId = &v
	return s
}

func (s *DataResultValue) SetCount(v int32) *DataResultValue {
	s.Count = &v
	return s
}

func (s *DataResultValue) Validate() error {
	return dara.Validate(s)
}
