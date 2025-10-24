// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateComputeQuotaPlanRequest
	GetName() *string
	SetQuota(v *UpdateComputeQuotaPlanRequestQuota) *UpdateComputeQuotaPlanRequest
	GetQuota() *UpdateComputeQuotaPlanRequestQuota
}

type UpdateComputeQuotaPlanRequest struct {
	// The name of quota plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameters of quota plan.
	Quota *UpdateComputeQuotaPlanRequestQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s UpdateComputeQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequest) GetName() *string {
	return s.Name
}

func (s *UpdateComputeQuotaPlanRequest) GetQuota() *UpdateComputeQuotaPlanRequestQuota {
	return s.Quota
}

func (s *UpdateComputeQuotaPlanRequest) SetName(v string) *UpdateComputeQuotaPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequest) SetQuota(v *UpdateComputeQuotaPlanRequestQuota) *UpdateComputeQuotaPlanRequest {
	s.Quota = v
	return s
}

func (s *UpdateComputeQuotaPlanRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeQuotaPlanRequestQuota struct {
	// The parameters of level-1 quota.
	Parameter *UpdateComputeQuotaPlanRequestQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The list of level-2 quotas.
	SubQuotaInfoList []*UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
}

func (s UpdateComputeQuotaPlanRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuota) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuota) GetParameter() *UpdateComputeQuotaPlanRequestQuotaParameter {
	return s.Parameter
}

func (s *UpdateComputeQuotaPlanRequestQuota) GetSubQuotaInfoList() []*UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *UpdateComputeQuotaPlanRequestQuota) SetParameter(v *UpdateComputeQuotaPlanRequestQuotaParameter) *UpdateComputeQuotaPlanRequestQuota {
	s.Parameter = v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuota) SetSubQuotaInfoList(v []*UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) *UpdateComputeQuotaPlanRequestQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuota) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
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

type UpdateComputeQuotaPlanRequestQuotaParameter struct {
	// The value of elastic Reserved CUs in the level-1 quota.
	//
	// > The default value is 0. The maximum value of this parameter must be equal to the number of subscription-based reserved CUs and cannot exceed 10,000 CUs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
}

func (s UpdateComputeQuotaPlanRequestQuotaParameter) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuotaParameter) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuotaParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *UpdateComputeQuotaPlanRequestQuotaParameter) SetElasticReservedCU(v int64) *UpdateComputeQuotaPlanRequestQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaParameter) Validate() error {
	return dara.Validate(s)
}

type UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList struct {
	// The nickname of the level-2 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// os_ComputeQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
	Parameter *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GetParameter() *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetNickName(v string) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetParameter(v *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoList) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter struct {
	// The value of elastic Reserved CUs.
	//
	// > The total number of elastically reserved CUs in all the level-2 quotas is equal to the number of elastically reserved CUs in the level-1 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// The value of maxCU in Reserved CUs.
	//
	// > The value of maxCU must be less than or equal to the value of maxCU in the level-1 quota that you purchased.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// The value of minCU in Reserved CUs.
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
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *UpdateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}
