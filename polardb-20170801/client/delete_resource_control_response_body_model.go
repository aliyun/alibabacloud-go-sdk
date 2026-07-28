// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceControlResponseBody
	GetRequestId() *string
}

type DeleteResourceControlResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceControlResponseBody) SetRequestId(v string) *DeleteResourceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceControlResponseBody) Validate() error {
	return dara.Validate(s)
}
