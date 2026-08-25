// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGroupResponseBody
	GetRequestId() *string
}

type DeleteGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F723DE01-6276-5DC4-9B1F-9CBE3E1748B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
