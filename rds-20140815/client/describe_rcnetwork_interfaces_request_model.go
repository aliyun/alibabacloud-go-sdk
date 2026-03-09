// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCNetworkInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRCNetworkInterfacesRequest
	GetInstanceId() *string
}

type DescribeRCNetworkInterfacesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rc-m5sc1271fv344a1r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeRCNetworkInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCNetworkInterfacesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCNetworkInterfacesRequest) SetInstanceId(v string) *DescribeRCNetworkInterfacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCNetworkInterfacesRequest) Validate() error {
	return dara.Validate(s)
}
