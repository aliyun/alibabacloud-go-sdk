// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditCountDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetAuditCountDistributionRequest
	GetBeginDate() *string
	SetDbId(v int32) *GetAuditCountDistributionRequest
	GetDbId() *int32
	SetEndDate(v string) *GetAuditCountDistributionRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetAuditCountDistributionRequest
	GetInstanceId() *string
	SetLang(v string) *GetAuditCountDistributionRequest
	GetLang() *string
	SetRegionId(v string) *GetAuditCountDistributionRequest
	GetRegionId() *string
}

type GetAuditCountDistributionRequest struct {
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 0
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
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

func (s GetAuditCountDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditCountDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetAuditCountDistributionRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetAuditCountDistributionRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetAuditCountDistributionRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetAuditCountDistributionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuditCountDistributionRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAuditCountDistributionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAuditCountDistributionRequest) SetBeginDate(v string) *GetAuditCountDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetDbId(v int32) *GetAuditCountDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetEndDate(v string) *GetAuditCountDistributionRequest {
	s.EndDate = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetInstanceId(v string) *GetAuditCountDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetLang(v string) *GetAuditCountDistributionRequest {
	s.Lang = &v
	return s
}

func (s *GetAuditCountDistributionRequest) SetRegionId(v string) *GetAuditCountDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetAuditCountDistributionRequest) Validate() error {
	return dara.Validate(s)
}
