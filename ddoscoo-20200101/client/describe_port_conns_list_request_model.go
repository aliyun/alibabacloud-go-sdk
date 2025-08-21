// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortConnsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortConnsListRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortConnsListRequest
	GetInstanceIds() []*string
	SetInterval(v int32) *DescribePortConnsListRequest
	GetInterval() *int32
	SetPort(v string) *DescribePortConnsListRequest
	GetPort() *string
	SetResourceGroupId(v string) *DescribePortConnsListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortConnsListRequest
	GetStartTime() *int64
}

type DescribePortConnsListRequest struct {
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
	// The interval for returning data. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of port that you want to query. If you do not specify this parameter, all ports are queried.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
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

func (s DescribePortConnsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortConnsListRequest) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortConnsListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortConnsListRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribePortConnsListRequest) GetPort() *string {
	return s.Port
}

func (s *DescribePortConnsListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortConnsListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortConnsListRequest) SetEndTime(v int64) *DescribePortConnsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortConnsListRequest) SetInstanceIds(v []*string) *DescribePortConnsListRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortConnsListRequest) SetInterval(v int32) *DescribePortConnsListRequest {
	s.Interval = &v
	return s
}

func (s *DescribePortConnsListRequest) SetPort(v string) *DescribePortConnsListRequest {
	s.Port = &v
	return s
}

func (s *DescribePortConnsListRequest) SetResourceGroupId(v string) *DescribePortConnsListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortConnsListRequest) SetStartTime(v int64) *DescribePortConnsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortConnsListRequest) Validate() error {
	return dara.Validate(s)
}
