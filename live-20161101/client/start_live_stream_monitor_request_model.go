// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveStreamMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorId(v string) *StartLiveStreamMonitorRequest
	GetMonitorId() *string
	SetOwnerId(v int64) *StartLiveStreamMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartLiveStreamMonitorRequest
	GetRegionId() *string
}

type StartLiveStreamMonitorRequest struct {
	// The ID of the monitoring session.
	//
	// >  You can obtain the monitoring session ID from the response of the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation.
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

func (s StartLiveStreamMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLiveStreamMonitorRequest) GoString() string {
	return s.String()
}

func (s *StartLiveStreamMonitorRequest) GetMonitorId() *string {
	return s.MonitorId
}

func (s *StartLiveStreamMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartLiveStreamMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartLiveStreamMonitorRequest) SetMonitorId(v string) *StartLiveStreamMonitorRequest {
	s.MonitorId = &v
	return s
}

func (s *StartLiveStreamMonitorRequest) SetOwnerId(v int64) *StartLiveStreamMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *StartLiveStreamMonitorRequest) SetRegionId(v string) *StartLiveStreamMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *StartLiveStreamMonitorRequest) Validate() error {
	return dara.Validate(s)
}
