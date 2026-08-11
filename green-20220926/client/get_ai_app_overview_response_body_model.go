// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppTotal(v int64) *GetAiAppOverviewResponseBody
	GetAppTotal() *int64
	SetRequestId(v string) *GetAiAppOverviewResponseBody
	GetRequestId() *string
	SetRiskEventLevelDistribution(v map[string]interface{}) *GetAiAppOverviewResponseBody
	GetRiskEventLevelDistribution() map[string]interface{}
	SetRiskEventResolvedTotal(v int64) *GetAiAppOverviewResponseBody
	GetRiskEventResolvedTotal() *int64
	SetRiskEventTotal(v int64) *GetAiAppOverviewResponseBody
	GetRiskEventTotal() *int64
	SetRiskEventUnhandledTotal(v int64) *GetAiAppOverviewResponseBody
	GetRiskEventUnhandledTotal() *int64
}

type GetAiAppOverviewResponseBody struct {
	// The total number of agents.
	//
	// example:
	//
	// 100
	AppTotal *int64 `json:"AppTotal,omitempty" xml:"AppTotal,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The distribution of risk events by level.
	RiskEventLevelDistribution map[string]interface{} `json:"RiskEventLevelDistribution,omitempty" xml:"RiskEventLevelDistribution,omitempty"`
	// The total number of resolved risk events.
	//
	// example:
	//
	// 10
	RiskEventResolvedTotal *int64 `json:"RiskEventResolvedTotal,omitempty" xml:"RiskEventResolvedTotal,omitempty"`
	// The total number of risk events.
	//
	// example:
	//
	// 20
	RiskEventTotal *int64 `json:"RiskEventTotal,omitempty" xml:"RiskEventTotal,omitempty"`
	// The total number of unhandled risk events.
	//
	// example:
	//
	// 10
	RiskEventUnhandledTotal *int64 `json:"RiskEventUnhandledTotal,omitempty" xml:"RiskEventUnhandledTotal,omitempty"`
}

func (s GetAiAppOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiAppOverviewResponseBody) GetAppTotal() *int64 {
	return s.AppTotal
}

func (s *GetAiAppOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiAppOverviewResponseBody) GetRiskEventLevelDistribution() map[string]interface{} {
	return s.RiskEventLevelDistribution
}

func (s *GetAiAppOverviewResponseBody) GetRiskEventResolvedTotal() *int64 {
	return s.RiskEventResolvedTotal
}

func (s *GetAiAppOverviewResponseBody) GetRiskEventTotal() *int64 {
	return s.RiskEventTotal
}

func (s *GetAiAppOverviewResponseBody) GetRiskEventUnhandledTotal() *int64 {
	return s.RiskEventUnhandledTotal
}

func (s *GetAiAppOverviewResponseBody) SetAppTotal(v int64) *GetAiAppOverviewResponseBody {
	s.AppTotal = &v
	return s
}

func (s *GetAiAppOverviewResponseBody) SetRequestId(v string) *GetAiAppOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiAppOverviewResponseBody) SetRiskEventLevelDistribution(v map[string]interface{}) *GetAiAppOverviewResponseBody {
	s.RiskEventLevelDistribution = v
	return s
}

func (s *GetAiAppOverviewResponseBody) SetRiskEventResolvedTotal(v int64) *GetAiAppOverviewResponseBody {
	s.RiskEventResolvedTotal = &v
	return s
}

func (s *GetAiAppOverviewResponseBody) SetRiskEventTotal(v int64) *GetAiAppOverviewResponseBody {
	s.RiskEventTotal = &v
	return s
}

func (s *GetAiAppOverviewResponseBody) SetRiskEventUnhandledTotal(v int64) *GetAiAppOverviewResponseBody {
	s.RiskEventUnhandledTotal = &v
	return s
}

func (s *GetAiAppOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}
