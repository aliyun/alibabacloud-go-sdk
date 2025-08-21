// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortMaxConnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortMaxConnsRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortMaxConnsRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribePortMaxConnsRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortMaxConnsRequest
	GetStartTime() *int64
}

type DescribePortMaxConnsRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of the Anti-DDoS Proxy instances.
	//
	// >  You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all Anti-DDoS Proxy instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePortMaxConnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortMaxConnsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortMaxConnsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortMaxConnsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortMaxConnsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortMaxConnsRequest) SetEndTime(v int64) *DescribePortMaxConnsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortMaxConnsRequest) SetInstanceIds(v []*string) *DescribePortMaxConnsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortMaxConnsRequest) SetResourceGroupId(v string) *DescribePortMaxConnsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortMaxConnsRequest) SetStartTime(v int64) *DescribePortMaxConnsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortMaxConnsRequest) Validate() error {
	return dara.Validate(s)
}
