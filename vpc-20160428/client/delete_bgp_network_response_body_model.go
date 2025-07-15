// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBgpNetworkResponseBody
	GetRequestId() *string
}

type DeleteBgpNetworkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBgpNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBgpNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBgpNetworkResponseBody) SetRequestId(v string) *DeleteBgpNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBgpNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
