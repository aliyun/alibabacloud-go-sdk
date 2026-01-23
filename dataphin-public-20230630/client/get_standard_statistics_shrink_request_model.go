// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardStatisticsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetStandardStatisticsShrinkRequest
	GetOpTenantId() *int64
	SetStatisticsQueryShrink(v string) *GetStandardStatisticsShrinkRequest
	GetStatisticsQueryShrink() *string
}

type GetStandardStatisticsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	StatisticsQueryShrink *string `json:"StatisticsQuery,omitempty" xml:"StatisticsQuery,omitempty"`
}

func (s GetStandardStatisticsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardStatisticsShrinkRequest) GetStatisticsQueryShrink() *string {
	return s.StatisticsQueryShrink
}

func (s *GetStandardStatisticsShrinkRequest) SetOpTenantId(v int64) *GetStandardStatisticsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardStatisticsShrinkRequest) SetStatisticsQueryShrink(v string) *GetStandardStatisticsShrinkRequest {
	s.StatisticsQueryShrink = &v
	return s
}

func (s *GetStandardStatisticsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
