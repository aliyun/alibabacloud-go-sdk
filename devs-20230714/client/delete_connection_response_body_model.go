// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConnectionResponseBody
	GetRequestId() *string
}

type DeleteConnectionResponseBody struct {
	// example:
	//
	// A5152937-1C8A-5260-90FA-520CEF028D2D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnectionResponseBody) SetRequestId(v string) *DeleteConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
