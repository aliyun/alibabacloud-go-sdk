// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerhubImageRiskStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDockerhubImageRiskStatisticResponseBody
	GetRequestId() *string
	SetRiskRankInfo(v *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) *GetDockerhubImageRiskStatisticResponseBody
	GetRiskRankInfo() *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo
}

type GetDockerhubImageRiskStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the risk source.
	RiskRankInfo *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo `json:"RiskRankInfo,omitempty" xml:"RiskRankInfo,omitempty" type:"Struct"`
}

func (s GetDockerhubImageRiskStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDockerhubImageRiskStatisticResponseBody) GetRiskRankInfo() *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo {
	return s.RiskRankInfo
}

func (s *GetDockerhubImageRiskStatisticResponseBody) SetRequestId(v string) *GetDockerhubImageRiskStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponseBody) SetRiskRankInfo(v *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) *GetDockerhubImageRiskStatisticResponseBody {
	s.RiskRankInfo = v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponseBody) Validate() error {
	if s.RiskRankInfo != nil {
		if err := s.RiskRankInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo struct {
	// The number of baseline risks.
	//
	// example:
	//
	// 1
	Baseline *int32 `json:"Baseline,omitempty" xml:"Baseline,omitempty"`
	// The timestamp when the scan was performed.
	//
	// example:
	//
	// 1693997625000
	ScanTime *int64 `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	// The timestamp when the scan was performed.
	//
	// example:
	//
	// 1693997625000
	ScanTimeTimestamp *int64 `json:"ScanTimeTimestamp,omitempty" xml:"ScanTimeTimestamp,omitempty"`
	// The number of scanned Docker Hub images.
	//
	// example:
	//
	// 1
	TotalScanned *int32 `json:"TotalScanned,omitempty" xml:"TotalScanned,omitempty"`
	// The number of high-risk vulnerabilities.
	//
	// example:
	//
	// 1
	VulAsap *int32 `json:"VulAsap,omitempty" xml:"VulAsap,omitempty"`
}

func (s GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) GoString() string {
	return s.String()
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) GetBaseline() *int32 {
	return s.Baseline
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) GetScanTime() *int64 {
	return s.ScanTime
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) GetScanTimeTimestamp() *int64 {
	return s.ScanTimeTimestamp
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) GetTotalScanned() *int32 {
	return s.TotalScanned
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) GetVulAsap() *int32 {
	return s.VulAsap
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) SetBaseline(v int32) *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo {
	s.Baseline = &v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) SetScanTime(v int64) *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo {
	s.ScanTime = &v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) SetScanTimeTimestamp(v int64) *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo {
	s.ScanTimeTimestamp = &v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) SetTotalScanned(v int32) *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo {
	s.TotalScanned = &v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) SetVulAsap(v int32) *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo {
	s.VulAsap = &v
	return s
}

func (s *GetDockerhubImageRiskStatisticResponseBodyRiskRankInfo) Validate() error {
	return dara.Validate(s)
}
