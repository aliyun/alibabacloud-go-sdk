// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRayHistoryServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRayHistoryServerResponseBody
	GetRequestId() *string
}

type StopRayHistoryServerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRayHistoryServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRayHistoryServerResponseBody) GoString() string {
	return s.String()
}

func (s *StopRayHistoryServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRayHistoryServerResponseBody) SetRequestId(v string) *StopRayHistoryServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRayHistoryServerResponseBody) Validate() error {
	return dara.Validate(s)
}
