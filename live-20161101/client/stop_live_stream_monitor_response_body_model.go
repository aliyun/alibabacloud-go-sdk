// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveStreamMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopLiveStreamMonitorResponseBody
	GetRequestId() *string
}

type StopLiveStreamMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLiveStreamMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLiveStreamMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveStreamMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLiveStreamMonitorResponseBody) SetRequestId(v string) *StopLiveStreamMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLiveStreamMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
