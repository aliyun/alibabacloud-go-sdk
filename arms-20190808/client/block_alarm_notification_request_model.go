// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockAlarmNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v int64) *BlockAlarmNotificationRequest
	GetAlarmId() *int64
	SetHandlerId(v int64) *BlockAlarmNotificationRequest
	GetHandlerId() *int64
	SetRegionId(v string) *BlockAlarmNotificationRequest
	GetRegionId() *string
	SetTimeout(v int64) *BlockAlarmNotificationRequest
	GetTimeout() *int64
}

type BlockAlarmNotificationRequest struct {
	// The ID of the alert.
	//
	// For more information about how to obtain the ID of an alert, see [ListAlertEvents](https://help.aliyun.com/document_detail/2612346.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 133
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The ID of the alert handler.
	//
	// example:
	//
	// 2044049
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of seconds that elapse before alert notifications are blocked. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 180
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s BlockAlarmNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s BlockAlarmNotificationRequest) GoString() string {
	return s.String()
}

func (s *BlockAlarmNotificationRequest) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *BlockAlarmNotificationRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *BlockAlarmNotificationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BlockAlarmNotificationRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *BlockAlarmNotificationRequest) SetAlarmId(v int64) *BlockAlarmNotificationRequest {
	s.AlarmId = &v
	return s
}

func (s *BlockAlarmNotificationRequest) SetHandlerId(v int64) *BlockAlarmNotificationRequest {
	s.HandlerId = &v
	return s
}

func (s *BlockAlarmNotificationRequest) SetRegionId(v string) *BlockAlarmNotificationRequest {
	s.RegionId = &v
	return s
}

func (s *BlockAlarmNotificationRequest) SetTimeout(v int64) *BlockAlarmNotificationRequest {
	s.Timeout = &v
	return s
}

func (s *BlockAlarmNotificationRequest) Validate() error {
	return dara.Validate(s)
}
