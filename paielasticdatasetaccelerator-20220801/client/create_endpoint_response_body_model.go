// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *CreateEndpointResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *CreateEndpointResponseBody
	GetRequestId() *string
}

type CreateEndpointResponseBody struct {
	// example:
	//
	// end-5zk866779me51jgu3w
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEndpointResponseBody) SetEndpointId(v string) *CreateEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateEndpointResponseBody) SetRequestId(v string) *CreateEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
