// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkResponseBody
	GetRequestId() *string
}

type DeleteNetworkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkResponseBody) SetRequestId(v string) *DeleteNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
