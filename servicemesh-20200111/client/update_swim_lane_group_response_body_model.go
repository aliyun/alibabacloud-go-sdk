// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSwimLaneGroupResponseBody
	GetRequestId() *string
}

type UpdateSwimLaneGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// yyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSwimLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSwimLaneGroupResponseBody) SetRequestId(v string) *UpdateSwimLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSwimLaneGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
