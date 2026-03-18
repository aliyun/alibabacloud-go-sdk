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
	// This parameter is required.
	Name  *string                             `json:"name,omitempty" xml:"name,omitempty"`
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
	Parameter        *UpdateComputeQuotaPlanRequestQuotaParameter          `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
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
	// This parameter is required.
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
	// This parameter is required.
	NickName  *string                                                      `json:"nickName,omitempty" xml:"nickName,omitempty"`
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
	// This parameter is required.
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// This parameter is required.
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// This parameter is required.
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
