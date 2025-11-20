// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoByRoomCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConferenceList(v []*QueryConferenceInfoByRoomCodeResponseBodyConferenceList) *QueryConferenceInfoByRoomCodeResponseBody
	GetConferenceList() []*QueryConferenceInfoByRoomCodeResponseBodyConferenceList
	SetDingtalkRequestId(v string) *QueryConferenceInfoByRoomCodeResponseBody
	GetDingtalkRequestId() *string
	SetHasMore(v bool) *QueryConferenceInfoByRoomCodeResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *QueryConferenceInfoByRoomCodeResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryConferenceInfoByRoomCodeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryConferenceInfoByRoomCodeResponseBody
	GetTotalCount() *int32
	SetVendorRequestId(v string) *QueryConferenceInfoByRoomCodeResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryConferenceInfoByRoomCodeResponseBody
	GetVendorType() *string
}

type QueryConferenceInfoByRoomCodeResponseBody struct {
	// example:
	//
	// []
	ConferenceList []*QueryConferenceInfoByRoomCodeResponseBodyConferenceList `json:"conferenceList,omitempty" xml:"conferenceList,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	DingtalkRequestId *string `json:"dingtalkRequestId,omitempty" xml:"dingtalkRequestId,omitempty"`
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 1296
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetConferenceList() []*QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	return s.ConferenceList
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetDingtalkRequestId() *string {
	return s.DingtalkRequestId
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetConferenceList(v []*QueryConferenceInfoByRoomCodeResponseBodyConferenceList) *QueryConferenceInfoByRoomCodeResponseBody {
	s.ConferenceList = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetDingtalkRequestId(v string) *QueryConferenceInfoByRoomCodeResponseBody {
	s.DingtalkRequestId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetHasMore(v bool) *QueryConferenceInfoByRoomCodeResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetNextToken(v string) *QueryConferenceInfoByRoomCodeResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetRequestId(v string) *QueryConferenceInfoByRoomCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetTotalCount(v int32) *QueryConferenceInfoByRoomCodeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetVendorRequestId(v string) *QueryConferenceInfoByRoomCodeResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetVendorType(v string) *QueryConferenceInfoByRoomCodeResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) Validate() error {
	if s.ConferenceList != nil {
		for _, item := range s.ConferenceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryConferenceInfoByRoomCodeResponseBodyConferenceList struct {
	// example:
	//
	// 3
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1000
	ConfDuration *int64 `json:"ConfDuration,omitempty" xml:"ConfDuration,omitempty"`
	// example:
	//
	// 636cf59f2b032f014ae32902
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// example:
	//
	// 527079
	CreatorId   *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorNick *string `json:"CreatorNick,omitempty" xml:"CreatorNick,omitempty"`
	// example:
	//
	// 1668087732000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// https//:xxx
	ExternalLinkUrl *string `json:"ExternalLinkUrl,omitempty" xml:"ExternalLinkUrl,omitempty"`
	// example:
	//
	// 123456789
	RoomCode *string `json:"RoomCode,omitempty" xml:"RoomCode,omitempty"`
	// example:
	//
	// 2d79cbde-b9d8-4256-9788-78b05834944e
	ScheduleConferenceId *string `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	// example:
	//
	// 1668087731000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeResponseBodyConferenceList) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetBizType() *string {
	return s.BizType
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetConfDuration() *int64 {
	return s.ConfDuration
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetCreatorNick() *string {
	return s.CreatorNick
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetExternalLinkUrl() *string {
	return s.ExternalLinkUrl
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetRoomCode() *string {
	return s.RoomCode
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GetTitle() *string {
	return s.Title
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetBizType(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.BizType = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetConfDuration(v int64) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ConfDuration = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetConferenceId(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetCreatorId(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.CreatorId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetCreatorNick(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.CreatorNick = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetEndTime(v int64) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.EndTime = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetExternalLinkUrl(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ExternalLinkUrl = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetRoomCode(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetScheduleConferenceId(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetStartTime(v int64) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.StartTime = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetStatus(v int32) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.Status = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetTitle(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.Title = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) Validate() error {
	return dara.Validate(s)
}
