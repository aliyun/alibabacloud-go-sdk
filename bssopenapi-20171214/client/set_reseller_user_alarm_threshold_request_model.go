// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserAlarmThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmThresholds(v string) *SetResellerUserAlarmThresholdRequest
	GetAlarmThresholds() *string
	SetAlarmType(v string) *SetResellerUserAlarmThresholdRequest
	GetAlarmType() *string
	SetOwnerId(v int64) *SetResellerUserAlarmThresholdRequest
	GetOwnerId() *int64
}

type SetResellerUserAlarmThresholdRequest struct {
	// example:
	//
	// [{\\"denominator\\":100,\\"numerator\\":30,\\"thresholdType\\":1}]
	AlarmThresholds *string `json:"AlarmThresholds,omitempty" xml:"AlarmThresholds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// quota_low_balance
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetResellerUserAlarmThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserAlarmThresholdRequest) GoString() string {
	return s.String()
}

func (s *SetResellerUserAlarmThresholdRequest) GetAlarmThresholds() *string {
	return s.AlarmThresholds
}

func (s *SetResellerUserAlarmThresholdRequest) GetAlarmType() *string {
	return s.AlarmType
}

func (s *SetResellerUserAlarmThresholdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetResellerUserAlarmThresholdRequest) SetAlarmThresholds(v string) *SetResellerUserAlarmThresholdRequest {
	s.AlarmThresholds = &v
	return s
}

func (s *SetResellerUserAlarmThresholdRequest) SetAlarmType(v string) *SetResellerUserAlarmThresholdRequest {
	s.AlarmType = &v
	return s
}

func (s *SetResellerUserAlarmThresholdRequest) SetOwnerId(v int64) *SetResellerUserAlarmThresholdRequest {
	s.OwnerId = &v
	return s
}

func (s *SetResellerUserAlarmThresholdRequest) Validate() error {
	return dara.Validate(s)
}
