// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMediaTagResponseBody
	GetRequestId() *string
}

type DeleteMediaTagResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 283DC68C-146F-4489-A2A1-2F88F1472A56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaTagResponseBody) SetRequestId(v string) *DeleteMediaTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaTagResponseBody) Validate() error {
	return dara.Validate(s)
}
