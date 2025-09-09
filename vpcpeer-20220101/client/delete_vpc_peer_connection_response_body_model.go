// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcPeerConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcPeerConnectionResponseBody
	GetRequestId() *string
}

type DeleteVpcPeerConnectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcPeerConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcPeerConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcPeerConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcPeerConnectionResponseBody) SetRequestId(v string) *DeleteVpcPeerConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcPeerConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
