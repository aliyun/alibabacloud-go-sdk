// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMemoryNodeResponseBody
	GetRequestId() *string
}

type UpdateMemoryNodeResponseBody struct {
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryNodeResponseBody) SetRequestId(v string) *UpdateMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
