// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteContactResponseBody
	GetRequestId() *string
}

type DeleteContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1F1E4D86-B70B-5352-A641-8CC80D13A37F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactResponseBody) SetRequestId(v string) *DeleteContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactResponseBody) Validate() error {
	return dara.Validate(s)
}
