// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNavigationConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeNavigationConfigRequest
	GetInstanceId() *string
}

type DescribeNavigationConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 287289b6-1510-4e64-9224-39b53ad89dd2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeNavigationConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNavigationConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNavigationConfigRequest) SetInstanceId(v string) *DescribeNavigationConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNavigationConfigRequest) Validate() error {
	return dara.Validate(s)
}
