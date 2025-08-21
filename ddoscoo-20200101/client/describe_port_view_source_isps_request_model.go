// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceIspsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortViewSourceIspsRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortViewSourceIspsRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribePortViewSourceIspsRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortViewSourceIspsRequest
	GetStartTime() *int64
}

type DescribePortViewSourceIspsRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of the Anti-DDoS Proxy instances to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the Anti-DDoS Proxy instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds.
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

func (s DescribePortViewSourceIspsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceIspsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortViewSourceIspsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortViewSourceIspsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortViewSourceIspsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortViewSourceIspsRequest) SetEndTime(v int64) *DescribePortViewSourceIspsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortViewSourceIspsRequest) SetInstanceIds(v []*string) *DescribePortViewSourceIspsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortViewSourceIspsRequest) SetResourceGroupId(v string) *DescribePortViewSourceIspsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortViewSourceIspsRequest) SetStartTime(v int64) *DescribePortViewSourceIspsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortViewSourceIspsRequest) Validate() error {
	return dara.Validate(s)
}
