// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSubQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubQuotaInfoList(v []*UpdateComputeSubQuotaRequestSubQuotaInfoList) *UpdateComputeSubQuotaRequest
	GetSubQuotaInfoList() []*UpdateComputeSubQuotaRequestSubQuotaInfoList
}

type UpdateComputeSubQuotaRequest struct {
	// The list of level-2 quotas.
	SubQuotaInfoList []*UpdateComputeSubQuotaRequestSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
}

func (s UpdateComputeSubQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSubQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaRequest) GetSubQuotaInfoList() []*UpdateComputeSubQuotaRequestSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *UpdateComputeSubQuotaRequest) SetSubQuotaInfoList(v []*UpdateComputeSubQuotaRequestSubQuotaInfoList) *UpdateComputeSubQuotaRequest {
	s.SubQuotaInfoList = v
	return s
}

func (s *UpdateComputeSubQuotaRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateComputeSubQuotaRequestSubQuotaInfoList struct {
	// The nickname of the level-2 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// os_ComputeQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
	Parameter *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The type of quota.
	//
	// >
	//
	// > - FUXI_OFFLINE(default) : Quotas of this type are used to run batch jobs.
	//
	// example:
	//
	// FUXI_OFFLINE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) GetParameter() *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) GetType() *string {
	return s.Type
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) SetNickName(v string) *UpdateComputeSubQuotaRequestSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) SetParameter(v *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) *UpdateComputeSubQuotaRequestSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) SetType(v string) *UpdateComputeSubQuotaRequestSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoList) Validate() error {
	return dara.Validate(s)
}

type UpdateComputeSubQuotaRequestSubQuotaInfoListParameter struct {
	// Specifies whether to enable the priority feature.
	//
	// example:
	//
	// false
	EnablePriority *bool `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	// Specifies whether the quota is strongly exclusive.
	//
	// example:
	//
	// false
	ForceReservedMin *bool `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	// The value of minCU in Reserved CUs.
	//
	// > The value of maxCU must be less than or equal to the value of maxCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// >
	//
	// >- The total value of minCU in all the level-2 quotas is equal to the value of minCU in the level-1 quota.
	//
	// >- The value of minCU must be less than or equal to the value of maxCU in the level-2 quota and less than or equal to the value of minCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
	// Scheduling policy of the quota.
	//
	// example:
	//
	// Fifo/Fair
	SchedulerType *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	// The upper limit for CUs that can be concurrently used by a job scheduled to the quota.
	//
	// example:
	//
	// 10
	SingleJobCULimit *int64 `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GetEnablePriority() *bool {
	return s.EnablePriority
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GetForceReservedMin() *bool {
	return s.ForceReservedMin
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) GetSingleJobCULimit() *int64 {
	return s.SingleJobCULimit
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetEnablePriority(v bool) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.EnablePriority = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetForceReservedMin(v bool) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.ForceReservedMin = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetMaxCU(v int64) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetMinCU(v int64) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetSchedulerType(v string) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.SchedulerType = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) SetSingleJobCULimit(v int64) *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter {
	s.SingleJobCULimit = &v
	return s
}

func (s *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}
