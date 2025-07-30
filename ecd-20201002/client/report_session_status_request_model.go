// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportSessionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *ReportSessionStatusRequest
	GetEndUserId() *string
	SetInstanceId(v string) *ReportSessionStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ReportSessionStatusRequest
	GetRegionId() *string
	SetSessionChangeTime(v int64) *ReportSessionStatusRequest
	GetSessionChangeTime() *int64
	SetSessionId(v string) *ReportSessionStatusRequest
	GetSessionId() *string
	SetSessionStatus(v string) *ReportSessionStatusRequest
	GetSessionStatus() *string
}

type ReportSessionStatusRequest struct {
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-bp167fcodoa90ixn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1642909143781
	SessionChangeTime *int64 `json:"SessionChangeTime,omitempty" xml:"SessionChangeTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SessionLogOn
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
}

func (s ReportSessionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportSessionStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ReportSessionStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReportSessionStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReportSessionStatusRequest) GetSessionChangeTime() *int64 {
	return s.SessionChangeTime
}

func (s *ReportSessionStatusRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ReportSessionStatusRequest) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *ReportSessionStatusRequest) SetEndUserId(v string) *ReportSessionStatusRequest {
	s.EndUserId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetInstanceId(v string) *ReportSessionStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetRegionId(v string) *ReportSessionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionChangeTime(v int64) *ReportSessionStatusRequest {
	s.SessionChangeTime = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionId(v string) *ReportSessionStatusRequest {
	s.SessionId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionStatus(v string) *ReportSessionStatusRequest {
	s.SessionStatus = &v
	return s
}

func (s *ReportSessionStatusRequest) Validate() error {
	return dara.Validate(s)
}
