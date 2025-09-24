// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResellerUserAlarmThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmType(v string) *QueryResellerUserAlarmThresholdRequest
	GetAlarmType() *string
	SetOwnerId(v int64) *QueryResellerUserAlarmThresholdRequest
	GetOwnerId() *int64
}

type QueryResellerUserAlarmThresholdRequest struct {
	// example:
	//
	// quota_low_balance
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryResellerUserAlarmThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryResellerUserAlarmThresholdRequest) GoString() string {
	return s.String()
}

func (s *QueryResellerUserAlarmThresholdRequest) GetAlarmType() *string {
	return s.AlarmType
}

func (s *QueryResellerUserAlarmThresholdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryResellerUserAlarmThresholdRequest) SetAlarmType(v string) *QueryResellerUserAlarmThresholdRequest {
	s.AlarmType = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdRequest) SetOwnerId(v int64) *QueryResellerUserAlarmThresholdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdRequest) Validate() error {
	return dara.Validate(s)
}
