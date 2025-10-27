// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushStatByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppPushStats(v *QueryPushStatByAppResponseBodyAppPushStats) *QueryPushStatByAppResponseBody
	GetAppPushStats() *QueryPushStatByAppResponseBodyAppPushStats
	SetRequestId(v string) *QueryPushStatByAppResponseBody
	GetRequestId() *string
}

type QueryPushStatByAppResponseBody struct {
	AppPushStats *QueryPushStatByAppResponseBodyAppPushStats `json:"AppPushStats,omitempty" xml:"AppPushStats,omitempty" type:"Struct"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPushStatByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByAppResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponseBody) GetAppPushStats() *QueryPushStatByAppResponseBodyAppPushStats {
	return s.AppPushStats
}

func (s *QueryPushStatByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPushStatByAppResponseBody) SetAppPushStats(v *QueryPushStatByAppResponseBodyAppPushStats) *QueryPushStatByAppResponseBody {
	s.AppPushStats = v
	return s
}

func (s *QueryPushStatByAppResponseBody) SetRequestId(v string) *QueryPushStatByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushStatByAppResponseBody) Validate() error {
	if s.AppPushStats != nil {
		if err := s.AppPushStats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPushStatByAppResponseBodyAppPushStats struct {
	AppPushStat []*QueryPushStatByAppResponseBodyAppPushStatsAppPushStat `json:"AppPushStat,omitempty" xml:"AppPushStat,omitempty" type:"Repeated"`
}

func (s QueryPushStatByAppResponseBodyAppPushStats) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByAppResponseBodyAppPushStats) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponseBodyAppPushStats) GetAppPushStat() []*QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	return s.AppPushStat
}

func (s *QueryPushStatByAppResponseBodyAppPushStats) SetAppPushStat(v []*QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) *QueryPushStatByAppResponseBodyAppPushStats {
	s.AppPushStat = v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStats) Validate() error {
	if s.AppPushStat != nil {
		for _, item := range s.AppPushStat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPushStatByAppResponseBodyAppPushStatsAppPushStat struct {
	// example:
	//
	// 120
	AcceptCount *int64 `json:"AcceptCount,omitempty" xml:"AcceptCount,omitempty"`
	// example:
	//
	// 10
	DeletedCount *int64 `json:"DeletedCount,omitempty" xml:"DeletedCount,omitempty"`
	// example:
	//
	// 30
	OpenedCount *int64 `json:"OpenedCount,omitempty" xml:"OpenedCount,omitempty"`
	// example:
	//
	// 60
	ReceivedCount *int64 `json:"ReceivedCount,omitempty" xml:"ReceivedCount,omitempty"`
	// example:
	//
	// 100
	SentCount *int64 `json:"SentCount,omitempty" xml:"SentCount,omitempty"`
	// example:
	//
	// 0
	SmsFailedCount *int64 `json:"SmsFailedCount,omitempty" xml:"SmsFailedCount,omitempty"`
	// example:
	//
	// 0
	SmsReceiveFailedCount *int64 `json:"SmsReceiveFailedCount,omitempty" xml:"SmsReceiveFailedCount,omitempty"`
	// example:
	//
	// 0
	SmsReceiveSuccessCount *int64 `json:"SmsReceiveSuccessCount,omitempty" xml:"SmsReceiveSuccessCount,omitempty"`
	// example:
	//
	// 0
	SmsSentCount *int64 `json:"SmsSentCount,omitempty" xml:"SmsSentCount,omitempty"`
	// example:
	//
	// 0
	SmsSkipCount *int64 `json:"SmsSkipCount,omitempty" xml:"SmsSkipCount,omitempty"`
	// example:
	//
	// 2016-07-25T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GoString() string {
	return s.String()
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetAcceptCount() *int64 {
	return s.AcceptCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetDeletedCount() *int64 {
	return s.DeletedCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetOpenedCount() *int64 {
	return s.OpenedCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetReceivedCount() *int64 {
	return s.ReceivedCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetSentCount() *int64 {
	return s.SentCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetSmsFailedCount() *int64 {
	return s.SmsFailedCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetSmsReceiveFailedCount() *int64 {
	return s.SmsReceiveFailedCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetSmsReceiveSuccessCount() *int64 {
	return s.SmsReceiveSuccessCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetSmsSentCount() *int64 {
	return s.SmsSentCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetSmsSkipCount() *int64 {
	return s.SmsSkipCount
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) GetTime() *string {
	return s.Time
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetAcceptCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.AcceptCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetDeletedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.DeletedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetOpenedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.OpenedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetReceivedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.ReceivedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSentCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SentCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsFailedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsFailedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsReceiveFailedCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsReceiveFailedCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsReceiveSuccessCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsReceiveSuccessCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsSentCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsSentCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetSmsSkipCount(v int64) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.SmsSkipCount = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) SetTime(v string) *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat {
	s.Time = &v
	return s
}

func (s *QueryPushStatByAppResponseBodyAppPushStatsAppPushStat) Validate() error {
	return dara.Validate(s)
}
