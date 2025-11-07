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
	// The response parameters.
	Data *GetOssBucketScanStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FAC50208-E56B-5CC8-8738-2B219D1A****
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssBucketScanStatisticResponseBodyData struct {
	// The expiration time of the purchased quota.
	//
	// example:
	//
	// 1714442403000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The number of high-risk objects.
	//
	// example:
	//
	// 0
	HighRisk *int64 `json:"HighRisk,omitempty" xml:"HighRisk,omitempty"`
	// The number of low-risk objects.
	//
	// example:
	//
	// 0
	LowRisk *int64 `json:"LowRisk,omitempty" xml:"LowRisk,omitempty"`
	// The number of medium-risk objects.
	//
	// example:
	//
	// 0
	MediumRisk *int64 `json:"MediumRisk,omitempty" xml:"MediumRisk,omitempty"`
	// The number of buckets that are not checked.
	//
	// example:
	//
	// 1
	NoScanBucket *int32 `json:"NoScanBucket,omitempty" xml:"NoScanBucket,omitempty"`
	// Postpaid usage count.
	//
	// example:
	//
	// 1000
	PostPayInvokeCount *int64 `json:"PostPayInvokeCount,omitempty" xml:"PostPayInvokeCount,omitempty"`
	// Prepaid authorized count.
	//
	// example:
	//
	// 10000
	PrePayAuthCount *int64 `json:"PrePayAuthCount,omitempty" xml:"PrePayAuthCount,omitempty"`
	// Prepaid usage count.
	//
	// example:
	//
	// 100
	PrePayInvokeCount *int64 `json:"PrePayInvokeCount,omitempty" xml:"PrePayInvokeCount,omitempty"`
	// The remaining quota.
	//
	// example:
	//
	// 1
	RemainAuth *int32 `json:"RemainAuth,omitempty" xml:"RemainAuth,omitempty"`
	// The number of buckets in which at-risk objects exist.
	//
	// example:
	//
	// 1
	RiskBucket *int32 `json:"RiskBucket,omitempty" xml:"RiskBucket,omitempty"`
	// The number of objects that are checked.
	//
	// example:
	//
	// 1
	ScanObject *int64 `json:"ScanObject,omitempty" xml:"ScanObject,omitempty"`
	// The total number of buckets.
	//
	// example:
	//
	// 1
	TotalBucket *int32 `json:"TotalBucket,omitempty" xml:"TotalBucket,omitempty"`
	// The total number of objects in the bucket.
	//
	// example:
	//
	// 1
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

func (s *GetOssBucketScanStatisticResponseBodyData) GetPostPayInvokeCount() *int64 {
	return s.PostPayInvokeCount
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetPrePayAuthCount() *int64 {
	return s.PrePayAuthCount
}

func (s *GetOssBucketScanStatisticResponseBodyData) GetPrePayInvokeCount() *int64 {
	return s.PrePayInvokeCount
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

func (s *GetOssBucketScanStatisticResponseBodyData) SetPostPayInvokeCount(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.PostPayInvokeCount = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetPrePayAuthCount(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.PrePayAuthCount = &v
	return s
}

func (s *GetOssBucketScanStatisticResponseBodyData) SetPrePayInvokeCount(v int64) *GetOssBucketScanStatisticResponseBodyData {
	s.PrePayInvokeCount = &v
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
