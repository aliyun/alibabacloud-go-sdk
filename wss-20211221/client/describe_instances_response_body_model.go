// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() []*DescribeInstancesResponseBodyInstances
	SetMaxResults(v int32) *DescribeInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
}

type DescribeInstancesResponseBody struct {
	// The list of sold instances that meet the query conditions.
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The maximum number of records to return per query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Leave it empty for the first query. For subsequent queries, pass the NextToken value returned in the previous response as is. Do not construct or decode it yourself.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6k=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot call issues.
	//
	// example:
	//
	// 72481C12-69AB-5CE1-8A35-A8EFA921****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() []*DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetMaxResults(v int32) *DescribeInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetNextToken(v string) *DescribeInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	// The Alibaba Cloud account ID to which the resource belongs.
	//
	// example:
	//
	// 123456789012****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The expiration time in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-10-01T00:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The sold instance ID, which is used to identify the purchased instance.
	//
	// example:
	//
	// cp-0bw2ew49****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type identifier.
	//
	// example:
	//
	// ecd.graphics.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The list of instance properties. Each item contains the property name Key and the property value Value.
	PropertyList []*DescribeInstancesResponseBodyInstancesPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The associated business resource ID, which is used to distinguish from the sold instance ID.
	//
	// example:
	//
	// ecd-7w78ozhj****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. For example, CorePackage indicates a core-hour package. It must be consistent with the resource type of the target instance.
	//
	// example:
	//
	// CorePackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The effective time in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The instance status. Valid values:
	//
	// - NORMAL: Normal.
	//
	// - DELETED: Deleted.
	//
	// - PENDING: Pending activation.
	//
	// - EXPIRED: Expired.
	//
	// - ARCHIVED: Historical archived status.
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The total metering quota included in the instance, using the same unit as UsedTime. A CorePackage core-hour package is metered in core-seconds. For example, 360000 indicates 100 core-hours.
	//
	// example:
	//
	// 360000
	TotalTime *int32 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The consumed metering quota of the instance. For a CorePackage core-hour package, the unit is core-seconds. Divide by 3600 to convert to core-hours. For other resource types, read according to the corresponding metering criteria.
	//
	// example:
	//
	// 0
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeInstancesResponseBodyInstances) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesResponseBodyInstances) GetPropertyList() []*DescribeInstancesResponseBodyInstancesPropertyList {
	return s.PropertyList
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeInstancesResponseBodyInstances) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstancesResponseBodyInstances) GetState() *string {
	return s.State
}

func (s *DescribeInstancesResponseBodyInstances) GetTotalTime() *int32 {
	return s.TotalTime
}

func (s *DescribeInstancesResponseBodyInstances) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *DescribeInstancesResponseBodyInstances) SetAliUid(v int64) *DescribeInstancesResponseBodyInstances {
	s.AliUid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpiredTime(v string) *DescribeInstancesResponseBodyInstances {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceType(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPropertyList(v []*DescribeInstancesResponseBodyInstancesPropertyList) *DescribeInstancesResponseBodyInstances {
	s.PropertyList = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceType(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v string) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetState(v string) *DescribeInstancesResponseBodyInstances {
	s.State = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetTotalTime(v int32) *DescribeInstancesResponseBodyInstances {
	s.TotalTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUsedTime(v int32) *DescribeInstancesResponseBodyInstances {
	s.UsedTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	if s.PropertyList != nil {
		for _, item := range s.PropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesPropertyList struct {
	// The instance property name.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The instance property value, which corresponds to the Key in the same entry.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesPropertyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesPropertyList) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesPropertyList) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesResponseBodyInstancesPropertyList) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyInstancesPropertyList) SetKey(v string) *DescribeInstancesResponseBodyInstancesPropertyList {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesPropertyList) SetValue(v string) *DescribeInstancesResponseBodyInstancesPropertyList {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesPropertyList) Validate() error {
	return dara.Validate(s)
}
