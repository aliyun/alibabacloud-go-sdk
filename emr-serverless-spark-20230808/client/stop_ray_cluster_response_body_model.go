// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRayClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRayClusterResponseBody
	GetRequestId() *string
}

type StopRayClusterResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopRayClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRayClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopRayClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRayClusterResponseBody) SetRequestId(v string) *StopRayClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRayClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
