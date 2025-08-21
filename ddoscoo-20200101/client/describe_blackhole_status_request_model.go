// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlackholeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeBlackholeStatusRequest
	GetInstanceIds() []*string
}

type DescribeBlackholeStatusRequest struct {
	// The ID of the instance.
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

func (s DescribeBlackholeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlackholeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeBlackholeStatusRequest) SetInstanceIds(v []*string) *DescribeBlackholeStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeBlackholeStatusRequest) Validate() error {
	return dara.Validate(s)
}
