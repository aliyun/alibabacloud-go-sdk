// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupScaleTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDesktopGroupScaleTimerResponseBody
	GetRequestId() *string
}

type SetDesktopGroupScaleTimerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDesktopGroupScaleTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupScaleTimerResponseBody) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupScaleTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDesktopGroupScaleTimerResponseBody) SetRequestId(v string) *SetDesktopGroupScaleTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDesktopGroupScaleTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
