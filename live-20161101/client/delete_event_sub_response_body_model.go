// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEventSubResponseBody
	GetRequestId() *string
}

type DeleteEventSubResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventSubResponseBody) SetRequestId(v string) *DeleteEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventSubResponseBody) Validate() error {
	return dara.Validate(s)
}
