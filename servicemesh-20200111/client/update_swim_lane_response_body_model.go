// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSwimLaneResponseBody
	GetRequestId() *string
}

type UpdateSwimLaneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// yyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSwimLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimLaneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSwimLaneResponseBody) SetRequestId(v string) *UpdateSwimLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSwimLaneResponseBody) Validate() error {
	return dara.Validate(s)
}
