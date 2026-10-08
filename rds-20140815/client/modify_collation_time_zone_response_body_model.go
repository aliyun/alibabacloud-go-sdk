// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCollationTimeZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollation(v string) *ModifyCollationTimeZoneResponseBody
	GetCollation() *string
	SetDBInstanceId(v string) *ModifyCollationTimeZoneResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *ModifyCollationTimeZoneResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyCollationTimeZoneResponseBody
	GetTaskId() *string
	SetTimezone(v string) *ModifyCollationTimeZoneResponseBody
	GetTimezone() *string
}

type ModifyCollationTimeZoneResponseBody struct {
	// The system character set collation.
	//
	// example:
	//
	// Chinese_PRC_CS_AS
	Collation *string `json:"Collation,omitempty" xml:"Collation,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-bp15qi0nd1u27****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58D48758-F035-52D3-A4FB-80C73DA3E95C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 56365****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The time zone.
	//
	// example:
	//
	// China Standard Time
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s ModifyCollationTimeZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCollationTimeZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCollationTimeZoneResponseBody) GetCollation() *string {
	return s.Collation
}

func (s *ModifyCollationTimeZoneResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyCollationTimeZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCollationTimeZoneResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyCollationTimeZoneResponseBody) GetTimezone() *string {
	return s.Timezone
}

func (s *ModifyCollationTimeZoneResponseBody) SetCollation(v string) *ModifyCollationTimeZoneResponseBody {
	s.Collation = &v
	return s
}

func (s *ModifyCollationTimeZoneResponseBody) SetDBInstanceId(v string) *ModifyCollationTimeZoneResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyCollationTimeZoneResponseBody) SetRequestId(v string) *ModifyCollationTimeZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCollationTimeZoneResponseBody) SetTaskId(v string) *ModifyCollationTimeZoneResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyCollationTimeZoneResponseBody) SetTimezone(v string) *ModifyCollationTimeZoneResponseBody {
	s.Timezone = &v
	return s
}

func (s *ModifyCollationTimeZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
