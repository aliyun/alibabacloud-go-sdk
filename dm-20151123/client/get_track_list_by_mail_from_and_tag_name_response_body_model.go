// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrackListByMailFromAndTagNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOffsetCreateTime(v string) *GetTrackListByMailFromAndTagNameResponseBody
	GetOffsetCreateTime() *string
	SetOffsetCreateTimeDesc(v string) *GetTrackListByMailFromAndTagNameResponseBody
	GetOffsetCreateTimeDesc() *string
	SetPageNo(v int32) *GetTrackListByMailFromAndTagNameResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetTrackListByMailFromAndTagNameResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetTrackListByMailFromAndTagNameResponseBody
	GetRequestId() *string
	SetTotal(v int32) *GetTrackListByMailFromAndTagNameResponseBody
	GetTotal() *int32
	SetTotalPages(v string) *GetTrackListByMailFromAndTagNameResponseBody
	GetTotalPages() *string
	SetTrackList(v *GetTrackListByMailFromAndTagNameResponseBodyTrackList) *GetTrackListByMailFromAndTagNameResponseBody
	GetTrackList() *GetTrackListByMailFromAndTagNameResponseBodyTrackList
}

type GetTrackListByMailFromAndTagNameResponseBody struct {
	// Used for pagination. Not set for the first query; for subsequent queries, set to the value of OffsetCreateTime from the previous response. (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	OffsetCreateTime *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	// (This field is deprecated)
	//
	// example:
	//
	// （本字段已废弃）
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// (This field is deprecated)
	//
	// example:
	//
	// 4
	Total      *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPages *string `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// Tracking data records
	TrackList *GetTrackListByMailFromAndTagNameResponseBodyTrackList `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Struct"`
}

func (s GetTrackListByMailFromAndTagNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetOffsetCreateTime() *string {
	return s.OffsetCreateTime
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetOffsetCreateTimeDesc() *string {
	return s.OffsetCreateTimeDesc
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetTotalPages() *string {
	return s.TotalPages
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) GetTrackList() *GetTrackListByMailFromAndTagNameResponseBodyTrackList {
	return s.TrackList
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetOffsetCreateTime(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetOffsetCreateTimeDesc(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetPageNo(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetPageSize(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetRequestId(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetTotal(v int32) *GetTrackListByMailFromAndTagNameResponseBody {
	s.Total = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetTotalPages(v string) *GetTrackListByMailFromAndTagNameResponseBody {
	s.TotalPages = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) SetTrackList(v *GetTrackListByMailFromAndTagNameResponseBodyTrackList) *GetTrackListByMailFromAndTagNameResponseBody {
	s.TrackList = v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBody) Validate() error {
	if s.TrackList != nil {
		if err := s.TrackList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrackListByMailFromAndTagNameResponseBodyTrackList struct {
	Stat []*GetTrackListByMailFromAndTagNameResponseBodyTrackListStat `json:"Stat,omitempty" xml:"Stat,omitempty" type:"Repeated"`
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackList) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackList) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackList) GetStat() []*GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	return s.Stat
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackList) SetStat(v []*GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) *GetTrackListByMailFromAndTagNameResponseBodyTrackList {
	s.Stat = v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackList) Validate() error {
	if s.Stat != nil {
		for _, item := range s.Stat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTrackListByMailFromAndTagNameResponseBodyTrackListStat struct {
	// Creation time
	//
	// example:
	//
	// 2025-01-11T10:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Click count
	//
	// example:
	//
	// 0
	RcptClickCount *string `json:"RcptClickCount,omitempty" xml:"RcptClickCount,omitempty"`
	// Click rate
	//
	// example:
	//
	// 0
	RcptClickRate *string `json:"RcptClickRate,omitempty" xml:"RcptClickRate,omitempty"`
	// Number of opens
	//
	// example:
	//
	// 0
	RcptOpenCount *string `json:"RcptOpenCount,omitempty" xml:"RcptOpenCount,omitempty"`
	// Open rate
	//
	// example:
	//
	// 0
	RcptOpenRate *string `json:"RcptOpenRate,omitempty" xml:"RcptOpenRate,omitempty"`
	// Unique click count
	//
	// example:
	//
	// 0
	RcptUniqueClickCount *string `json:"RcptUniqueClickCount,omitempty" xml:"RcptUniqueClickCount,omitempty"`
	// Unique click rate
	//
	// example:
	//
	// 0
	RcptUniqueClickRate *string `json:"RcptUniqueClickRate,omitempty" xml:"RcptUniqueClickRate,omitempty"`
	// Unique open count
	//
	// example:
	//
	// 0
	RcptUniqueOpenCount *string `json:"RcptUniqueOpenCount,omitempty" xml:"RcptUniqueOpenCount,omitempty"`
	// Unique open rate
	//
	// example:
	//
	// 0
	RcptUniqueOpenRate *string `json:"RcptUniqueOpenRate,omitempty" xml:"RcptUniqueOpenRate,omitempty"`
	// Total number
	//
	// example:
	//
	// 0
	TotalNumber *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptClickCount() *string {
	return s.RcptClickCount
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptClickRate() *string {
	return s.RcptClickRate
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptOpenCount() *string {
	return s.RcptOpenCount
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptOpenRate() *string {
	return s.RcptOpenRate
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptUniqueClickCount() *string {
	return s.RcptUniqueClickCount
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptUniqueClickRate() *string {
	return s.RcptUniqueClickRate
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptUniqueOpenCount() *string {
	return s.RcptUniqueOpenCount
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetRcptUniqueOpenRate() *string {
	return s.RcptUniqueOpenRate
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) GetTotalNumber() *string {
	return s.TotalNumber
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetCreateTime(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.CreateTime = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptClickCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptClickCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptClickRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptClickRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptOpenCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptOpenCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptOpenRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptOpenRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueClickCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueClickCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueClickRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueClickRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueOpenCount(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueOpenCount = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetRcptUniqueOpenRate(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.RcptUniqueOpenRate = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) SetTotalNumber(v string) *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat {
	s.TotalNumber = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponseBodyTrackListStat) Validate() error {
	return dara.Validate(s)
}
