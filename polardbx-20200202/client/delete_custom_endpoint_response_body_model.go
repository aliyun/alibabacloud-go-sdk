// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomEndpointResponseBody
	GetRequestId() *string
}

type DeleteCustomEndpointResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-****-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomEndpointResponseBody) SetRequestId(v string) *DeleteCustomEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
