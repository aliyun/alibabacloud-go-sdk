// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveStreamMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartLiveStreamMonitorResponseBody
	GetRequestId() *string
}

type StartLiveStreamMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartLiveStreamMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartLiveStreamMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *StartLiveStreamMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartLiveStreamMonitorResponseBody) SetRequestId(v string) *StartLiveStreamMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLiveStreamMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
