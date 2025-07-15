// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveStreamMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorId(v string) *StopLiveStreamMonitorRequest
	GetMonitorId() *string
	SetOwnerId(v int64) *StopLiveStreamMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopLiveStreamMonitorRequest
	GetRegionId() *string
}

type StopLiveStreamMonitorRequest struct {
	// The ID of the monitoring session.
	//
	// >  You can obtain the monitoring session ID***	- from the response of the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	MonitorId *string `json:"MonitorId,omitempty" xml:"MonitorId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopLiveStreamMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLiveStreamMonitorRequest) GoString() string {
	return s.String()
}

func (s *StopLiveStreamMonitorRequest) GetMonitorId() *string {
	return s.MonitorId
}

func (s *StopLiveStreamMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopLiveStreamMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopLiveStreamMonitorRequest) SetMonitorId(v string) *StopLiveStreamMonitorRequest {
	s.MonitorId = &v
	return s
}

func (s *StopLiveStreamMonitorRequest) SetOwnerId(v int64) *StopLiveStreamMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *StopLiveStreamMonitorRequest) SetRegionId(v string) *StopLiveStreamMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *StopLiveStreamMonitorRequest) Validate() error {
	return dara.Validate(s)
}
