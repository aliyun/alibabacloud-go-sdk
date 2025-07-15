// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcAccessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpcAccessesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcAccessesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpcAccessesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcAccessesResponseBody
	GetTotalCount() *int32
	SetVpcAccessAttributes(v *DescribeVpcAccessesResponseBodyVpcAccessAttributes) *DescribeVpcAccessesResponseBody
	GetVpcAccessAttributes() *DescribeVpcAccessesResponseBodyVpcAccessAttributes
}

type DescribeVpcAccessesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8883AC74-259D-4C0B-99FC-0B7F9A588B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the VPC access authorization. The information is an array consisting of VpcAccessAttribute data.
	VpcAccessAttributes *DescribeVpcAccessesResponseBodyVpcAccessAttributes `json:"VpcAccessAttributes,omitempty" xml:"VpcAccessAttributes,omitempty" type:"Struct"`
}

func (s DescribeVpcAccessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcAccessesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcAccessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcAccessesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcAccessesResponseBody) GetVpcAccessAttributes() *DescribeVpcAccessesResponseBodyVpcAccessAttributes {
	return s.VpcAccessAttributes
}

func (s *DescribeVpcAccessesResponseBody) SetPageNumber(v int32) *DescribeVpcAccessesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetPageSize(v int32) *DescribeVpcAccessesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetRequestId(v string) *DescribeVpcAccessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetTotalCount(v int32) *DescribeVpcAccessesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetVpcAccessAttributes(v *DescribeVpcAccessesResponseBodyVpcAccessAttributes) *DescribeVpcAccessesResponseBody {
	s.VpcAccessAttributes = v
	return s
}

func (s *DescribeVpcAccessesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAccessesResponseBodyVpcAccessAttributes struct {
	VpcAccessAttribute []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute `json:"VpcAccessAttribute,omitempty" xml:"VpcAccessAttribute,omitempty" type:"Repeated"`
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributes) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributes) GetVpcAccessAttribute() []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	return s.VpcAccessAttribute
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributes) SetVpcAccessAttribute(v []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) *DescribeVpcAccessesResponseBodyVpcAccessAttributes {
	s.VpcAccessAttribute = v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute struct {
	// The time when the authorization was created.
	//
	// example:
	//
	// 2017-01-30T04:10:19Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the VPC access authorization.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of an Elastic Compute Service (ECS) or Server Load Balancer (SLB) instance in the VPC.
	//
	// example:
	//
	// i-uf6bzcg1pr4oh5jjmxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the authorization.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port number that corresponds to the instance.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	Tags *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the VPC access authorization.
	//
	// example:
	//
	// vpc-*****ssds24
	VpcAccessId *string `json:"VpcAccessId,omitempty" xml:"VpcAccessId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf657qec7lx42paw3qxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The host of the backend service.
	//
	// example:
	//
	// hos-a***.fh-**nc.com
	VpcTargetHostName *string `json:"VpcTargetHostName,omitempty" xml:"VpcTargetHostName,omitempty"`
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetName() *string {
	return s.Name
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetPort() *int32 {
	return s.Port
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetTags() *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags {
	return s.Tags
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetVpcAccessId() *string {
	return s.VpcAccessId
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GetVpcTargetHostName() *string {
	return s.VpcTargetHostName
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetCreatedTime(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetDescription(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Description = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetInstanceId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetName(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Name = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetPort(v int32) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Port = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetRegionId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetTags(v *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Tags = v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetVpcAccessId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.VpcAccessId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetVpcId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetVpcTargetHostName(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.VpcTargetHostName = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags struct {
	TagInfo []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags) GetTagInfo() []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo {
	return s.TagInfo
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags) SetTagInfo(v []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags {
	s.TagInfo = v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTags) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo struct {
	// The tag key.
	//
	// example:
	//
	// PROJECT
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 6427a17ae6041d1be62414e4
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) SetKey(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) SetValue(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo {
	s.Value = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttributeTagsTagInfo) Validate() error {
	return dara.Validate(s)
}
