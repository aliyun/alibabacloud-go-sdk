// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeQuotaPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateComputeQuotaPlanRequest
	GetName() *string
	SetQuota(v *CreateComputeQuotaPlanRequestQuota) *CreateComputeQuotaPlanRequest
	GetQuota() *CreateComputeQuotaPlanRequestQuota
}

type CreateComputeQuotaPlanRequest struct {
	// This parameter is required.
	Name  *string                             `json:"name,omitempty" xml:"name,omitempty"`
	Quota *CreateComputeQuotaPlanRequestQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s CreateComputeQuotaPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequest) GetName() *string {
	return s.Name
}

func (s *CreateComputeQuotaPlanRequest) GetQuota() *CreateComputeQuotaPlanRequestQuota {
	return s.Quota
}

func (s *CreateComputeQuotaPlanRequest) SetName(v string) *CreateComputeQuotaPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateComputeQuotaPlanRequest) SetQuota(v *CreateComputeQuotaPlanRequestQuota) *CreateComputeQuotaPlanRequest {
	s.Quota = v
	return s
}

func (s *CreateComputeQuotaPlanRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateComputeQuotaPlanRequestQuota struct {
	Parameter        *CreateComputeQuotaPlanRequestQuotaParameter          `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	SubQuotaInfoList []*CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
}

func (s CreateComputeQuotaPlanRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuota) GetParameter() *CreateComputeQuotaPlanRequestQuotaParameter {
	return s.Parameter
}

func (s *CreateComputeQuotaPlanRequestQuota) GetSubQuotaInfoList() []*CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	return s.SubQuotaInfoList
}

func (s *CreateComputeQuotaPlanRequestQuota) SetParameter(v *CreateComputeQuotaPlanRequestQuotaParameter) *CreateComputeQuotaPlanRequestQuota {
	s.Parameter = v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuota) SetSubQuotaInfoList(v []*CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) *CreateComputeQuotaPlanRequestQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuota) Validate() error {
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

type CreateComputeQuotaPlanRequestQuotaParameter struct {
	// This parameter is required.
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
}

func (s CreateComputeQuotaPlanRequestQuotaParameter) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuotaParameter) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuotaParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *CreateComputeQuotaPlanRequestQuotaParameter) SetElasticReservedCU(v int64) *CreateComputeQuotaPlanRequestQuotaParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaParameter) Validate() error {
	return dara.Validate(s)
}

type CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList struct {
	// This parameter is required.
	NickName  *string                                                      `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GetNickName() *string {
	return s.NickName
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) GetParameter() *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	return s.Parameter
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetNickName(v string) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) SetParameter(v *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoList) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter struct {
	// This parameter is required.
	ElasticReservedCU *int64 `json:"elasticReservedCU,omitempty" xml:"elasticReservedCU,omitempty"`
	// This parameter is required.
	MaxCU *int64 `json:"maxCU,omitempty" xml:"maxCU,omitempty"`
	// This parameter is required.
	MinCU *int64 `json:"minCU,omitempty" xml:"minCU,omitempty"`
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GoString() string {
	return s.String()
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GetElasticReservedCU() *int64 {
	return s.ElasticReservedCU
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GetMaxCU() *int64 {
	return s.MaxCU
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) GetMinCU() *int64 {
	return s.MinCU
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetElasticReservedCU(v int64) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.ElasticReservedCU = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMaxCU(v int64) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MaxCU = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) SetMinCU(v int64) *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter {
	s.MinCU = &v
	return s
}

func (s *CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter) Validate() error {
	return dara.Validate(s)
}
