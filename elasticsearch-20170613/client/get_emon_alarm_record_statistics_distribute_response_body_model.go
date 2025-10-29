// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonAlarmRecordStatisticsDistributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBody
	GetCode() *string
	SetMessage(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBody
	GetRequestId() *string
	SetResult(v *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) *GetEmonAlarmRecordStatisticsDistributeResponseBody
	GetResult() *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult
	SetSuccess(v bool) *GetEmonAlarmRecordStatisticsDistributeResponseBody
	GetSuccess() *bool
}

type GetEmonAlarmRecordStatisticsDistributeResponseBody struct {
	Code      *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) GetResult() *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	return s.Result
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) SetCode(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) SetMessage(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) SetRequestId(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) SetResult(v *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) *GetEmonAlarmRecordStatisticsDistributeResponseBody {
	s.Result = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) SetSuccess(v bool) *GetEmonAlarmRecordStatisticsDistributeResponseBody {
	s.Success = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEmonAlarmRecordStatisticsDistributeResponseBodyResult struct {
	AlarmGroup      *string                                                                    `json:"alarmGroup,omitempty" xml:"alarmGroup,omitempty"`
	AlarmGroupTotal []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal `json:"alarmGroupTotal,omitempty" xml:"alarmGroupTotal,omitempty" type:"Repeated"`
	ChannelTotal    []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal    `json:"channelTotal,omitempty" xml:"channelTotal,omitempty" type:"Repeated"`
	Count           *int32                                                                     `json:"count,omitempty" xml:"count,omitempty"`
	LevelTotal      []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal      `json:"levelTotal,omitempty" xml:"levelTotal,omitempty" type:"Repeated"`
	ReceiverTotal   []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal   `json:"receiverTotal,omitempty" xml:"receiverTotal,omitempty" type:"Repeated"`
	Statistics      []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics      `json:"statistics,omitempty" xml:"statistics,omitempty" type:"Repeated"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GetAlarmGroup() *string {
	return s.AlarmGroup
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GetAlarmGroupTotal() []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	return s.AlarmGroupTotal
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GetChannelTotal() []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	return s.ChannelTotal
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GetCount() *int32 {
	return s.Count
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GetLevelTotal() []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	return s.LevelTotal
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GetReceiverTotal() []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	return s.ReceiverTotal
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) GetStatistics() []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	return s.Statistics
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) SetAlarmGroup(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	s.AlarmGroup = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) SetAlarmGroupTotal(v []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	s.AlarmGroupTotal = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) SetChannelTotal(v []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	s.ChannelTotal = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) SetCount(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	s.Count = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) SetLevelTotal(v []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	s.LevelTotal = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) SetReceiverTotal(v []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	s.ReceiverTotal = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) SetStatistics(v []*GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult {
	s.Statistics = v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResult) Validate() error {
	if s.AlarmGroupTotal != nil {
		for _, item := range s.AlarmGroupTotal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChannelTotal != nil {
		for _, item := range s.ChannelTotal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LevelTotal != nil {
		for _, item := range s.LevelTotal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReceiverTotal != nil {
		for _, item := range s.ReceiverTotal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal struct {
	AlarmGroup *string `json:"alarmGroup,omitempty" xml:"alarmGroup,omitempty"`
	Compare    *string `json:"compare,omitempty" xml:"compare,omitempty"`
	Count      *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
	Receiver   *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	Time       *int32  `json:"time,omitempty" xml:"time,omitempty"`
	Today      *int32  `json:"today,omitempty" xml:"today,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	Yesterday  *int32  `json:"yesterday,omitempty" xml:"yesterday,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetAlarmGroup() *string {
	return s.AlarmGroup
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetCompare() *string {
	return s.Compare
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetCount() *int32 {
	return s.Count
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetLevel() *string {
	return s.Level
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetReceiver() *string {
	return s.Receiver
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetTime() *int32 {
	return s.Time
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetToday() *int32 {
	return s.Today
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetType() *string {
	return s.Type
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) GetYesterday() *int32 {
	return s.Yesterday
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetAlarmGroup(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.AlarmGroup = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetCompare(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Compare = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetCount(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Count = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetLevel(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Level = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetReceiver(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Receiver = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetTime(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Time = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetToday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Today = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetType(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Type = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) SetYesterday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal {
	s.Yesterday = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultAlarmGroupTotal) Validate() error {
	return dara.Validate(s)
}

type GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal struct {
	AlarmGroup *string `json:"alarmGroup,omitempty" xml:"alarmGroup,omitempty"`
	Compare    *string `json:"compare,omitempty" xml:"compare,omitempty"`
	Count      *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
	Receiver   *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	Time       *int32  `json:"time,omitempty" xml:"time,omitempty"`
	Today      *int32  `json:"today,omitempty" xml:"today,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	Yesterday  *int32  `json:"yesterday,omitempty" xml:"yesterday,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetAlarmGroup() *string {
	return s.AlarmGroup
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetCompare() *string {
	return s.Compare
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetCount() *int32 {
	return s.Count
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetLevel() *string {
	return s.Level
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetReceiver() *string {
	return s.Receiver
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetTime() *int32 {
	return s.Time
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetToday() *int32 {
	return s.Today
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetType() *string {
	return s.Type
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) GetYesterday() *int32 {
	return s.Yesterday
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetAlarmGroup(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.AlarmGroup = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetCompare(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Compare = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetCount(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Count = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetLevel(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Level = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetReceiver(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Receiver = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetTime(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Time = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetToday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Today = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetType(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Type = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) SetYesterday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal {
	s.Yesterday = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultChannelTotal) Validate() error {
	return dara.Validate(s)
}

type GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal struct {
	AlarmGroup *string `json:"alarmGroup,omitempty" xml:"alarmGroup,omitempty"`
	Compare    *string `json:"compare,omitempty" xml:"compare,omitempty"`
	Count      *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
	Receiver   *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	Time       *int32  `json:"time,omitempty" xml:"time,omitempty"`
	Today      *int32  `json:"today,omitempty" xml:"today,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	Yesterday  *int32  `json:"yesterday,omitempty" xml:"yesterday,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetAlarmGroup() *string {
	return s.AlarmGroup
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetCompare() *string {
	return s.Compare
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetCount() *int32 {
	return s.Count
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetLevel() *string {
	return s.Level
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetReceiver() *string {
	return s.Receiver
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetTime() *int32 {
	return s.Time
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetToday() *int32 {
	return s.Today
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetType() *string {
	return s.Type
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) GetYesterday() *int32 {
	return s.Yesterday
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetAlarmGroup(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.AlarmGroup = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetCompare(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Compare = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetCount(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Count = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetLevel(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Level = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetReceiver(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Receiver = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetTime(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Time = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetToday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Today = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetType(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Type = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) SetYesterday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal {
	s.Yesterday = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultLevelTotal) Validate() error {
	return dara.Validate(s)
}

type GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal struct {
	AlarmGroup *string `json:"alarmGroup,omitempty" xml:"alarmGroup,omitempty"`
	Compare    *string `json:"compare,omitempty" xml:"compare,omitempty"`
	Count      *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
	Receiver   *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	Time       *int32  `json:"time,omitempty" xml:"time,omitempty"`
	Today      *int32  `json:"today,omitempty" xml:"today,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	Yesterday  *int32  `json:"yesterday,omitempty" xml:"yesterday,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetAlarmGroup() *string {
	return s.AlarmGroup
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetCompare() *string {
	return s.Compare
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetCount() *int32 {
	return s.Count
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetLevel() *string {
	return s.Level
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetReceiver() *string {
	return s.Receiver
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetTime() *int32 {
	return s.Time
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetToday() *int32 {
	return s.Today
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetType() *string {
	return s.Type
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) GetYesterday() *int32 {
	return s.Yesterday
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetAlarmGroup(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.AlarmGroup = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetCompare(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Compare = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetCount(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Count = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetLevel(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Level = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetReceiver(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Receiver = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetTime(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Time = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetToday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Today = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetType(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Type = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) SetYesterday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal {
	s.Yesterday = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultReceiverTotal) Validate() error {
	return dara.Validate(s)
}

type GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics struct {
	AlarmGroup *string `json:"alarmGroup,omitempty" xml:"alarmGroup,omitempty"`
	Compare    *string `json:"compare,omitempty" xml:"compare,omitempty"`
	Count      *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
	Receiver   *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	Time       *int32  `json:"time,omitempty" xml:"time,omitempty"`
	Today      *int32  `json:"today,omitempty" xml:"today,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	Yesterday  *int32  `json:"yesterday,omitempty" xml:"yesterday,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetAlarmGroup() *string {
	return s.AlarmGroup
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetCompare() *string {
	return s.Compare
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetCount() *int32 {
	return s.Count
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetLevel() *string {
	return s.Level
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetReceiver() *string {
	return s.Receiver
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetTime() *int32 {
	return s.Time
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetToday() *int32 {
	return s.Today
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetType() *string {
	return s.Type
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) GetYesterday() *int32 {
	return s.Yesterday
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetAlarmGroup(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.AlarmGroup = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetCompare(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Compare = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetCount(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Count = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetLevel(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Level = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetReceiver(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Receiver = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetTime(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Time = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetToday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Today = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetType(v string) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Type = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) SetYesterday(v int32) *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics {
	s.Yesterday = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeResponseBodyResultStatistics) Validate() error {
	return dara.Validate(s)
}
