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
	if s.SubQuotaInfoList != nil {
		for _, item := range s.SubQuotaInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateComputeSubQuotaRequestSubQuotaInfoList struct {
	// This parameter is required.
	NickName  *string                                                `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter *UpdateComputeSubQuotaRequestSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	Type      *string                                                `json:"type,omitempty" xml:"type,omitempty"`
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
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeSubQuotaRequestSubQuotaInfoListParameter struct {
	EnablePriority   *bool `json:"enablePriority,omitempty" xml:"enablePriority,omitempty"`
	ForceReservedMin *bool `json:"forceReservedMin,omitempty" xml:"forceReservedMin,omitempty"`
	// This parameter is required.
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// This parameter is required.
	MinCU            *int64  `json:"minCU,omitempty" xml:"minCU,omitempty"`
	SchedulerType    *string `json:"schedulerType,omitempty" xml:"schedulerType,omitempty"`
	SingleJobCULimit *int64  `json:"singleJobCULimit,omitempty" xml:"singleJobCULimit,omitempty"`
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
