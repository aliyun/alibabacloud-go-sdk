// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOffsetCreateTime(v string) *GetTrackListResponseBody
	GetOffsetCreateTime() *string
	SetOffsetCreateTimeDesc(v string) *GetTrackListResponseBody
	GetOffsetCreateTimeDesc() *string
	SetPageNo(v int32) *GetTrackListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetTrackListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetTrackListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *GetTrackListResponseBody
	GetTotal() *int32
	SetTotalPages(v int32) *GetTrackListResponseBody
	GetTotalPages() *int32
	SetData(v *GetTrackListResponseBodyData) *GetTrackListResponseBody
	GetData() *GetTrackListResponseBodyData
}

type GetTrackListResponseBody struct {
	// Used for pagination. Not set for the first query, but for subsequent queries, it should be set to the value of OffsetCreateTime from the previous response. (This field is deprecated)
	//
	// example:
	//
	// (This field is deprecated)
	OffsetCreateTime *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	// (This field is deprecated)
	//
	// example:
	//
	// (This field is deprecated)
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of items
	//
	// example:
	//
	// 100
	Total      *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// Tracking data records
	Data *GetTrackListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetTrackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBody) GetOffsetCreateTime() *string {
	return s.OffsetCreateTime
}

func (s *GetTrackListResponseBody) GetOffsetCreateTimeDesc() *string {
	return s.OffsetCreateTimeDesc
}

func (s *GetTrackListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetTrackListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTrackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrackListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetTrackListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *GetTrackListResponseBody) GetData() *GetTrackListResponseBodyData {
	return s.Data
}

func (s *GetTrackListResponseBody) SetOffsetCreateTime(v string) *GetTrackListResponseBody {
	s.OffsetCreateTime = &v
	return s
}

func (s *GetTrackListResponseBody) SetOffsetCreateTimeDesc(v string) *GetTrackListResponseBody {
	s.OffsetCreateTimeDesc = &v
	return s
}

func (s *GetTrackListResponseBody) SetPageNo(v int32) *GetTrackListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetTrackListResponseBody) SetPageSize(v int32) *GetTrackListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTrackListResponseBody) SetRequestId(v string) *GetTrackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrackListResponseBody) SetTotal(v int32) *GetTrackListResponseBody {
	s.Total = &v
	return s
}

func (s *GetTrackListResponseBody) SetTotalPages(v int32) *GetTrackListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *GetTrackListResponseBody) SetData(v *GetTrackListResponseBodyData) *GetTrackListResponseBody {
	s.Data = v
	return s
}

func (s *GetTrackListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrackListResponseBodyData struct {
	Stat []*GetTrackListResponseBodyDataStat `json:"stat,omitempty" xml:"stat,omitempty" type:"Repeated"`
}

func (s GetTrackListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBodyData) GetStat() []*GetTrackListResponseBodyDataStat {
	return s.Stat
}

func (s *GetTrackListResponseBodyData) SetStat(v []*GetTrackListResponseBodyDataStat) *GetTrackListResponseBodyData {
	s.Stat = v
	return s
}

func (s *GetTrackListResponseBodyData) Validate() error {
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

type GetTrackListResponseBodyDataStat struct {
	// Creation time
	//
	// example:
	//
	// 2019-09-29T13:28Z
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
	// Number of Opens
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

func (s GetTrackListResponseBodyDataStat) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListResponseBodyDataStat) GoString() string {
	return s.String()
}

func (s *GetTrackListResponseBodyDataStat) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTrackListResponseBodyDataStat) GetRcptClickCount() *string {
	return s.RcptClickCount
}

func (s *GetTrackListResponseBodyDataStat) GetRcptClickRate() *string {
	return s.RcptClickRate
}

func (s *GetTrackListResponseBodyDataStat) GetRcptOpenCount() *string {
	return s.RcptOpenCount
}

func (s *GetTrackListResponseBodyDataStat) GetRcptOpenRate() *string {
	return s.RcptOpenRate
}

func (s *GetTrackListResponseBodyDataStat) GetRcptUniqueClickCount() *string {
	return s.RcptUniqueClickCount
}

func (s *GetTrackListResponseBodyDataStat) GetRcptUniqueClickRate() *string {
	return s.RcptUniqueClickRate
}

func (s *GetTrackListResponseBodyDataStat) GetRcptUniqueOpenCount() *string {
	return s.RcptUniqueOpenCount
}

func (s *GetTrackListResponseBodyDataStat) GetRcptUniqueOpenRate() *string {
	return s.RcptUniqueOpenRate
}

func (s *GetTrackListResponseBodyDataStat) GetTotalNumber() *string {
	return s.TotalNumber
}

func (s *GetTrackListResponseBodyDataStat) SetCreateTime(v string) *GetTrackListResponseBodyDataStat {
	s.CreateTime = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptOpenCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueClickRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueClickRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenCount(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenCount = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetRcptUniqueOpenRate(v string) *GetTrackListResponseBodyDataStat {
	s.RcptUniqueOpenRate = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) SetTotalNumber(v string) *GetTrackListResponseBodyDataStat {
	s.TotalNumber = &v
	return s
}

func (s *GetTrackListResponseBodyDataStat) Validate() error {
	return dara.Validate(s)
}
