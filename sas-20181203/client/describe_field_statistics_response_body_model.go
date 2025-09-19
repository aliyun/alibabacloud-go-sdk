// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupedFields(v *DescribeFieldStatisticsResponseBodyGroupedFields) *DescribeFieldStatisticsResponseBody
	GetGroupedFields() *DescribeFieldStatisticsResponseBodyGroupedFields
	SetRequestId(v string) *DescribeFieldStatisticsResponseBody
	GetRequestId() *string
}

type DescribeFieldStatisticsResponseBody struct {
	// The information about servers that are returned.
	GroupedFields *DescribeFieldStatisticsResponseBodyGroupedFields `json:"GroupedFields,omitempty" xml:"GroupedFields,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFieldStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsResponseBody) GetGroupedFields() *DescribeFieldStatisticsResponseBodyGroupedFields {
	return s.GroupedFields
}

func (s *DescribeFieldStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFieldStatisticsResponseBody) SetGroupedFields(v *DescribeFieldStatisticsResponseBodyGroupedFields) *DescribeFieldStatisticsResponseBody {
	s.GroupedFields = v
	return s
}

func (s *DescribeFieldStatisticsResponseBody) SetRequestId(v string) *DescribeFieldStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFieldStatisticsResponseBodyGroupedFields struct {
	// The number of assets that are deployed on Alibaba Cloud.
	//
	// example:
	//
	// 100
	AliYunInstanceCount *int32 `json:"AliYunInstanceCount,omitempty" xml:"AliYunInstanceCount,omitempty"`
	// The number of servers.
	//
	// example:
	//
	// 100
	AwsInstanceCount *int32 `json:"AwsInstanceCount,omitempty" xml:"AwsInstanceCount,omitempty"`
	// The number of third-party cloud servers.
	//
	// example:
	//
	// 5
	AzureInstanceCount *int32 `json:"AzureInstanceCount,omitempty" xml:"AzureInstanceCount,omitempty"`
	// The number of cores of exposed assets.
	//
	// example:
	//
	// 30
	ExposedInstanceCoreCount *int64 `json:"ExposedInstanceCoreCount,omitempty" xml:"ExposedInstanceCoreCount,omitempty"`
	// The number of exposed servers.
	//
	// example:
	//
	// 1
	ExposedInstanceCount *int32 `json:"ExposedInstanceCount,omitempty" xml:"ExposedInstanceCount,omitempty"`
	// The number of assets whose importance is common.
	//
	// example:
	//
	// 10
	GeneralAssetCount *int32 `json:"GeneralAssetCount,omitempty" xml:"GeneralAssetCount,omitempty"`
	// The number of instances that are provisioned by third-party providers.
	//
	// example:
	//
	// 10
	GoogleInstanceCount *int32 `json:"GoogleInstanceCount,omitempty" xml:"GoogleInstanceCount,omitempty"`
	// The number of server groups.
	//
	// example:
	//
	// 20
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The number of instances that are provisioned by third-party providers.
	//
	// example:
	//
	// 0
	HuaweiInstanceCount *int32 `json:"HuaweiInstanceCount,omitempty" xml:"HuaweiInstanceCount,omitempty"`
	// The number of assets that can be protected by Security Center.
	//
	// example:
	//
	// 100
	IdcInstanceCount *int32 `json:"IdcInstanceCount,omitempty" xml:"IdcInstanceCount,omitempty"`
	// The number of assets whose importance is important.
	//
	// example:
	//
	// 10
	ImportantAssetCount *int32 `json:"ImportantAssetCount,omitempty" xml:"ImportantAssetCount,omitempty"`
	// The number of cores of assets in the specified asset type. If the asset type is not specified, the value of this parameter indicates the total number of cores of servers and Alibaba Cloud services within your account.
	//
	// example:
	//
	// 301
	InstanceCoreCount *int64 `json:"InstanceCoreCount,omitempty" xml:"InstanceCoreCount,omitempty"`
	// The total number of assets of the specified type. If no asset types are specified, this parameter indicates the total number of all servers and Alibaba Cloud services within your account.
	//
	// example:
	//
	// 100
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The total number of tasks for the specified type of assets. If no asset types are specified, this parameter indicates the total number of all servers and Alibaba Cloud services within your account.
	//
	// example:
	//
	// 10
	InstanceSyncTaskCount *int32 `json:"InstanceSyncTaskCount,omitempty" xml:"InstanceSyncTaskCount,omitempty"`
	// The number of cores of new servers.
	//
	// example:
	//
	// 30
	NewInstanceCoreCount *int64 `json:"NewInstanceCoreCount,omitempty" xml:"NewInstanceCoreCount,omitempty"`
	// The number of newly added servers.
	//
	// example:
	//
	// 10
	NewInstanceCount *int32 `json:"NewInstanceCount,omitempty" xml:"NewInstanceCount,omitempty"`
	// The number of servers on which no risks are detected.
	//
	// example:
	//
	// 10
	NoRiskInstanceCount *int32 `json:"NoRiskInstanceCount,omitempty" xml:"NoRiskInstanceCount,omitempty"`
	// The number of assets that are not added to Security Center of the specified asset type.
	//
	// example:
	//
	// 10
	NotBindMachineInstanceCount *int32 `json:"NotBindMachineInstanceCount,omitempty" xml:"NotBindMachineInstanceCount,omitempty"`
	// The number of cores of servers that are not started.
	//
	// example:
	//
	// 30
	NotRunningStatusCoreCount *int64 `json:"NotRunningStatusCoreCount,omitempty" xml:"NotRunningStatusCoreCount,omitempty"`
	// The number of servers that are shut down.
	//
	// example:
	//
	// 10
	NotRunningStatusCount *int32 `json:"NotRunningStatusCount,omitempty" xml:"NotRunningStatusCount,omitempty"`
	// The number of servers whose Security Center agent status is Offline.
	//
	// example:
	//
	// 21
	OfflineInstanceCount *int32 `json:"OfflineInstanceCount,omitempty" xml:"OfflineInstanceCount,omitempty"`
	// The number of servers outside the cloud.
	//
	// example:
	//
	// 20
	OutMachineInstanceCount *int32 `json:"OutMachineInstanceCount,omitempty" xml:"OutMachineInstanceCount,omitempty"`
	// The number of servers for which the Security Center agent suspends protection.
	//
	// example:
	//
	// 10
	PauseInstanceCount *int32 `json:"PauseInstanceCount,omitempty" xml:"PauseInstanceCount,omitempty"`
	// The number of regions to which the servers belong.
	//
	// example:
	//
	// 11
	RegionCount *int32 `json:"RegionCount,omitempty" xml:"RegionCount,omitempty"`
	// The number of cores of vulnerable assets.
	//
	// example:
	//
	// 201
	RiskInstanceCoreCount *int64 `json:"RiskInstanceCoreCount,omitempty" xml:"RiskInstanceCoreCount,omitempty"`
	// The number of assets that are at risk.
	//
	// example:
	//
	// 90
	RiskInstanceCount *int32 `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
	// The total number of cloud services that are protected by Security Center.
	//
	// example:
	//
	// 10
	TencentInstanceCount *int32 `json:"TencentInstanceCount,omitempty" xml:"TencentInstanceCount,omitempty"`
	// The number of assets whose importance is test.
	//
	// example:
	//
	// 10
	TestAssetCount *int32 `json:"TestAssetCount,omitempty" xml:"TestAssetCount,omitempty"`
	// The number of simple application servers.
	//
	// example:
	//
	// 2
	TripartiteInstanceCount *int32 `json:"TripartiteInstanceCount,omitempty" xml:"TripartiteInstanceCount,omitempty"`
	// The number of servers that are in the Unknown state.
	//
	// example:
	//
	// 1
	UnKnowStatusInstanceCount *int32 `json:"UnKnowStatusInstanceCount,omitempty" xml:"UnKnowStatusInstanceCount,omitempty"`
	// The number of cores of unprotected assets.
	//
	// example:
	//
	// 30
	UnprotectedInstanceCoreCount *int64 `json:"UnprotectedInstanceCoreCount,omitempty" xml:"UnprotectedInstanceCoreCount,omitempty"`
	// The number of unprotected assets.
	//
	// example:
	//
	// 10
	UnprotectedInstanceCount *int32 `json:"UnprotectedInstanceCount,omitempty" xml:"UnprotectedInstanceCount,omitempty"`
	// The number of instances that are provisioned by third-party providers.
	//
	// example:
	//
	// 20
	VolcengineInstanceCount *int32 `json:"VolcengineInstanceCount,omitempty" xml:"VolcengineInstanceCount,omitempty"`
	// The number of virtual private clouds (VPCs).
	//
	// example:
	//
	// 5
	VpcCount *int32 `json:"VpcCount,omitempty" xml:"VpcCount,omitempty"`
}

func (s DescribeFieldStatisticsResponseBodyGroupedFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldStatisticsResponseBodyGroupedFields) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetAliYunInstanceCount() *int32 {
	return s.AliYunInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetAwsInstanceCount() *int32 {
	return s.AwsInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetAzureInstanceCount() *int32 {
	return s.AzureInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetExposedInstanceCoreCount() *int64 {
	return s.ExposedInstanceCoreCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetExposedInstanceCount() *int32 {
	return s.ExposedInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetGeneralAssetCount() *int32 {
	return s.GeneralAssetCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetGoogleInstanceCount() *int32 {
	return s.GoogleInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetHuaweiInstanceCount() *int32 {
	return s.HuaweiInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetIdcInstanceCount() *int32 {
	return s.IdcInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetImportantAssetCount() *int32 {
	return s.ImportantAssetCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetInstanceCoreCount() *int64 {
	return s.InstanceCoreCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetInstanceSyncTaskCount() *int32 {
	return s.InstanceSyncTaskCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetNewInstanceCoreCount() *int64 {
	return s.NewInstanceCoreCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetNewInstanceCount() *int32 {
	return s.NewInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetNoRiskInstanceCount() *int32 {
	return s.NoRiskInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetNotBindMachineInstanceCount() *int32 {
	return s.NotBindMachineInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetNotRunningStatusCoreCount() *int64 {
	return s.NotRunningStatusCoreCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetNotRunningStatusCount() *int32 {
	return s.NotRunningStatusCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetOfflineInstanceCount() *int32 {
	return s.OfflineInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetOutMachineInstanceCount() *int32 {
	return s.OutMachineInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetPauseInstanceCount() *int32 {
	return s.PauseInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetRegionCount() *int32 {
	return s.RegionCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetRiskInstanceCoreCount() *int64 {
	return s.RiskInstanceCoreCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetRiskInstanceCount() *int32 {
	return s.RiskInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetTencentInstanceCount() *int32 {
	return s.TencentInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetTestAssetCount() *int32 {
	return s.TestAssetCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetTripartiteInstanceCount() *int32 {
	return s.TripartiteInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetUnKnowStatusInstanceCount() *int32 {
	return s.UnKnowStatusInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetUnprotectedInstanceCoreCount() *int64 {
	return s.UnprotectedInstanceCoreCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetUnprotectedInstanceCount() *int32 {
	return s.UnprotectedInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetVolcengineInstanceCount() *int32 {
	return s.VolcengineInstanceCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) GetVpcCount() *int32 {
	return s.VpcCount
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetAliYunInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.AliYunInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetAwsInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.AwsInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetAzureInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.AzureInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetExposedInstanceCoreCount(v int64) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.ExposedInstanceCoreCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetExposedInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.ExposedInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetGeneralAssetCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.GeneralAssetCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetGoogleInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.GoogleInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetGroupCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.GroupCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetHuaweiInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.HuaweiInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetIdcInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.IdcInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetImportantAssetCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.ImportantAssetCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetInstanceCoreCount(v int64) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.InstanceCoreCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.InstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetInstanceSyncTaskCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.InstanceSyncTaskCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNewInstanceCoreCount(v int64) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NewInstanceCoreCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNewInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NewInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNoRiskInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NoRiskInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNotBindMachineInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NotBindMachineInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNotRunningStatusCoreCount(v int64) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NotRunningStatusCoreCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNotRunningStatusCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NotRunningStatusCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetOfflineInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.OfflineInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetOutMachineInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.OutMachineInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetPauseInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.PauseInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetRegionCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.RegionCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetRiskInstanceCoreCount(v int64) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.RiskInstanceCoreCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetRiskInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.RiskInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetTencentInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.TencentInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetTestAssetCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.TestAssetCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetTripartiteInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.TripartiteInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetUnKnowStatusInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.UnKnowStatusInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetUnprotectedInstanceCoreCount(v int64) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.UnprotectedInstanceCoreCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetUnprotectedInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.UnprotectedInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetVolcengineInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.VolcengineInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetVpcCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.VpcCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) Validate() error {
	return dara.Validate(s)
}
