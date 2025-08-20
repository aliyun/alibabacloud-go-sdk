// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStackGroupResponseBody
	GetRequestId() *string
}

type DeleteStackGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStackGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStackGroupResponseBody) SetRequestId(v string) *DeleteStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStackGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
