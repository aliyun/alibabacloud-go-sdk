// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspiciousStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemindCount(v int32) *GetSuspiciousStatisticsResponseBody
	GetRemindCount() *int32
	SetRequestId(v string) *GetSuspiciousStatisticsResponseBody
	GetRequestId() *string
	SetSeriousCount(v int32) *GetSuspiciousStatisticsResponseBody
	GetSeriousCount() *int32
	SetSuspiciousCount(v int32) *GetSuspiciousStatisticsResponseBody
	GetSuspiciousCount() *int32
	SetTotalCount(v int32) *GetSuspiciousStatisticsResponseBody
	GetTotalCount() *int32
}

type GetSuspiciousStatisticsResponseBody struct {
	// The number of alerts whose Emergency level is Reminder.
	//
	// example:
	//
	// 0
	RemindCount *int32 `json:"RemindCount,omitempty" xml:"RemindCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 84092B42-1A59-4F34-8DF8-1D93520990A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of alerts whose Emergency level is Urgent.
	//
	// example:
	//
	// 1
	SeriousCount *int32 `json:"SeriousCount,omitempty" xml:"SeriousCount,omitempty"`
	// The number of alerts whose Emergency level is Suspicious.
	//
	// example:
	//
	// 8
	SuspiciousCount *int32 `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
	// The total number of alerts.
	//
	// example:
	//
	// 9
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSuspiciousStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSuspiciousStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspiciousStatisticsResponseBody) GetRemindCount() *int32 {
	return s.RemindCount
}

func (s *GetSuspiciousStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSuspiciousStatisticsResponseBody) GetSeriousCount() *int32 {
	return s.SeriousCount
}

func (s *GetSuspiciousStatisticsResponseBody) GetSuspiciousCount() *int32 {
	return s.SuspiciousCount
}

func (s *GetSuspiciousStatisticsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetSuspiciousStatisticsResponseBody) SetRemindCount(v int32) *GetSuspiciousStatisticsResponseBody {
	s.RemindCount = &v
	return s
}

func (s *GetSuspiciousStatisticsResponseBody) SetRequestId(v string) *GetSuspiciousStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspiciousStatisticsResponseBody) SetSeriousCount(v int32) *GetSuspiciousStatisticsResponseBody {
	s.SeriousCount = &v
	return s
}

func (s *GetSuspiciousStatisticsResponseBody) SetSuspiciousCount(v int32) *GetSuspiciousStatisticsResponseBody {
	s.SuspiciousCount = &v
	return s
}

func (s *GetSuspiciousStatisticsResponseBody) SetTotalCount(v int32) *GetSuspiciousStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSuspiciousStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
