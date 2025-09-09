// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectVpcPeerConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RejectVpcPeerConnectionResponseBody
	GetRequestId() *string
}

type RejectVpcPeerConnectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD2E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectVpcPeerConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *RejectVpcPeerConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectVpcPeerConnectionResponseBody) SetRequestId(v string) *RejectVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectVpcPeerConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
