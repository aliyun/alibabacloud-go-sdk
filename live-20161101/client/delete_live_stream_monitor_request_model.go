// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorId(v string) *DeleteLiveStreamMonitorRequest
	GetMonitorId() *string
	SetOwnerId(v int64) *DeleteLiveStreamMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveStreamMonitorRequest
	GetRegionId() *string
}

type DeleteLiveStreamMonitorRequest struct {
	// The ID of the monitoring session.
	//
	// >  You can obtain the monitoring session ID from the response parameter **MonitorId*	- of the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation.
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

func (s DeleteLiveStreamMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamMonitorRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamMonitorRequest) GetMonitorId() *string {
	return s.MonitorId
}

func (s *DeleteLiveStreamMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveStreamMonitorRequest) SetMonitorId(v string) *DeleteLiveStreamMonitorRequest {
	s.MonitorId = &v
	return s
}

func (s *DeleteLiveStreamMonitorRequest) SetOwnerId(v int64) *DeleteLiveStreamMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamMonitorRequest) SetRegionId(v string) *DeleteLiveStreamMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveStreamMonitorRequest) Validate() error {
	return dara.Validate(s)
}
