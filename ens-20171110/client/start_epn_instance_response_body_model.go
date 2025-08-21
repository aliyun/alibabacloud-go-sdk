// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartEpnInstanceResponseBody
	GetRequestId() *string
}

type StartEpnInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartEpnInstanceResponseBody) SetRequestId(v string) *StartEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
