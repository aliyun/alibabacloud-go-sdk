// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTokenCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendTokenCodeResponseBody
	GetRequestId() *string
}

type SendTokenCodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 134BD0B2-B848-5743-9CE2-C1FD3D5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTokenCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendTokenCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendTokenCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendTokenCodeResponseBody) SetRequestId(v string) *SendTokenCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendTokenCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
