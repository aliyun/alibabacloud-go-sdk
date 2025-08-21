// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcEndpointResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteVpcEndpointResponseBody
	GetResult() *bool
}

type DeleteVpcEndpointResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcEndpointResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteVpcEndpointResponseBody) SetRequestId(v string) *DeleteVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcEndpointResponseBody) SetResult(v bool) *DeleteVpcEndpointResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteVpcEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
