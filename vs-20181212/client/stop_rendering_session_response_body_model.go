// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRenderingSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRenderingSessionResponseBody
	GetRequestId() *string
}

type StopRenderingSessionResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRenderingSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRenderingSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StopRenderingSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRenderingSessionResponseBody) SetRequestId(v string) *StopRenderingSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRenderingSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
