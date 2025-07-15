// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupTimerStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDesktopGroupTimerStatusResponseBody
	GetRequestId() *string
}

type SetDesktopGroupTimerStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AC7912E7-8BDF-547F-BCAC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDesktopGroupTimerStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupTimerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupTimerStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDesktopGroupTimerStatusResponseBody) SetRequestId(v string) *SetDesktopGroupTimerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDesktopGroupTimerStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
