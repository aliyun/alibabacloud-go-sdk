// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMemoryResponseBody
	GetRequestId() *string
}

type UpdateMemoryResponseBody struct {
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryResponseBody) SetRequestId(v string) *UpdateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
