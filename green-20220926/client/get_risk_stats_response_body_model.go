// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRiskStatsResponseBody
	GetRequestId() *string
	SetRiskStats(v []*GetRiskStatsResponseBodyRiskStats) *GetRiskStatsResponseBody
	GetRiskStats() []*GetRiskStatsResponseBodyRiskStats
}

type GetRiskStatsResponseBody struct {
	// The ID assigned by the backend to uniquely identify a request. This ID can be used to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of risk posture statistics.
	RiskStats []*GetRiskStatsResponseBodyRiskStats `json:"RiskStats,omitempty" xml:"RiskStats,omitempty" type:"Repeated"`
}

func (s GetRiskStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRiskStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRiskStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRiskStatsResponseBody) GetRiskStats() []*GetRiskStatsResponseBodyRiskStats {
	return s.RiskStats
}

func (s *GetRiskStatsResponseBody) SetRequestId(v string) *GetRiskStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRiskStatsResponseBody) SetRiskStats(v []*GetRiskStatsResponseBodyRiskStats) *GetRiskStatsResponseBody {
	s.RiskStats = v
	return s
}

func (s *GetRiskStatsResponseBody) Validate() error {
	if s.RiskStats != nil {
		for _, item := range s.RiskStats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRiskStatsResponseBodyRiskStats struct {
	// The total number of requests.
	//
	// example:
	//
	// 100
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The number of detected risks.
	//
	// example:
	//
	// 1
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The type.
	//
	// example:
	//
	// prompt_attack
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRiskStatsResponseBodyRiskStats) String() string {
	return dara.Prettify(s)
}

func (s GetRiskStatsResponseBodyRiskStats) GoString() string {
	return s.String()
}

func (s *GetRiskStatsResponseBodyRiskStats) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *GetRiskStatsResponseBodyRiskStats) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *GetRiskStatsResponseBodyRiskStats) GetType() *string {
	return s.Type
}

func (s *GetRiskStatsResponseBodyRiskStats) SetRequestCount(v int64) *GetRiskStatsResponseBodyRiskStats {
	s.RequestCount = &v
	return s
}

func (s *GetRiskStatsResponseBodyRiskStats) SetRiskCount(v int64) *GetRiskStatsResponseBodyRiskStats {
	s.RiskCount = &v
	return s
}

func (s *GetRiskStatsResponseBodyRiskStats) SetType(v string) *GetRiskStatsResponseBodyRiskStats {
	s.Type = &v
	return s
}

func (s *GetRiskStatsResponseBodyRiskStats) Validate() error {
	return dara.Validate(s)
}
