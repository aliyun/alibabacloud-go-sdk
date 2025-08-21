// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceCountriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePortViewSourceCountriesRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePortViewSourceCountriesRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribePortViewSourceCountriesRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribePortViewSourceCountriesRequest
	GetStartTime() *int64
}

type DescribePortViewSourceCountriesRequest struct {
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
	// The IDs of the Anti-DDoS Proxy instances to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the Anti-DDoS Proxy instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
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

func (s DescribePortViewSourceCountriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceCountriesRequest) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePortViewSourceCountriesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePortViewSourceCountriesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortViewSourceCountriesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePortViewSourceCountriesRequest) SetEndTime(v int64) *DescribePortViewSourceCountriesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) SetInstanceIds(v []*string) *DescribePortViewSourceCountriesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) SetResourceGroupId(v string) *DescribePortViewSourceCountriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) SetStartTime(v int64) *DescribePortViewSourceCountriesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) Validate() error {
	return dara.Validate(s)
}
