// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceResponseBody
	GetCode() *string
	SetInstances(v []*ListInstanceResponseBodyInstances) *ListInstanceResponseBody
	GetInstances() []*ListInstanceResponseBodyInstances
	SetIsSuccess(v bool) *ListInstanceResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListInstanceResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListInstanceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListInstanceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInstanceResponseBody
	GetTotalCount() *int32
}

type ListInstanceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried instances.
	Instances []*ListInstanceResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A2A9BA68-B264-4953-9154-CE61B1C03BA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12121
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceResponseBody) GetInstances() []*ListInstanceResponseBodyInstances {
	return s.Instances
}

func (s *ListInstanceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListInstanceResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListInstanceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstanceResponseBody) SetCode(v string) *ListInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceResponseBody) SetInstances(v []*ListInstanceResponseBodyInstances) *ListInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstanceResponseBody) SetIsSuccess(v bool) *ListInstanceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListInstanceResponseBody) SetPageNo(v int32) *ListInstanceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListInstanceResponseBody) SetPageSize(v int32) *ListInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetTotalCount(v int32) *ListInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceResponseBody) Validate() error {
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

type ListInstanceResponseBodyInstances struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 1562849679000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-sgedpenzw80e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issue occurs on the instance.
	//
	// example:
	//
	// oss bucket already exists
	InstanceIssue *string `json:"InstanceIssue,omitempty" xml:"InstanceIssue,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The edition of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// Enterprise_Basic
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 1562849760000
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek2h3aexpy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the instance.
	Tags []*ListInstanceResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstanceResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceResponseBodyInstances) GetInstanceIssue() *string {
	return s.InstanceIssue
}

func (s *ListInstanceResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceResponseBodyInstances) GetInstanceSpecification() *string {
	return s.InstanceSpecification
}

func (s *ListInstanceResponseBodyInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListInstanceResponseBodyInstances) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListInstanceResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstanceResponseBodyInstances) GetTags() []*ListInstanceResponseBodyInstancesTags {
	return s.Tags
}

func (s *ListInstanceResponseBodyInstances) SetCreateTime(v string) *ListInstanceResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceId(v string) *ListInstanceResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceIssue(v string) *ListInstanceResponseBodyInstances {
	s.InstanceIssue = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceName(v string) *ListInstanceResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceSpecification(v string) *ListInstanceResponseBodyInstances {
	s.InstanceSpecification = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceStatus(v string) *ListInstanceResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetModifiedTime(v string) *ListInstanceResponseBodyInstances {
	s.ModifiedTime = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetRegionId(v string) *ListInstanceResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetResourceGroupId(v string) *ListInstanceResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetTags(v []*ListInstanceResponseBodyInstancesTags) *ListInstanceResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *ListInstanceResponseBodyInstances) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResponseBodyInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListInstanceResponseBodyInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstancesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListInstanceResponseBodyInstancesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListInstanceResponseBodyInstancesTags) SetTagKey(v string) *ListInstanceResponseBodyInstancesTags {
	s.TagKey = &v
	return s
}

func (s *ListInstanceResponseBodyInstancesTags) SetTagValue(v string) *ListInstanceResponseBodyInstancesTags {
	s.TagValue = &v
	return s
}

func (s *ListInstanceResponseBodyInstancesTags) Validate() error {
	return dara.Validate(s)
}
