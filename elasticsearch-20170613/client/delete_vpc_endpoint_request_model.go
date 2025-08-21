// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpcEndpointRequest
	GetClientToken() *string
}

type DeleteVpcEndpointRequest struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteVpcEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpcEndpointRequest) SetClientToken(v string) *DeleteVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcEndpointRequest) Validate() error {
	return dara.Validate(s)
}
