// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeGroupedInstancesResponseBodyInstances) *DescribeGroupedInstancesResponseBody
	GetInstances() []*DescribeGroupedInstancesResponseBodyInstances
	SetPageInfo(v *DescribeGroupedInstancesResponseBodyPageInfo) *DescribeGroupedInstancesResponseBody
	GetPageInfo() *DescribeGroupedInstancesResponseBodyPageInfo
	SetRequestId(v string) *DescribeGroupedInstancesResponseBody
	GetRequestId() *string
}

type DescribeGroupedInstancesResponseBody struct {
	// The information about the assets.
	Instances []*DescribeGroupedInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeGroupedInstancesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 52A3AEE6-114A-499D-8990-4BA9B27FE0AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupedInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedInstancesResponseBody) GetInstances() []*DescribeGroupedInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeGroupedInstancesResponseBody) GetPageInfo() *DescribeGroupedInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeGroupedInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupedInstancesResponseBody) SetInstances(v []*DescribeGroupedInstancesResponseBodyInstances) *DescribeGroupedInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeGroupedInstancesResponseBody) SetPageInfo(v *DescribeGroupedInstancesResponseBodyPageInfo) *DescribeGroupedInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeGroupedInstancesResponseBody) SetRequestId(v string) *DescribeGroupedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupedInstancesResponseBodyInstances struct {
	// The number of assets on which high-risk vulnerabilities are detected.
	//
	// example:
	//
	// 11
	AsapVulInstanceCount *int64 `json:"AsapVulInstanceCount,omitempty" xml:"AsapVulInstanceCount,omitempty"`
	// The number of assets that are protected by the specified edition.
	//
	// example:
	//
	// 205
	AuthVersionCheckCount *int32 `json:"AuthVersionCheckCount,omitempty" xml:"AuthVersionCheckCount,omitempty"`
	// The name of the server group.
	//
	// example:
	//
	// testA
	FieldAliasName *string `json:"FieldAliasName,omitempty" xml:"FieldAliasName,omitempty"`
	// The type of the server group. Valid values:
	//
	// 	- **0**: the default group
	//
	// 	- **1**: other group
	//
	// example:
	//
	// 1
	GroupFlag *int32 `json:"GroupFlag,omitempty" xml:"GroupFlag,omitempty"`
	// The number of cores of assets in the specified asset type.
	//
	// >  If the **MachineTypes*	- request parameter is not specified, the value of the InstanceCoreCount parameter indicates the total number of cores of assets within your account.
	//
	// example:
	//
	// 610
	InstanceCoreCount *int64 `json:"InstanceCoreCount,omitempty" xml:"InstanceCoreCount,omitempty"`
	// The total number of assets that belong to the specified type.
	//
	// >  If the **MachineTypes*	- request parameter is not specified, the value of the InstanceCount parameter is the total number of your assets.
	//
	// example:
	//
	// 205
	InstanceCount *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The operating system type of the asset. Valid values:
	//
	// 	- **windows**
	//
	// 	- **linux**
	//
	// > This parameter is returned only when Lang is set to zh.
	//
	// example:
	//
	// windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The number of assets that are at risk.
	//
	// example:
	//
	// 172
	RiskInstanceCount *string `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
	// The number of assets that are not protected by Security Center.
	//
	// example:
	//
	// 32
	UnProtectedInstanceCount *string `json:"UnProtectedInstanceCount,omitempty" xml:"UnProtectedInstanceCount,omitempty"`
}

func (s DescribeGroupedInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetAsapVulInstanceCount() *int64 {
	return s.AsapVulInstanceCount
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetAuthVersionCheckCount() *int32 {
	return s.AuthVersionCheckCount
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetFieldAliasName() *string {
	return s.FieldAliasName
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetGroupFlag() *int32 {
	return s.GroupFlag
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetInstanceCoreCount() *int64 {
	return s.InstanceCoreCount
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetInstanceCount() *string {
	return s.InstanceCount
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetOs() *string {
	return s.Os
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetRiskInstanceCount() *string {
	return s.RiskInstanceCount
}

func (s *DescribeGroupedInstancesResponseBodyInstances) GetUnProtectedInstanceCount() *string {
	return s.UnProtectedInstanceCount
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetAsapVulInstanceCount(v int64) *DescribeGroupedInstancesResponseBodyInstances {
	s.AsapVulInstanceCount = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetAuthVersionCheckCount(v int32) *DescribeGroupedInstancesResponseBodyInstances {
	s.AuthVersionCheckCount = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetFieldAliasName(v string) *DescribeGroupedInstancesResponseBodyInstances {
	s.FieldAliasName = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetGroupFlag(v int32) *DescribeGroupedInstancesResponseBodyInstances {
	s.GroupFlag = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetInstanceCoreCount(v int64) *DescribeGroupedInstancesResponseBodyInstances {
	s.InstanceCoreCount = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetInstanceCount(v string) *DescribeGroupedInstancesResponseBodyInstances {
	s.InstanceCount = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetOs(v string) *DescribeGroupedInstancesResponseBodyInstances {
	s.Os = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetRiskInstanceCount(v string) *DescribeGroupedInstancesResponseBodyInstances {
	s.RiskInstanceCount = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) SetUnProtectedInstanceCount(v string) *DescribeGroupedInstancesResponseBodyInstances {
	s.UnProtectedInstanceCount = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupedInstancesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 5
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupedInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeGroupedInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeGroupedInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeGroupedInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeGroupedInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
