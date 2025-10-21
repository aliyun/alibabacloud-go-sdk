// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEndpointResponseBody
	GetRequestId() *string
}

type CreateEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEndpointResponseBody) SetRequestId(v string) *CreateEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
