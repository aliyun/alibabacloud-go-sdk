// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushStatByMsgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushStats(v *QueryPushStatByMsgResponseBodyPushStats) *QueryPushStatByMsgResponseBody
	GetPushStats() *QueryPushStatByMsgResponseBodyPushStats
	SetRequestId(v string) *QueryPushStatByMsgResponseBody
	GetRequestId() *string
}

type QueryPushStatByMsgResponseBody struct {
	PushStats *QueryPushStatByMsgResponseBodyPushStats `json:"PushStats,omitempty" xml:"PushStats,omitempty" type:"Struct"`
	// example:
	//
	// CF195C34-98FB-491A-98D7-19CBC1FA880B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPushStatByMsgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByMsgResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponseBody) GetPushStats() *QueryPushStatByMsgResponseBodyPushStats {
	return s.PushStats
}

func (s *QueryPushStatByMsgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPushStatByMsgResponseBody) SetPushStats(v *QueryPushStatByMsgResponseBodyPushStats) *QueryPushStatByMsgResponseBody {
	s.PushStats = v
	return s
}

func (s *QueryPushStatByMsgResponseBody) SetRequestId(v string) *QueryPushStatByMsgResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushStatByMsgResponseBody) Validate() error {
	if s.PushStats != nil {
		if err := s.PushStats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPushStatByMsgResponseBodyPushStats struct {
	PushStat []*QueryPushStatByMsgResponseBodyPushStatsPushStat `json:"PushStat,omitempty" xml:"PushStat,omitempty" type:"Repeated"`
}

func (s QueryPushStatByMsgResponseBodyPushStats) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByMsgResponseBodyPushStats) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponseBodyPushStats) GetPushStat() []*QueryPushStatByMsgResponseBodyPushStatsPushStat {
	return s.PushStat
}

func (s *QueryPushStatByMsgResponseBodyPushStats) SetPushStat(v []*QueryPushStatByMsgResponseBodyPushStatsPushStat) *QueryPushStatByMsgResponseBodyPushStats {
	s.PushStat = v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStats) Validate() error {
	if s.PushStat != nil {
		for _, item := range s.PushStat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPushStatByMsgResponseBodyPushStatsPushStat struct {
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
	// 510427
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
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
}

func (s QueryPushStatByMsgResponseBodyPushStatsPushStat) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByMsgResponseBodyPushStatsPushStat) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetAcceptCount() *int64 {
	return s.AcceptCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetDeletedCount() *int64 {
	return s.DeletedCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetOpenedCount() *int64 {
	return s.OpenedCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetReceivedCount() *int64 {
	return s.ReceivedCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetSentCount() *int64 {
	return s.SentCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetSmsFailedCount() *int64 {
	return s.SmsFailedCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetSmsReceiveFailedCount() *int64 {
	return s.SmsReceiveFailedCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetSmsReceiveSuccessCount() *int64 {
	return s.SmsReceiveSuccessCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetSmsSentCount() *int64 {
	return s.SmsSentCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) GetSmsSkipCount() *int64 {
	return s.SmsSkipCount
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetAcceptCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.AcceptCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetDeletedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.DeletedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetMessageId(v string) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.MessageId = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetOpenedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.OpenedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetReceivedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.ReceivedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSentCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SentCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsFailedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsFailedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsReceiveFailedCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsReceiveFailedCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsReceiveSuccessCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsReceiveSuccessCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsSentCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsSentCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) SetSmsSkipCount(v int64) *QueryPushStatByMsgResponseBodyPushStatsPushStat {
	s.SmsSkipCount = &v
	return s
}

func (s *QueryPushStatByMsgResponseBodyPushStatsPushStat) Validate() error {
	return dara.Validate(s)
}
