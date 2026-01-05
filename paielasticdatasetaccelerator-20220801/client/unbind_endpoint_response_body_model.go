// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindEndpointResponseBody
	GetRequestId() *string
}

type UnbindEndpointResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindEndpointResponseBody) SetRequestId(v string) *UnbindEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
