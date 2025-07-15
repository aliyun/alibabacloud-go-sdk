// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeInstanceFromCenResponseBody
	GetRequestId() *string
}

type RevokeInstanceFromCenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeInstanceFromCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromCenResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeInstanceFromCenResponseBody) SetRequestId(v string) *RevokeInstanceFromCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeInstanceFromCenResponseBody) Validate() error {
	return dara.Validate(s)
}
