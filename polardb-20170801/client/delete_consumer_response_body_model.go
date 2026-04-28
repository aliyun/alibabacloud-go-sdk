// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConsumerResponseBody
	GetRequestId() *string
}

type DeleteConsumerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerResponseBody) SetRequestId(v string) *DeleteConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
