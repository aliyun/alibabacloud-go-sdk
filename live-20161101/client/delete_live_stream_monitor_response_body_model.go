// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveStreamMonitorResponseBody
	GetRequestId() *string
}

type DeleteLiveStreamMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveStreamMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamMonitorResponseBody) SetRequestId(v string) *DeleteLiveStreamMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
