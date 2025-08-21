// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEpnInstanceResponseBody
	GetRequestId() *string
}

type DeleteEpnInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEpnInstanceResponseBody) SetRequestId(v string) *DeleteEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
