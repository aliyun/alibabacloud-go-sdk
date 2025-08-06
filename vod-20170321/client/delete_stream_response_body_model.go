// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStreamResponseBody
	GetRequestId() *string
}

type DeleteStreamResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A7U43F6-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStreamResponseBody) SetRequestId(v string) *DeleteStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
