// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePdpConfigResponseBody
	GetRequestId() *string
}

type DeletePdpConfigResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePdpConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePdpConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePdpConfigResponseBody) SetRequestId(v string) *DeletePdpConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePdpConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
