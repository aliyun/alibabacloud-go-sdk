// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonAlarmRecordStatisticsDistributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetEmonAlarmRecordStatisticsDistributeRequest
	GetBody() *string
	SetGroupId(v string) *GetEmonAlarmRecordStatisticsDistributeRequest
	GetGroupId() *string
	SetTimeEnd(v int64) *GetEmonAlarmRecordStatisticsDistributeRequest
	GetTimeEnd() *int64
	SetTimeStart(v int64) *GetEmonAlarmRecordStatisticsDistributeRequest
	GetTimeStart() *int64
}

type GetEmonAlarmRecordStatisticsDistributeRequest struct {
	Body      *string `json:"body,omitempty" xml:"body,omitempty"`
	GroupId   *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	TimeEnd   *int64  `json:"timeEnd,omitempty" xml:"timeEnd,omitempty"`
	TimeStart *int64  `json:"timeStart,omitempty" xml:"timeStart,omitempty"`
}

func (s GetEmonAlarmRecordStatisticsDistributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmonAlarmRecordStatisticsDistributeRequest) GoString() string {
	return s.String()
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) GetBody() *string {
	return s.Body
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) GetTimeEnd() *int64 {
	return s.TimeEnd
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) GetTimeStart() *int64 {
	return s.TimeStart
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) SetBody(v string) *GetEmonAlarmRecordStatisticsDistributeRequest {
	s.Body = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) SetGroupId(v string) *GetEmonAlarmRecordStatisticsDistributeRequest {
	s.GroupId = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) SetTimeEnd(v int64) *GetEmonAlarmRecordStatisticsDistributeRequest {
	s.TimeEnd = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) SetTimeStart(v int64) *GetEmonAlarmRecordStatisticsDistributeRequest {
	s.TimeStart = &v
	return s
}

func (s *GetEmonAlarmRecordStatisticsDistributeRequest) Validate() error {
	return dara.Validate(s)
}
