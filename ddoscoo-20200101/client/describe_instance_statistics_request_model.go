// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeInstanceStatisticsRequest
	GetInstanceIds() []*string
}

type DescribeInstanceStatisticsRequest struct {
	// The ID of the instance that you want to query.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInstanceStatisticsRequest) SetInstanceIds(v []*string) *DescribeInstanceStatisticsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstanceStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
