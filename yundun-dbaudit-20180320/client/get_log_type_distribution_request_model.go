// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogTypeDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetLogTypeDistributionRequest
	GetBeginDate() *string
	SetDbId(v int32) *GetLogTypeDistributionRequest
	GetDbId() *int32
	SetEndDate(v string) *GetLogTypeDistributionRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetLogTypeDistributionRequest
	GetInstanceId() *string
	SetLang(v string) *GetLogTypeDistributionRequest
	GetLang() *string
	SetRegionId(v string) *GetLogTypeDistributionRequest
	GetRegionId() *string
}

type GetLogTypeDistributionRequest struct {
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

func (s GetLogTypeDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogTypeDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetLogTypeDistributionRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetLogTypeDistributionRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetLogTypeDistributionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLogTypeDistributionRequest) GetLang() *string {
	return s.Lang
}

func (s *GetLogTypeDistributionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLogTypeDistributionRequest) SetBeginDate(v string) *GetLogTypeDistributionRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetDbId(v int32) *GetLogTypeDistributionRequest {
	s.DbId = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetEndDate(v string) *GetLogTypeDistributionRequest {
	s.EndDate = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetInstanceId(v string) *GetLogTypeDistributionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetLang(v string) *GetLogTypeDistributionRequest {
	s.Lang = &v
	return s
}

func (s *GetLogTypeDistributionRequest) SetRegionId(v string) *GetLogTypeDistributionRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogTypeDistributionRequest) Validate() error {
	return dara.Validate(s)
}
