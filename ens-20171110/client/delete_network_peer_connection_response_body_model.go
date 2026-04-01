// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPeerConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkPeerConnectionResponseBody
	GetRequestId() *string
}

type DeleteNetworkPeerConnectionResponseBody struct {
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkPeerConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPeerConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkPeerConnectionResponseBody) SetRequestId(v string) *DeleteNetworkPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkPeerConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
