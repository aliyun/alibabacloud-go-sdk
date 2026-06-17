// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourcePackagesResponseBody
	GetRequestId() *string
	SetResourcePackageList(v []*DescribeResourcePackagesResponseBodyResourcePackageList) *DescribeResourcePackagesResponseBody
	GetResourcePackageList() []*DescribeResourcePackagesResponseBodyResourcePackageList
}

type DescribeResourcePackagesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of cross-cloud resource plans.
	ResourcePackageList []*DescribeResourcePackagesResponseBodyResourcePackageList `json:"ResourcePackageList,omitempty" xml:"ResourcePackageList,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourcePackagesResponseBody) GetResourcePackageList() []*DescribeResourcePackagesResponseBodyResourcePackageList {
	return s.ResourcePackageList
}

func (s *DescribeResourcePackagesResponseBody) SetRequestId(v string) *DescribeResourcePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackagesResponseBody) SetResourcePackageList(v []*DescribeResourcePackagesResponseBodyResourcePackageList) *DescribeResourcePackagesResponseBody {
	s.ResourcePackageList = v
	return s
}

func (s *DescribeResourcePackagesResponseBody) Validate() error {
	if s.ResourcePackageList != nil {
		for _, item := range s.ResourcePackageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcePackagesResponseBodyResourcePackageList struct {
	// Indicates whether automatic quota allocation is enabled.
	//
	// example:
	//
	// true
	AutoQuota *bool `json:"AutoQuota,omitempty" xml:"AutoQuota,omitempty"`
	// The time when the resource plan was created.
	//
	// example:
	//
	// 1744621511000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource plan expires.
	//
	// example:
	//
	// 1747238400000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the cross-cloud resource plan.
	//
	// example:
	//
	// pm-bp11b0i9389******
	ResourcePackageId *string `json:"ResourcePackageId,omitempty" xml:"ResourcePackageId,omitempty"`
	// The quota allocation details.
	ResourcePackageQuotaList []*DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList `json:"ResourcePackageQuotaList,omitempty" xml:"ResourcePackageQuotaList,omitempty" type:"Repeated"`
	// The type of the cross-cloud resource plan.
	//
	// example:
	//
	// MySQL
	ResourcePackageType *string `json:"ResourcePackageType,omitempty" xml:"ResourcePackageType,omitempty"`
	// The status of the cross-cloud resource plan. Valid values:
	//
	// - Normal: Normal.
	//
	// - Maintaining: Under maintenance.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*DescribeResourcePackagesResponseBodyResourcePackageListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total capacity.
	//
	// example:
	//
	// 8
	TotalCapacity *int64 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// The used capacity.
	//
	// example:
	//
	// 4
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s DescribeResourcePackagesResponseBodyResourcePackageList) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackagesResponseBodyResourcePackageList) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetAutoQuota() *bool {
	return s.AutoQuota
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetResourcePackageId() *string {
	return s.ResourcePackageId
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetResourcePackageQuotaList() []*DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList {
	return s.ResourcePackageQuotaList
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetResourcePackageType() *string {
	return s.ResourcePackageType
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetTags() []*DescribeResourcePackagesResponseBodyResourcePackageListTags {
	return s.Tags
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetTotalCapacity() *int64 {
	return s.TotalCapacity
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetAutoQuota(v bool) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.AutoQuota = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetCreateTime(v int64) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetExpireTime(v int64) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetResourcePackageId(v string) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.ResourcePackageId = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetResourcePackageQuotaList(v []*DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.ResourcePackageQuotaList = v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetResourcePackageType(v string) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.ResourcePackageType = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetStatus(v string) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.Status = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetTags(v []*DescribeResourcePackagesResponseBodyResourcePackageListTags) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.Tags = v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetTotalCapacity(v int64) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) SetUsedCapacity(v int64) *DescribeResourcePackagesResponseBodyResourcePackageList {
	s.UsedCapacity = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageList) Validate() error {
	if s.ResourcePackageQuotaList != nil {
		for _, item := range s.ResourcePackageQuotaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList struct {
	// The capacity allocated to the resource pool.
	//
	// example:
	//
	// 6
	AllocatedCapacity *int64 `json:"AllocatedCapacity,omitempty" xml:"AllocatedCapacity,omitempty"`
	// The ID of the resource pool.
	//
	// example:
	//
	// pj-87681rbcef6******
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The used capacity of the resource pool.
	//
	// example:
	//
	// 2
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) GetAllocatedCapacity() *int64 {
	return s.AllocatedCapacity
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) SetAllocatedCapacity(v int64) *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList {
	s.AllocatedCapacity = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) SetProjectId(v string) *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList {
	s.ProjectId = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) SetUsedCapacity(v int64) *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList {
	s.UsedCapacity = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListResourcePackageQuotaList) Validate() error {
	return dara.Validate(s)
}

type DescribeResourcePackagesResponseBodyResourcePackageListTags struct {
	// The tag key.
	//
	// example:
	//
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourcePackagesResponseBodyResourcePackageListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackagesResponseBodyResourcePackageListTags) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListTags) GetKey() *string {
	return s.Key
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListTags) GetValue() *string {
	return s.Value
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListTags) SetKey(v string) *DescribeResourcePackagesResponseBodyResourcePackageListTags {
	s.Key = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListTags) SetValue(v string) *DescribeResourcePackagesResponseBodyResourcePackageListTags {
	s.Value = &v
	return s
}

func (s *DescribeResourcePackagesResponseBodyResourcePackageListTags) Validate() error {
	return dara.Validate(s)
}
