// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSwitchNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeInstanceSwitchNetworkRequest
	GetDrdsInstanceId() *string
}

type DescribeInstanceSwitchNetworkRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstanceSwitchNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeInstanceSwitchNetworkRequest) SetDrdsInstanceId(v string) *DescribeInstanceSwitchNetworkRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkRequest) Validate() error {
	return dara.Validate(s)
}
