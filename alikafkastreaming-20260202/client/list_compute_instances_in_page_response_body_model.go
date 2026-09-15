// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesInPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListComputeInstancesInPageResponseBody
	GetCode() *int64
	SetCurrentPage(v int32) *ListComputeInstancesInPageResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListComputeInstancesInPageResponseBodyData) *ListComputeInstancesInPageResponseBody
	GetData() []*ListComputeInstancesInPageResponseBodyData
	SetPageSize(v int32) *ListComputeInstancesInPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListComputeInstancesInPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListComputeInstancesInPageResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListComputeInstancesInPageResponseBody
	GetTotal() *int64
}

type ListComputeInstancesInPageResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListComputeInstancesInPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9079DD86-09F1-5303-A64F-D9089F96BC16
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListComputeInstancesInPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesInPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesInPageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListComputeInstancesInPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListComputeInstancesInPageResponseBody) GetData() []*ListComputeInstancesInPageResponseBodyData {
	return s.Data
}

func (s *ListComputeInstancesInPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeInstancesInPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeInstancesInPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListComputeInstancesInPageResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListComputeInstancesInPageResponseBody) SetCode(v int64) *ListComputeInstancesInPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBody) SetCurrentPage(v int32) *ListComputeInstancesInPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBody) SetData(v []*ListComputeInstancesInPageResponseBodyData) *ListComputeInstancesInPageResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeInstancesInPageResponseBody) SetPageSize(v int32) *ListComputeInstancesInPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBody) SetRequestId(v string) *ListComputeInstancesInPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBody) SetSuccess(v bool) *ListComputeInstancesInPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBody) SetTotal(v int64) *ListComputeInstancesInPageResponseBody {
	s.Total = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeInstancesInPageResponseBodyData struct {
	// example:
	//
	// false
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" xml:"AutoRenewFlag,omitempty"`
	// example:
	//
	// POST_PAID
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-09-02T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 4.0
	CuLimitSum *float64 `json:"CuLimitSum,omitempty" xml:"CuLimitSum,omitempty"`
	// example:
	//
	// 2.0
	CuReservedSum *float64 `json:"CuReservedSum,omitempty" xml:"CuReservedSum,omitempty"`
	// example:
	//
	// 2.5
	CuUsedSum *float64 `json:"CuUsedSum,omitempty" xml:"CuUsedSum,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2027-09-02T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// alikafka_streaming-cn-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// streaming-prod
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1234567890123456
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// RUNNING
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// example:
	//
	// 3
	TotalJobs *int64 `json:"TotalJobs,omitempty" xml:"TotalJobs,omitempty"`
	// example:
	//
	// 2
	TotalRunningJobs *int64    `json:"TotalRunningJobs,omitempty" xml:"TotalRunningJobs,omitempty"`
	VSwitchIds       []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-bp1abcdefg
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListComputeInstancesInPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesInPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesInPageResponseBodyData) GetAutoRenewFlag() *bool {
	return s.AutoRenewFlag
}

func (s *ListComputeInstancesInPageResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListComputeInstancesInPageResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListComputeInstancesInPageResponseBodyData) GetCuLimitSum() *float64 {
	return s.CuLimitSum
}

func (s *ListComputeInstancesInPageResponseBodyData) GetCuReservedSum() *float64 {
	return s.CuReservedSum
}

func (s *ListComputeInstancesInPageResponseBodyData) GetCuUsedSum() *float64 {
	return s.CuUsedSum
}

func (s *ListComputeInstancesInPageResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListComputeInstancesInPageResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeInstancesInPageResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListComputeInstancesInPageResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *ListComputeInstancesInPageResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeInstancesInPageResponseBodyData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListComputeInstancesInPageResponseBodyData) GetTotalJobs() *int64 {
	return s.TotalJobs
}

func (s *ListComputeInstancesInPageResponseBodyData) GetTotalRunningJobs() *int64 {
	return s.TotalRunningJobs
}

func (s *ListComputeInstancesInPageResponseBodyData) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListComputeInstancesInPageResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *ListComputeInstancesInPageResponseBodyData) SetAutoRenewFlag(v bool) *ListComputeInstancesInPageResponseBodyData {
	s.AutoRenewFlag = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetChargeType(v string) *ListComputeInstancesInPageResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetCreateTime(v string) *ListComputeInstancesInPageResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetCuLimitSum(v float64) *ListComputeInstancesInPageResponseBodyData {
	s.CuLimitSum = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetCuReservedSum(v float64) *ListComputeInstancesInPageResponseBodyData {
	s.CuReservedSum = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetCuUsedSum(v float64) *ListComputeInstancesInPageResponseBodyData {
	s.CuUsedSum = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetExpireTime(v string) *ListComputeInstancesInPageResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetInstanceId(v string) *ListComputeInstancesInPageResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetInstanceName(v string) *ListComputeInstancesInPageResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetOwner(v string) *ListComputeInstancesInPageResponseBodyData {
	s.Owner = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetRegionId(v string) *ListComputeInstancesInPageResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetServiceStatus(v string) *ListComputeInstancesInPageResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetTotalJobs(v int64) *ListComputeInstancesInPageResponseBodyData {
	s.TotalJobs = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetTotalRunningJobs(v int64) *ListComputeInstancesInPageResponseBodyData {
	s.TotalRunningJobs = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetVSwitchIds(v []*string) *ListComputeInstancesInPageResponseBodyData {
	s.VSwitchIds = v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) SetVpcId(v string) *ListComputeInstancesInPageResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *ListComputeInstancesInPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
