// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveStreamMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorId(v string) *CreateLiveStreamMonitorResponseBody
	GetMonitorId() *string
	SetRequestId(v string) *CreateLiveStreamMonitorResponseBody
	GetRequestId() *string
}

type CreateLiveStreamMonitorResponseBody struct {
	// The ID of the monitoring session.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	MonitorId *string `json:"MonitorId,omitempty" xml:"MonitorId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLiveStreamMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveStreamMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamMonitorResponseBody) GetMonitorId() *string {
	return s.MonitorId
}

func (s *CreateLiveStreamMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveStreamMonitorResponseBody) SetMonitorId(v string) *CreateLiveStreamMonitorResponseBody {
	s.MonitorId = &v
	return s
}

func (s *CreateLiveStreamMonitorResponseBody) SetRequestId(v string) *CreateLiveStreamMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveStreamMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
