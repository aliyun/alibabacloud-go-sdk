// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpPeerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBgpPeerResponseBody
	GetRequestId() *string
}

type DeleteBgpPeerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 611CB80C-B6A9-43DB-9E38-0B0AC3D9B58F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBgpPeerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpPeerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBgpPeerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBgpPeerResponseBody) SetRequestId(v string) *DeleteBgpPeerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBgpPeerResponseBody) Validate() error {
	return dara.Validate(s)
}
