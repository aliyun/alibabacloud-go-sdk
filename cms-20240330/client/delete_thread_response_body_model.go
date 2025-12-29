// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteThreadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteThreadResponseBody
	GetRequestId() *string
}

type DeleteThreadResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 123-0F43-23423-AC43-34234
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteThreadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteThreadResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteThreadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteThreadResponseBody) SetRequestId(v string) *DeleteThreadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteThreadResponseBody) Validate() error {
	return dara.Validate(s)
}
