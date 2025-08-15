// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssBucketScanStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetOssBucketScanStatisticResponseBodyData) *GetOssBucketScanStatisticResponseBody
	GetData() *GetOssBucketScanStatisticResponseBodyData
	SetRequestId(v string) *GetOssBucketScanStatisticResponseBody
	GetRequestId() *string
}

type GetOssBucketScanStatisticResponseBody struct {
	Data *GetOssBucketScanStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E14CECBF-8CAC-520C-ACC3-9503D5******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssBucketScanStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssBucketScanStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssBucketScanStatisticResponseBody) GetData() *GetOssBucketScanStatisticResponseBodyData {
	return s.Data
}

func (s *GetOssBucketScanStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssBucketScanStatisticResponseBody) SetData(v *GetOssBucketScanStatisticResponseBodyData) *GetOssBucketScanStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetOssBucketScanStatisticResponseBody) SetRequestId(v string) *GetOssBucketScanStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOssBucketScanStatisticResponseBodyData struct {
	// example:
	//
	// 1605905716750024
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1
	HighRisk *int64 `json:"HighRisk,omitempty" xml:"HighRisk,omitempty"`
	// example:
	//
	// 1
	LowRisk *int64 `json:"LowRisk,omitempty" xml:"LowRisk,omitempty"`
	// example:
	//
	// 1
	MediumRisk *int64 `json:"MediumRisk,omitempty" xml:"MediumRisk,omitempty"`
	// example:
	//
	// 1
	NoScanBucket *int32 `json:"NoScanBucket,omitempty" xml:"NoScanBucket,omitempty"`
	// example:
	//
	// 1
	RemainAuth *int32 `json:"RemainAuth,omitempty" xml:"RemainAuth,omitempty"`
	// example:
	//
	// 1
	RiskBucket *int32 `json:"RiskBucket,omitempty" xml:"RiskBucket,omitempty"`
	// example:
	//
	// 10
	ScanObject *int64 `json:"ScanObject,omitempty" xml:"ScanObject,omitempty"`
	// example:
	//
	// 10
	TotalBucket *int32 `json:"TotalBucket,omitempty" xml:"TotalBucket,omitempty"`
	// example:
	//
	// 111
	TotalObject *int64 `json:"TotalObject,omitempty" xml:"TotalObject,omitempty"`
}

func (s GetOssBucketScanStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOssBucketScanStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetHighRisk() *int64 {
	return s.HighRisk
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetLowRisk() *int64 {
	return s.LowRisk
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetMediumRisk() *int64 {
	return s.MediumRisk
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetNoScanBucket() *int32 {
	return s.NoScanBucket
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetRemainAuth() *int32 {
	return s.RemainAuth
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetRiskBucket() *int32 {
	return s.RiskBucket
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetScanObject() *int64 {
	return s.ScanObject
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetTotalBucket() *int32 {
	return s.TotalBucket
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetTotalObject() *int64 {
	return s.TotalObject
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetExpireTime(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetHighRisk(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.HighRisk = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetLowRisk(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.LowRisk = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetMediumRisk(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.MediumRisk = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetNoScanBucket(v int32) *GetOssBucketScanStatisticResponseBodyData {
	s.NoScanBucket = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetRemainAuth(v int32) *GetOssBucketScanStatisticResponseBodyData {
	s.RemainAuth = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetRiskBucket(v int32) *GetOssBucketScanStatisticResponseBodyData {
	s.RiskBucket = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetScanObject(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.ScanObject = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetTotalBucket(v int32) *GetOssBucketScanStatisticResponseBodyData {
	s.TotalBucket = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetTotalObject(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.TotalObject = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
