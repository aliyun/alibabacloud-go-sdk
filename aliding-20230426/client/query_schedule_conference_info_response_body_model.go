// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConferenceList(v []*QueryScheduleConferenceInfoResponseBodyConferenceList) *QueryScheduleConferenceInfoResponseBody
	GetConferenceList() []*QueryScheduleConferenceInfoResponseBodyConferenceList
	SetNextToken(v string) *QueryScheduleConferenceInfoResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryScheduleConferenceInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryScheduleConferenceInfoResponseBody
	GetTotalCount() *int32
	SetVendorRequestId(v string) *QueryScheduleConferenceInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryScheduleConferenceInfoResponseBody
	GetVendorType() *string
}

type QueryScheduleConferenceInfoResponseBody struct {
	ConferenceList []*QueryScheduleConferenceInfoResponseBodyConferenceList `json:"conferenceList,omitempty" xml:"conferenceList,omitempty" type:"Repeated"`
	// example:
	//
	// 19
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

func (s QueryScheduleConferenceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoResponseBody) GetConferenceList() []*QueryScheduleConferenceInfoResponseBodyConferenceList {
	return s.ConferenceList
}

func (s *QueryScheduleConferenceInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryScheduleConferenceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryScheduleConferenceInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryScheduleConferenceInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryScheduleConferenceInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryScheduleConferenceInfoResponseBody) SetConferenceList(v []*QueryScheduleConferenceInfoResponseBodyConferenceList) *QueryScheduleConferenceInfoResponseBody {
	s.ConferenceList = v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) SetNextToken(v string) *QueryScheduleConferenceInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) SetRequestId(v string) *QueryScheduleConferenceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) SetTotalCount(v int32) *QueryScheduleConferenceInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) SetVendorRequestId(v string) *QueryScheduleConferenceInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) SetVendorType(v string) *QueryScheduleConferenceInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) Validate() error {
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

type QueryScheduleConferenceInfoResponseBodyConferenceList struct {
	// example:
	//
	// 636cf59f2b032f014ae32902
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// example:
	//
	// 1668087732000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 123456789
	RoomCode *string `json:"RoomCode,omitempty" xml:"RoomCode,omitempty"`
	// example:
	//
	// 1668087731000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xxx发起的视频会议
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryScheduleConferenceInfoResponseBodyConferenceList) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoResponseBodyConferenceList) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) GetRoomCode() *string {
	return s.RoomCode
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) GetTitle() *string {
	return s.Title
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetConferenceId(v string) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.ConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetEndTime(v int64) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.EndTime = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetRoomCode(v string) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.RoomCode = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetStartTime(v int64) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.StartTime = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetStatus(v int32) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.Status = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetTitle(v string) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.Title = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) Validate() error {
	return dara.Validate(s)
}
