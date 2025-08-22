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
	// example:
	//
	// 1E7BA3EB-B0EF-53F5-9999-07CAD6D9F8A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
