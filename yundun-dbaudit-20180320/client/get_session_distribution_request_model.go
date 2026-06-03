// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetSessionDistributionRequest
	GetBeginDate() *string
	SetDbId(v int32) *GetSessionDistributionRequest
	GetDbId() *int32
	SetEndDate(v string) *GetSessionDistributionRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetSessionDistributionRequest
	GetInstanceId() *string
	SetLang(v string) *GetSessionDistributionRequest
	GetLang() *string
	SetRegionId(v string) *GetSessionDistributionRequest
	GetRegionId() *string
}

type GetSessionDistributionRequest struct {
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

func (s GetSessionDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSessionDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetSessionDistributionRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetSessionDistributionRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetSessionDistributionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSessionDistributionRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSessionDistributionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSessionDistributionRequest) SetBeginDate(v string) *GetSessionDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetSessionDistributionRequest) SetDbId(v int32) *GetSessionDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetSessionDistributionRequest) SetEndDate(v string) *GetSessionDistributionRequest {
	s.EndDate = &v
	return s
}

func (s *GetSessionDistributionRequest) SetInstanceId(v string) *GetSessionDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSessionDistributionRequest) SetLang(v string) *GetSessionDistributionRequest {
	s.Lang = &v
	return s
}

func (s *GetSessionDistributionRequest) SetRegionId(v string) *GetSessionDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetSessionDistributionRequest) Validate() error {
	return dara.Validate(s)
}
