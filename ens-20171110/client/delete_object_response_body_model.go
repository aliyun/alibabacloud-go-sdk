// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteObjectResponseBody
	GetRequestId() *string
}

type DeleteObjectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A5626B44-0189-443E-9816-D951F596CC89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteObjectResponseBody) SetRequestId(v string) *DeleteObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
