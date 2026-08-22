// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) *DescribeOpenSearchInstancesResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchInstancesResponseBodyData) *DescribeOpenSearchInstancesResponseBody
	GetData() *DescribeOpenSearchInstancesResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchInstancesResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchInstancesResponseBody struct {
	AccessDeniedDetail *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *DescribeOpenSearchInstancesResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInstancesResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchInstancesResponseBody) GetData() *DescribeOpenSearchInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchInstancesResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) *DescribeOpenSearchInstancesResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBody) SetData(v *DescribeOpenSearchInstancesResponseBodyData) *DescribeOpenSearchInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBody) SetRequestId(v string) *DescribeOpenSearchInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchInstancesResponseBodyData struct {
	Instances []*DescribeOpenSearchInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kmMV9kamx92yNWehxph5Fw
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 130
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeOpenSearchInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInstancesResponseBodyData) GetInstances() []*DescribeOpenSearchInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *DescribeOpenSearchInstancesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeOpenSearchInstancesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOpenSearchInstancesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeOpenSearchInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOpenSearchInstancesResponseBodyData) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *DescribeOpenSearchInstancesResponseBodyData) SetInstances(v []*DescribeOpenSearchInstancesResponseBodyDataInstances) *DescribeOpenSearchInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyData) SetMaxResults(v int32) *DescribeOpenSearchInstancesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyData) SetNextToken(v string) *DescribeOpenSearchInstancesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyData) SetPageNumber(v int32) *DescribeOpenSearchInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyData) SetPageSize(v int32) *DescribeOpenSearchInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyData) SetTotalNumber(v int32) *DescribeOpenSearchInstancesResponseBodyData {
	s.TotalNumber = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyData) Validate() error {
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

type DescribeOpenSearchInstancesResponseBodyDataInstances struct {
	// example:
	//
	// t1222576965886205
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 4000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 2026-06-08T07:19:05.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 3
	DataNodeCount *int32 `json:"DataNodeCount,omitempty" xml:"DataNodeCount,omitempty"`
	// example:
	//
	// 我的 Supabase 项目
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// pxc-shrdb7a2t8w3c1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 16
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
	// example:
	//
	// 1
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 8 核 32 GB
	SpecDisplay *string `json:"SpecDisplay,omitempty" xml:"SpecDisplay,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 500
	StorageSizeGB *int32 `json:"StorageSizeGB,omitempty" xml:"StorageSizeGB,omitempty"`
}

func (s DescribeOpenSearchInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetDataNodeCount() *int32 {
	return s.DataNodeCount
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetNetType() *string {
	return s.NetType
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetSpecDisplay() *string {
	return s.SpecDisplay
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) GetStorageSizeGB() *int32 {
	return s.StorageSizeGB
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetAvailabilityZone(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.AvailabilityZone = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetChargeType(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetCpu(v int32) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.Cpu = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetCreateTime(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetDataNodeCount(v int32) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.DataNodeCount = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetDescription(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.Description = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetEngineVersion(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetInstanceId(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetMemoryGB(v int32) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.MemoryGB = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetNetType(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.NetType = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetRegionId(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetSpecDisplay(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.SpecDisplay = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetStatus(v string) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) SetStorageSizeGB(v int32) *DescribeOpenSearchInstancesResponseBodyDataInstances {
	s.StorageSizeGB = &v
	return s
}

func (s *DescribeOpenSearchInstancesResponseBodyDataInstances) Validate() error {
	return dara.Validate(s)
}
