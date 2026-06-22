// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathEventStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAttackPathEventStatisticsResponseBody
	GetRequestId() *string
	SetSeriousPathRiskNum(v int64) *GetAttackPathEventStatisticsResponseBody
	GetSeriousPathRiskNum() *int64
	SetTotalAssetRiskNum(v int64) *GetAttackPathEventStatisticsResponseBody
	GetTotalAssetRiskNum() *int64
	SetTotalPathRiskNum(v int64) *GetAttackPathEventStatisticsResponseBody
	GetTotalPathRiskNum() *int64
}

type GetAttackPathEventStatisticsResponseBody struct {
	// The ID of the request. The China Chinese Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 89AD16CC-97EE-50F3-9B12-9E28E5C8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of attack paths that require immediate handling.
	//
	// example:
	//
	// 1
	SeriousPathRiskNum *int64 `json:"SeriousPathRiskNum,omitempty" xml:"SeriousPathRiskNum,omitempty"`
	// The number of risky assets.
	//
	// example:
	//
	// 5
	TotalAssetRiskNum *int64 `json:"TotalAssetRiskNum,omitempty" xml:"TotalAssetRiskNum,omitempty"`
	// The number of risky paths.
	//
	// example:
	//
	// 10
	TotalPathRiskNum *int64 `json:"TotalPathRiskNum,omitempty" xml:"TotalPathRiskNum,omitempty"`
}

func (s GetAttackPathEventStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathEventStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttackPathEventStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttackPathEventStatisticsResponseBody) GetSeriousPathRiskNum() *int64 {
	return s.SeriousPathRiskNum
}

func (s *GetAttackPathEventStatisticsResponseBody) GetTotalAssetRiskNum() *int64 {
	return s.TotalAssetRiskNum
}

func (s *GetAttackPathEventStatisticsResponseBody) GetTotalPathRiskNum() *int64 {
	return s.TotalPathRiskNum
}

func (s *GetAttackPathEventStatisticsResponseBody) SetRequestId(v string) *GetAttackPathEventStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttackPathEventStatisticsResponseBody) SetSeriousPathRiskNum(v int64) *GetAttackPathEventStatisticsResponseBody {
	s.SeriousPathRiskNum = &v
	return s
}

func (s *GetAttackPathEventStatisticsResponseBody) SetTotalAssetRiskNum(v int64) *GetAttackPathEventStatisticsResponseBody {
	s.TotalAssetRiskNum = &v
	return s
}

func (s *GetAttackPathEventStatisticsResponseBody) SetTotalPathRiskNum(v int64) *GetAttackPathEventStatisticsResponseBody {
	s.TotalPathRiskNum = &v
	return s
}

func (s *GetAttackPathEventStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
