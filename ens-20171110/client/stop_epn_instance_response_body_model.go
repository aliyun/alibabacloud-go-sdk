// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopEpnInstanceResponseBody
	GetRequestId() *string
}

type StopEpnInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopEpnInstanceResponseBody) SetRequestId(v string) *StopEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
