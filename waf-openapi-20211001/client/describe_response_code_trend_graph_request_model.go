// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResponseCodeTrendGraphRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeResponseCodeTrendGraphRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeResponseCodeTrendGraphRequest
	GetInstanceId() *string
	SetInterval(v string) *DescribeResponseCodeTrendGraphRequest
	GetInterval() *string
	SetRegionId(v string) *DescribeResponseCodeTrendGraphRequest
	GetRegionId() *string
	SetResource(v string) *DescribeResponseCodeTrendGraphRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResponseCodeTrendGraphRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeResponseCodeTrendGraphRequest
	GetStartTimestamp() *string
	SetType(v string) *DescribeResponseCodeTrendGraphRequest
	GetType() *string
}

type DescribeResponseCodeTrendGraphRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386280
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval. Unit: seconds. The value must be an integral multiple of 60.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ID of the region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The type of the error codes. Valid values:
	//
	// 	- **waf:*	- error codes that are returned to clients from WAF.
	//
	// 	- **upstream:*	- error codes that are returned to WAF from the origin server.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResponseCodeTrendGraphRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeResponseCodeTrendGraphRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResponseCodeTrendGraphRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeResponseCodeTrendGraphRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResponseCodeTrendGraphRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeResponseCodeTrendGraphRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResponseCodeTrendGraphRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeResponseCodeTrendGraphRequest) GetType() *string {
	return s.Type
}

func (s *DescribeResponseCodeTrendGraphRequest) SetEndTimestamp(v string) *DescribeResponseCodeTrendGraphRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetInstanceId(v string) *DescribeResponseCodeTrendGraphRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetInterval(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Interval = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetRegionId(v string) *DescribeResponseCodeTrendGraphRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetResource(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Resource = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetResourceManagerResourceGroupId(v string) *DescribeResponseCodeTrendGraphRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetStartTimestamp(v string) *DescribeResponseCodeTrendGraphRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetType(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Type = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) Validate() error {
	return dara.Validate(s)
}
