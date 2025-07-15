// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDesktopGroupTimerResponseBody
	GetRequestId() *string
}

type SetDesktopGroupTimerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 388CF76E-FFB3-5174-9F91-CDD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDesktopGroupTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupTimerResponseBody) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDesktopGroupTimerResponseBody) SetRequestId(v string) *SetDesktopGroupTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDesktopGroupTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
