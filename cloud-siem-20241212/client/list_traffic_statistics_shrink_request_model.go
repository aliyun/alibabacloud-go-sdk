// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficStatisticsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListTrafficStatisticsShrinkRequest
	GetLang() *string
	SetLogUserIdsShrink(v string) *ListTrafficStatisticsShrinkRequest
	GetLogUserIdsShrink() *string
	SetProductId(v string) *ListTrafficStatisticsShrinkRequest
	GetProductId() *string
	SetRegionId(v string) *ListTrafficStatisticsShrinkRequest
	GetRegionId() *string
	SetRegionTag(v int32) *ListTrafficStatisticsShrinkRequest
	GetRegionTag() *int32
	SetRoleFor(v int64) *ListTrafficStatisticsShrinkRequest
	GetRoleFor() *int64
	SetTrafficStatisticPeriod(v string) *ListTrafficStatisticsShrinkRequest
	GetTrafficStatisticPeriod() *string
	SetTrafficStatisticPeriodType(v string) *ListTrafficStatisticsShrinkRequest
	GetTrafficStatisticPeriodType() *string
	SetTrafficStatisticType(v string) *ListTrafficStatisticsShrinkRequest
	GetTrafficStatisticType() *string
}

type ListTrafficStatisticsShrinkRequest struct {
	// example:
	//
	// zh。
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LogUserIdsShrink *string `json:"LogUserIds,omitempty" xml:"LogUserIds,omitempty"`
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

func (s ListTrafficStatisticsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficStatisticsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTrafficStatisticsShrinkRequest) GetLogUserIdsShrink() *string {
	return s.LogUserIdsShrink
}

func (s *ListTrafficStatisticsShrinkRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListTrafficStatisticsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTrafficStatisticsShrinkRequest) GetRegionTag() *int32 {
	return s.RegionTag
}

func (s *ListTrafficStatisticsShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListTrafficStatisticsShrinkRequest) GetTrafficStatisticPeriod() *string {
	return s.TrafficStatisticPeriod
}

func (s *ListTrafficStatisticsShrinkRequest) GetTrafficStatisticPeriodType() *string {
	return s.TrafficStatisticPeriodType
}

func (s *ListTrafficStatisticsShrinkRequest) GetTrafficStatisticType() *string {
	return s.TrafficStatisticType
}

func (s *ListTrafficStatisticsShrinkRequest) SetLang(v string) *ListTrafficStatisticsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetLogUserIdsShrink(v string) *ListTrafficStatisticsShrinkRequest {
	s.LogUserIdsShrink = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetProductId(v string) *ListTrafficStatisticsShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetRegionId(v string) *ListTrafficStatisticsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetRegionTag(v int32) *ListTrafficStatisticsShrinkRequest {
	s.RegionTag = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetRoleFor(v int64) *ListTrafficStatisticsShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetTrafficStatisticPeriod(v string) *ListTrafficStatisticsShrinkRequest {
	s.TrafficStatisticPeriod = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetTrafficStatisticPeriodType(v string) *ListTrafficStatisticsShrinkRequest {
	s.TrafficStatisticPeriodType = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) SetTrafficStatisticType(v string) *ListTrafficStatisticsShrinkRequest {
	s.TrafficStatisticType = &v
	return s
}

func (s *ListTrafficStatisticsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
