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
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
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
