// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEndpointResponseBody
	GetRequestId() *string
}

type DeleteEndpointResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEndpointResponseBody) SetRequestId(v string) *DeleteEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
