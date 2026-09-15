// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetComputeInstanceResponseBody
	GetCode() *int64
	SetData(v *GetComputeInstanceResponseBodyData) *GetComputeInstanceResponseBody
	GetData() *GetComputeInstanceResponseBodyData
	SetRequestId(v string) *GetComputeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetComputeInstanceResponseBody
	GetSuccess() *bool
}

type GetComputeInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *int64                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetComputeInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 019F5939-507B-57C1-B82C-E55821456FA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetComputeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeInstanceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetComputeInstanceResponseBody) GetData() *GetComputeInstanceResponseBodyData {
	return s.Data
}

func (s *GetComputeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetComputeInstanceResponseBody) SetCode(v int64) *GetComputeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetComputeInstanceResponseBody) SetData(v *GetComputeInstanceResponseBodyData) *GetComputeInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeInstanceResponseBody) SetRequestId(v string) *GetComputeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeInstanceResponseBody) SetSuccess(v bool) *GetComputeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetComputeInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeInstanceResponseBodyData struct {
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
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2027-09-02T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// alikafka_streaming-cn-pe333xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// streaming-prod
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
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

func (s GetComputeInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetComputeInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeInstanceResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetComputeInstanceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetComputeInstanceResponseBodyData) GetCuLimitSum() *float64 {
	return s.CuLimitSum
}

func (s *GetComputeInstanceResponseBodyData) GetCuReservedSum() *float64 {
	return s.CuReservedSum
}

func (s *GetComputeInstanceResponseBodyData) GetCuUsedSum() *float64 {
	return s.CuUsedSum
}

func (s *GetComputeInstanceResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetComputeInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetComputeInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetComputeInstanceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeInstanceResponseBodyData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *GetComputeInstanceResponseBodyData) GetTotalJobs() *int64 {
	return s.TotalJobs
}

func (s *GetComputeInstanceResponseBodyData) GetTotalRunningJobs() *int64 {
	return s.TotalRunningJobs
}

func (s *GetComputeInstanceResponseBodyData) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetComputeInstanceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetComputeInstanceResponseBodyData) SetChargeType(v string) *GetComputeInstanceResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetCreateTime(v string) *GetComputeInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetCuLimitSum(v float64) *GetComputeInstanceResponseBodyData {
	s.CuLimitSum = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetCuReservedSum(v float64) *GetComputeInstanceResponseBodyData {
	s.CuReservedSum = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetCuUsedSum(v float64) *GetComputeInstanceResponseBodyData {
	s.CuUsedSum = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetExpireTime(v string) *GetComputeInstanceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetInstanceId(v string) *GetComputeInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetInstanceName(v string) *GetComputeInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetRegionId(v string) *GetComputeInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetServiceStatus(v string) *GetComputeInstanceResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetTotalJobs(v int64) *GetComputeInstanceResponseBodyData {
	s.TotalJobs = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetTotalRunningJobs(v int64) *GetComputeInstanceResponseBodyData {
	s.TotalRunningJobs = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetVSwitchIds(v []*string) *GetComputeInstanceResponseBodyData {
	s.VSwitchIds = v
	return s
}

func (s *GetComputeInstanceResponseBodyData) SetVpcId(v string) *GetComputeInstanceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetComputeInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
