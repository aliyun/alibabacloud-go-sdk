// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortConnsCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortConnsCountRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortConnsCountRequest
	GetInstanceIds() []*string
	SetPort(v string) *DescribePortConnsCountRequest
	GetPort() *string
	SetResourceGroupId(v string) *DescribePortConnsCountRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortConnsCountRequest
	GetStartTime() *int64
}

type DescribePortConnsCountRequest struct {
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
	// An array that consists of the IDs of instances.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s DescribePortConnsCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortConnsCountRequest) GoString() string {
	return s.String()
}

func (s *DescribePortConnsCountRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortConnsCountRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortConnsCountRequest) GetPort() *string {
	return s.Port
}

func (s *DescribePortConnsCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortConnsCountRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortConnsCountRequest) SetEndTime(v int64) *DescribePortConnsCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetInstanceIds(v []*string) *DescribePortConnsCountRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortConnsCountRequest) SetPort(v string) *DescribePortConnsCountRequest {
	s.Port = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetResourceGroupId(v string) *DescribePortConnsCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetStartTime(v int64) *DescribePortConnsCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortConnsCountRequest) Validate() error {
	return dara.Validate(s)
}
