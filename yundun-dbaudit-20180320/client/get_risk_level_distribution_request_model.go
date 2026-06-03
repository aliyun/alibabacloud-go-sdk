// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskLevelDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetRiskLevelDistributionRequest
	GetBeginDate() *string
	SetDbId(v int32) *GetRiskLevelDistributionRequest
	GetDbId() *int32
	SetEndDate(v string) *GetRiskLevelDistributionRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetRiskLevelDistributionRequest
	GetInstanceId() *string
	SetLang(v string) *GetRiskLevelDistributionRequest
	GetLang() *string
	SetRegionId(v string) *GetRiskLevelDistributionRequest
	GetRegionId() *string
}

type GetRiskLevelDistributionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 0
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 23:59:59
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetRiskLevelDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRiskLevelDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetRiskLevelDistributionRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetRiskLevelDistributionRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetRiskLevelDistributionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRiskLevelDistributionRequest) GetLang() *string {
	return s.Lang
}

func (s *GetRiskLevelDistributionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRiskLevelDistributionRequest) SetBeginDate(v string) *GetRiskLevelDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetDbId(v int32) *GetRiskLevelDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetEndDate(v string) *GetRiskLevelDistributionRequest {
	s.EndDate = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetInstanceId(v string) *GetRiskLevelDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetLang(v string) *GetRiskLevelDistributionRequest {
	s.Lang = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) SetRegionId(v string) *GetRiskLevelDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetRiskLevelDistributionRequest) Validate() error {
	return dara.Validate(s)
}
