// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteListResponseBody
	GetRequestId() *string
}

type DeleteListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteListResponseBody) SetRequestId(v string) *DeleteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteListResponseBody) Validate() error {
	return dara.Validate(s)
}
