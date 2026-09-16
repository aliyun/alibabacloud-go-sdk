// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBusinessHoursRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CheckBusinessHoursRequest
	GetInstanceId() *string
	SetTime(v int64) *CheckBusinessHoursRequest
	GetTime() *int64
}

type CheckBusinessHoursRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1789526665860
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s CheckBusinessHoursRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckBusinessHoursRequest) GoString() string {
	return s.String()
}

func (s *CheckBusinessHoursRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckBusinessHoursRequest) GetTime() *int64 {
	return s.Time
}

func (s *CheckBusinessHoursRequest) SetInstanceId(v string) *CheckBusinessHoursRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckBusinessHoursRequest) SetTime(v int64) *CheckBusinessHoursRequest {
	s.Time = &v
	return s
}

func (s *CheckBusinessHoursRequest) Validate() error {
	return dara.Validate(s)
}
