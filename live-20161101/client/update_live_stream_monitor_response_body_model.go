// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveStreamMonitorResponseBody
	GetRequestId() *string
}

type UpdateLiveStreamMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveStreamMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveStreamMonitorResponseBody) SetRequestId(v string) *UpdateLiveStreamMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveStreamMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
