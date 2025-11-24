// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSwimLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSwimLaneResponseBody
	GetRequestId() *string
}

type CreateSwimLaneResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// *****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSwimLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSwimLaneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSwimLaneResponseBody) SetRequestId(v string) *CreateSwimLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSwimLaneResponseBody) Validate() error {
	return dara.Validate(s)
}
