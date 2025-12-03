// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEaisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeEaisResponseBodyInstances) *DescribeEaisResponseBody
	GetInstances() *DescribeEaisResponseBodyInstances
	SetPageNumber(v int32) *DescribeEaisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEaisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEaisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEaisResponseBody
	GetTotalCount() *int32
}

type DescribeEaisResponseBody struct {
	Instances *DescribeEaisResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1E23D585-BBD8-436F-9615-54CACD6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEaisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBody) GetInstances() *DescribeEaisResponseBodyInstances {
	return s.Instances
}

func (s *DescribeEaisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEaisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEaisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEaisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEaisResponseBody) SetInstances(v *DescribeEaisResponseBodyInstances) *DescribeEaisResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeEaisResponseBody) SetPageNumber(v int32) *DescribeEaisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEaisResponseBody) SetPageSize(v int32) *DescribeEaisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEaisResponseBody) SetRequestId(v string) *DescribeEaisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEaisResponseBody) SetTotalCount(v int32) *DescribeEaisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEaisResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEaisResponseBodyInstances struct {
	Instance []*DescribeEaisResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeEaisResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstances) GetInstance() []*DescribeEaisResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeEaisResponseBodyInstances) SetInstance(v []*DescribeEaisResponseBodyInstancesInstance) *DescribeEaisResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeEaisResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEaisResponseBodyInstancesInstance struct {
	// example:
	//
	// jupyter
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// i-wz93g6pyat2g****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// test1
	ClientInstanceName *string `json:"ClientInstanceName,omitempty" xml:"ClientInstanceName,omitempty"`
	// example:
	//
	// ecs.g5ne.large
	ClientInstanceType *string `json:"ClientInstanceType,omitempty" xml:"ClientInstanceType,omitempty"`
	// example:
	//
	// 2020-11-11T03:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// eais-sz8t15a7gt7****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// testName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// eais.ei-a6.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// http://121.41.**.24:8888
	JupyterUrl  *string `json:"JupyterUrl,omitempty" xml:"JupyterUrl,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// sg-bp1gppir818lx4******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// InUse
	Status *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeEaisResponseBodyInstancesInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// vsw-bp1sd131hfmd76r******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeEaisResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetCategory() *string {
	return s.Category
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetClientInstanceName() *string {
	return s.ClientInstanceName
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetClientInstanceType() *string {
	return s.ClientInstanceType
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetJupyterUrl() *string {
	return s.JupyterUrl
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetTags() *DescribeEaisResponseBodyInstancesInstanceTags {
	return s.Tags
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeEaisResponseBodyInstancesInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetCategory(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.Category = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetClientInstanceId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ClientInstanceId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetClientInstanceName(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ClientInstanceName = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetClientInstanceType(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ClientInstanceType = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetCreationTime(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetDescription(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetElasticAcceleratedInstanceId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetInstanceType(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetJupyterUrl(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.JupyterUrl = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetPaymentType(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.PaymentType = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetRegionId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetSecurityGroupId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetStartTime(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetStatus(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetTags(v *DescribeEaisResponseBodyInstancesInstanceTags) *DescribeEaisResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetVSwitchId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) SetZoneId(v string) *DescribeEaisResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstance) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEaisResponseBodyInstancesInstanceTags struct {
	Tag []*DescribeEaisResponseBodyInstancesInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEaisResponseBodyInstancesInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstancesInstanceTags) GetTag() []*DescribeEaisResponseBodyInstancesInstanceTagsTag {
	return s.Tag
}

func (s *DescribeEaisResponseBodyInstancesInstanceTags) SetTag(v []*DescribeEaisResponseBodyInstancesInstanceTagsTag) *DescribeEaisResponseBodyInstancesInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstanceTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEaisResponseBodyInstancesInstanceTagsTag struct {
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeEaisResponseBodyInstancesInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisResponseBodyInstancesInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponseBodyInstancesInstanceTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeEaisResponseBodyInstancesInstanceTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeEaisResponseBodyInstancesInstanceTagsTag) SetTagKey(v string) *DescribeEaisResponseBodyInstancesInstanceTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstanceTagsTag) SetTagValue(v string) *DescribeEaisResponseBodyInstancesInstanceTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeEaisResponseBodyInstancesInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
