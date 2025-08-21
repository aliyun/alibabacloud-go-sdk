// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRegionBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeNetworkRegionBlockRequest
	GetInstanceId() *string
}

type DescribeNetworkRegionBlockRequest struct {
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeNetworkRegionBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRegionBlockRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkRegionBlockRequest) SetInstanceId(v string) *DescribeNetworkRegionBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkRegionBlockRequest) Validate() error {
	return dara.Validate(s)
}
