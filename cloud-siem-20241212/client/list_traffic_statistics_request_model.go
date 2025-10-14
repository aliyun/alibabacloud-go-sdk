// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListTrafficStatisticsRequest
	GetLang() *string
	SetLogUserIds(v []*int64) *ListTrafficStatisticsRequest
	GetLogUserIds() []*int64
	SetProductId(v string) *ListTrafficStatisticsRequest
	GetProductId() *string
	SetRegionId(v string) *ListTrafficStatisticsRequest
	GetRegionId() *string
	SetRegionTag(v int32) *ListTrafficStatisticsRequest
	GetRegionTag() *int32
	SetRoleFor(v int64) *ListTrafficStatisticsRequest
	GetRoleFor() *int64
	SetTrafficStatisticPeriod(v string) *ListTrafficStatisticsRequest
	GetTrafficStatisticPeriod() *string
	SetTrafficStatisticPeriodType(v string) *ListTrafficStatisticsRequest
	GetTrafficStatisticPeriodType() *string
	SetTrafficStatisticType(v string) *ListTrafficStatisticsRequest
	GetTrafficStatisticType() *string
}

type ListTrafficStatisticsRequest struct {
	// example:
	//
	// zh。
	Lang       *string  `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LogUserIds []*int64 `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1。
	RegionTag *int32 `json:"RegionTag,omitempty" xml:"RegionTag,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 30。
	TrafficStatisticPeriod *string `json:"TrafficStatisticPeriod,omitempty" xml:"TrafficStatisticPeriod,omitempty"`
	// example:
	//
	// day。
	TrafficStatisticPeriodType *string `json:"TrafficStatisticPeriodType,omitempty" xml:"TrafficStatisticPeriodType,omitempty"`
	// example:
	//
	// Region。
	TrafficStatisticType *string `json:"TrafficStatisticType,omitempty" xml:"TrafficStatisticType,omitempty"`
}

func (s ListTrafficStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTrafficStatisticsRequest) GetLogUserIds() []*int64 {
	return s.LogUserIds
}

func (s *ListTrafficStatisticsRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListTrafficStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTrafficStatisticsRequest) GetRegionTag() *int32 {
	return s.RegionTag
}

func (s *ListTrafficStatisticsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListTrafficStatisticsRequest) GetTrafficStatisticPeriod() *string {
	return s.TrafficStatisticPeriod
}

func (s *ListTrafficStatisticsRequest) GetTrafficStatisticPeriodType() *string {
	return s.TrafficStatisticPeriodType
}

func (s *ListTrafficStatisticsRequest) GetTrafficStatisticType() *string {
	return s.TrafficStatisticType
}

func (s *ListTrafficStatisticsRequest) SetLang(v string) *ListTrafficStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *ListTrafficStatisticsRequest) SetLogUserIds(v []*int64) *ListTrafficStatisticsRequest {
	s.LogUserIds = v
	return s
}

func (s *ListTrafficStatisticsRequest) SetProductId(v string) *ListTrafficStatisticsRequest {
	s.ProductId = &v
	return s
}

func (s *ListTrafficStatisticsRequest) SetRegionId(v string) *ListTrafficStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTrafficStatisticsRequest) SetRegionTag(v int32) *ListTrafficStatisticsRequest {
	s.RegionTag = &v
	return s
}

func (s *ListTrafficStatisticsRequest) SetRoleFor(v int64) *ListTrafficStatisticsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListTrafficStatisticsRequest) SetTrafficStatisticPeriod(v string) *ListTrafficStatisticsRequest {
	s.TrafficStatisticPeriod = &v
	return s
}

func (s *ListTrafficStatisticsRequest) SetTrafficStatisticPeriodType(v string) *ListTrafficStatisticsRequest {
	s.TrafficStatisticPeriodType = &v
	return s
}

func (s *ListTrafficStatisticsRequest) SetTrafficStatisticType(v string) *ListTrafficStatisticsRequest {
	s.TrafficStatisticType = &v
	return s
}

func (s *ListTrafficStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
