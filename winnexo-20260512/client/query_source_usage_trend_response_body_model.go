// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySourceUsageTrendResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySourceUsageTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySourceUsageTrendResponseBody
	GetRequestId() *string
	SetTrend(v []*QuerySourceUsageTrendResponseBodyTrend) *QuerySourceUsageTrendResponseBody
	GetTrend() []*QuerySourceUsageTrendResponseBodyTrend
}

type QuerySourceUsageTrendResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error details.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The trend details.
	//
	// example:
	//
	// [{"bucketStart":"2026-11-01T00:00:00Z","retrievalCount":12,"usageCount":5}]
	Trend []*QuerySourceUsageTrendResponseBodyTrend `json:"trend,omitempty" xml:"trend,omitempty" type:"Repeated"`
}

func (s QuerySourceUsageTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageTrendResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySourceUsageTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySourceUsageTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySourceUsageTrendResponseBody) GetTrend() []*QuerySourceUsageTrendResponseBodyTrend {
	return s.Trend
}

func (s *QuerySourceUsageTrendResponseBody) SetCode(v string) *QuerySourceUsageTrendResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBody) SetMessage(v string) *QuerySourceUsageTrendResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBody) SetRequestId(v string) *QuerySourceUsageTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBody) SetTrend(v []*QuerySourceUsageTrendResponseBodyTrend) *QuerySourceUsageTrendResponseBody {
	s.Trend = v
	return s
}

func (s *QuerySourceUsageTrendResponseBody) Validate() error {
	if s.Trend != nil {
		for _, item := range s.Trend {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySourceUsageTrendResponseBodyTrend struct {
	// The start time of the natural bucket with the selected time zone. The first bucket may start earlier than the startTime.
	//
	// example:
	//
	// Start time of the natural bucket with the selected time zone. The first bucket may start earlier than the startTime
	BucketStart *string `json:"bucketStart,omitempty" xml:"bucketStart,omitempty"`
	// The number of retrieval recalls.
	//
	// example:
	//
	// 1
	RetrievalCount *int64 `json:"retrievalCount,omitempty" xml:"retrievalCount,omitempty"`
	// The deduplicated number of sources with retrieval events.
	//
	// example:
	//
	// 5
	RetrievedSourceCount *int64 `json:"retrievedSourceCount,omitempty" xml:"retrievedSourceCount,omitempty"`
	// The number of rows in the source data.
	//
	// example:
	//
	// 4
	SourceCount *int64 `json:"sourceCount,omitempty" xml:"sourceCount,omitempty"`
	// The number of times used by the agent.
	//
	// example:
	//
	// 2
	UsageCount *int64 `json:"usageCount,omitempty" xml:"usageCount,omitempty"`
	// The deduplicated number of sources with usage events.
	//
	// example:
	//
	// 1
	UsedSourceCount *int64 `json:"usedSourceCount,omitempty" xml:"usedSourceCount,omitempty"`
}

func (s QuerySourceUsageTrendResponseBodyTrend) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageTrendResponseBodyTrend) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageTrendResponseBodyTrend) GetBucketStart() *string {
	return s.BucketStart
}

func (s *QuerySourceUsageTrendResponseBodyTrend) GetRetrievalCount() *int64 {
	return s.RetrievalCount
}

func (s *QuerySourceUsageTrendResponseBodyTrend) GetRetrievedSourceCount() *int64 {
	return s.RetrievedSourceCount
}

func (s *QuerySourceUsageTrendResponseBodyTrend) GetSourceCount() *int64 {
	return s.SourceCount
}

func (s *QuerySourceUsageTrendResponseBodyTrend) GetUsageCount() *int64 {
	return s.UsageCount
}

func (s *QuerySourceUsageTrendResponseBodyTrend) GetUsedSourceCount() *int64 {
	return s.UsedSourceCount
}

func (s *QuerySourceUsageTrendResponseBodyTrend) SetBucketStart(v string) *QuerySourceUsageTrendResponseBodyTrend {
	s.BucketStart = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBodyTrend) SetRetrievalCount(v int64) *QuerySourceUsageTrendResponseBodyTrend {
	s.RetrievalCount = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBodyTrend) SetRetrievedSourceCount(v int64) *QuerySourceUsageTrendResponseBodyTrend {
	s.RetrievedSourceCount = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBodyTrend) SetSourceCount(v int64) *QuerySourceUsageTrendResponseBodyTrend {
	s.SourceCount = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBodyTrend) SetUsageCount(v int64) *QuerySourceUsageTrendResponseBodyTrend {
	s.UsageCount = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBodyTrend) SetUsedSourceCount(v int64) *QuerySourceUsageTrendResponseBodyTrend {
	s.UsedSourceCount = &v
	return s
}

func (s *QuerySourceUsageTrendResponseBodyTrend) Validate() error {
	return dara.Validate(s)
}
