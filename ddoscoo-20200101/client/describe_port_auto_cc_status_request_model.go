// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortAutoCcStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribePortAutoCcStatusRequest
	GetInstanceIds() []*string
}

type DescribePortAutoCcStatusRequest struct {
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

func (s DescribePortAutoCcStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortAutoCcStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortAutoCcStatusRequest) SetInstanceIds(v []*string) *DescribePortAutoCcStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortAutoCcStatusRequest) Validate() error {
	return dara.Validate(s)
}
