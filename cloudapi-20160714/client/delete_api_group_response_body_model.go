// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApiGroupResponseBody
	GetRequestId() *string
}

type DeleteApiGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E8515BA6-81CD-4191-A7CF-C4FCDD3C0D99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiGroupResponseBody) SetRequestId(v string) *DeleteApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
