// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkId(v string) *DescribeNetworkAttributeRequest
	GetNetworkId() *string
}

type DescribeNetworkAttributeRequest struct {
	// The ID of the network.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s DescribeNetworkAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworkAttributeRequest) SetNetworkId(v string) *DescribeNetworkAttributeRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkAttributeRequest) Validate() error {
	return dara.Validate(s)
}
