// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSwimLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSwimLaneGroupResponseBody
	GetRequestId() *string
}

type CreateSwimLaneGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// yyyy
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSwimLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSwimLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSwimLaneGroupResponseBody) SetRequestId(v string) *CreateSwimLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSwimLaneGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
