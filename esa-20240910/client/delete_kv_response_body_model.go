// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKvResponseBody
	GetRequestId() *string
}

type DeleteKvResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKvResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKvResponseBody) SetRequestId(v string) *DeleteKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKvResponseBody) Validate() error {
	return dara.Validate(s)
}
