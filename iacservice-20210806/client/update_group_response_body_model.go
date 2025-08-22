// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGroupResponseBody
	GetRequestId() *string
}

type UpdateGroupResponseBody struct {
	// example:
	//
	// 4EF5E823-AC0D-5CB5-8F18-1352455A488D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
