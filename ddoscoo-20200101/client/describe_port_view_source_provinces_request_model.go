// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceProvincesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortViewSourceProvincesRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortViewSourceProvincesRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribePortViewSourceProvincesRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortViewSourceProvincesRequest
	GetStartTime() *int64
}

type DescribePortViewSourceProvincesRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds. If you do not configure this parameter, the current system time is used as the end time.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
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
	// The ID of the resource group to which the Anti-DDoS Proxy instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
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

func (s DescribePortViewSourceProvincesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceProvincesRequest) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortViewSourceProvincesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortViewSourceProvincesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortViewSourceProvincesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortViewSourceProvincesRequest) SetEndTime(v int64) *DescribePortViewSourceProvincesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) SetInstanceIds(v []*string) *DescribePortViewSourceProvincesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) SetResourceGroupId(v string) *DescribePortViewSourceProvincesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) SetStartTime(v int64) *DescribePortViewSourceProvincesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) Validate() error {
	return dara.Validate(s)
}
