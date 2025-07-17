// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v int64) *CloseAlarmRequest
	GetAlarmId() *int64
	SetHandlerId(v int64) *CloseAlarmRequest
	GetHandlerId() *int64
	SetRegionId(v string) *CloseAlarmRequest
	GetRegionId() *string
	SetSolution(v string) *CloseAlarmRequest
	GetSolution() *string
}

type CloseAlarmRequest struct {
	// The ID of the alert.
	//
	// For more information about how to obtain the ID of an alert, see [ListAlertEvents](https://help.aliyun.com/document_detail/2612346.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 163
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The ID of the alert handler.
	//
	// example:
	//
	// 2048065
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The alert solution.
	//
	// example:
	//
	// Restart Repair
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s CloseAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseAlarmRequest) GoString() string {
	return s.String()
}

func (s *CloseAlarmRequest) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *CloseAlarmRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *CloseAlarmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseAlarmRequest) GetSolution() *string {
	return s.Solution
}

func (s *CloseAlarmRequest) SetAlarmId(v int64) *CloseAlarmRequest {
	s.AlarmId = &v
	return s
}

func (s *CloseAlarmRequest) SetHandlerId(v int64) *CloseAlarmRequest {
	s.HandlerId = &v
	return s
}

func (s *CloseAlarmRequest) SetRegionId(v string) *CloseAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *CloseAlarmRequest) SetSolution(v string) *CloseAlarmRequest {
	s.Solution = &v
	return s
}

func (s *CloseAlarmRequest) Validate() error {
	return dara.Validate(s)
}
