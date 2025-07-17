// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v int64) *ClaimAlarmRequest
	GetAlarmId() *int64
	SetHandlerId(v int64) *ClaimAlarmRequest
	GetHandlerId() *int64
	SetRegionId(v string) *ClaimAlarmRequest
	GetRegionId() *string
}

type ClaimAlarmRequest struct {
	// The ID of the alert.
	//
	// For more information about how to obtain the ID of an alert, see [ListAlertEvents](https://help.aliyun.com/document_detail/2612346.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The ID of the handler.
	//
	// example:
	//
	// 2046076
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ClaimAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s ClaimAlarmRequest) GoString() string {
	return s.String()
}

func (s *ClaimAlarmRequest) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *ClaimAlarmRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *ClaimAlarmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ClaimAlarmRequest) SetAlarmId(v int64) *ClaimAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *ClaimAlarmRequest) SetHandlerId(v int64) *ClaimAlarmRequest {
	s.HandlerId = &v
	return s
}

func (s *ClaimAlarmRequest) SetRegionId(v string) *ClaimAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *ClaimAlarmRequest) Validate() error {
	return dara.Validate(s)
}
