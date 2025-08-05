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
	// The name of quota plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameters of quota plan.
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
	return dara.Validate(s)
}

type CreateComputeQuotaPlanRequestQuota struct {
	// The parameters of level-1 quota.
	Parameter *CreateComputeQuotaPlanRequestQuotaParameter `json:"parameter,omitempty" xml:"parameter,omitempty" type:"Struct"`
	// The list of level-2 quotas.
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
	return dara.Validate(s)
}

type CreateComputeQuotaPlanRequestQuotaParameter struct {
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
	// The nickname of the level-2 quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// os_ComputeQuota
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The parameters of the level-2 quota.
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
	return dara.Validate(s)
}

type CreateComputeQuotaPlanRequestQuotaSubQuotaInfoListParameter struct {
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
