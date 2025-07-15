// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBgpPeerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpPeerId(v string) *CreateBgpPeerResponseBody
	GetBgpPeerId() *string
	SetRequestId(v string) *CreateBgpPeerResponseBody
	GetRequestId() *string
}

type CreateBgpPeerResponseBody struct {
	// The ID of the BGP peer.
	//
	// example:
	//
	// bgp-m5eoyp2mwegk8ce9v****
	BgpPeerId *string `json:"BgpPeerId,omitempty" xml:"BgpPeerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D4B7649A-61BB-4C64-A586-1DFF1EDA6A42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBgpPeerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBgpPeerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBgpPeerResponseBody) GetBgpPeerId() *string {
	return s.BgpPeerId
}

func (s *CreateBgpPeerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBgpPeerResponseBody) SetBgpPeerId(v string) *CreateBgpPeerResponseBody {
	s.BgpPeerId = &v
	return s
}

func (s *CreateBgpPeerResponseBody) SetRequestId(v string) *CreateBgpPeerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBgpPeerResponseBody) Validate() error {
	return dara.Validate(s)
}
