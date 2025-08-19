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
	SetResult(v bool) *DeleteEndpointResponseBody
	GetResult() *bool
}

type DeleteEndpointResponseBody struct {
	// example:
	//
	// 1305A3E0-A291-54BA-A3B2-7D1C12EC4112
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
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

func (s *DeleteEndpointResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteEndpointResponseBody) SetRequestId(v string) *DeleteEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEndpointResponseBody) SetResult(v bool) *DeleteEndpointResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
