// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptVpcPeerConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcceptVpcPeerConnectionResponseBody
	GetRequestId() *string
}

type AcceptVpcPeerConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptVpcPeerConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptVpcPeerConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptVpcPeerConnectionResponseBody) SetRequestId(v string) *AcceptVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptVpcPeerConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
