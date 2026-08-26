// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRtcMPUEventSubRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListRtcMPUEventSubRecordRequest
	GetAppId() *string
	SetEndTime(v string) *ListRtcMPUEventSubRecordRequest
	GetEndTime() *string
	SetPageNo(v int32) *ListRtcMPUEventSubRecordRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRtcMPUEventSubRecordRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListRtcMPUEventSubRecordRequest
	GetStartTime() *string
	SetSubId(v string) *ListRtcMPUEventSubRecordRequest
	GetSubId() *string
}

type ListRtcMPUEventSubRecordRequest struct {
	// The ID of the subscribed application. You can view your application IDs by navigating to **ApsaraVideo Live > Live+ > ApsaraVideo Real-time Communication > Application Management**.
	//
	// > - The application ID consists of uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters.
	//
	// > - You must first call CreateRtcMPUEventSub to create a stream mixing and forwarding event subscription for this application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query.
	//
	// Format: yyyy-MM-ddTHH:mm:ssZ (UTC). The value cannot be later than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1970-01-01T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query.
	//
	// Format: yyyy-MM-ddTHH:mm:ssZ (UTC). The value cannot be earlier than seven days before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1970-01-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The callback ID of the subscription.
	//
	// example:
	//
	// yourSubId
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
}

func (s ListRtcMPUEventSubRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUEventSubRecordRequest) GoString() string {
	return s.String()
}

func (s *ListRtcMPUEventSubRecordRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListRtcMPUEventSubRecordRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRtcMPUEventSubRecordRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRtcMPUEventSubRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRtcMPUEventSubRecordRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRtcMPUEventSubRecordRequest) GetSubId() *string {
	return s.SubId
}

func (s *ListRtcMPUEventSubRecordRequest) SetAppId(v string) *ListRtcMPUEventSubRecordRequest {
	s.AppId = &v
	return s
}

func (s *ListRtcMPUEventSubRecordRequest) SetEndTime(v string) *ListRtcMPUEventSubRecordRequest {
	s.EndTime = &v
	return s
}

func (s *ListRtcMPUEventSubRecordRequest) SetPageNo(v int32) *ListRtcMPUEventSubRecordRequest {
	s.PageNo = &v
	return s
}

func (s *ListRtcMPUEventSubRecordRequest) SetPageSize(v int32) *ListRtcMPUEventSubRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListRtcMPUEventSubRecordRequest) SetStartTime(v string) *ListRtcMPUEventSubRecordRequest {
	s.StartTime = &v
	return s
}

func (s *ListRtcMPUEventSubRecordRequest) SetSubId(v string) *ListRtcMPUEventSubRecordRequest {
	s.SubId = &v
	return s
}

func (s *ListRtcMPUEventSubRecordRequest) Validate() error {
	return dara.Validate(s)
}
