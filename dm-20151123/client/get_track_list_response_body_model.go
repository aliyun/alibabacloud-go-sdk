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
	// Used for pagination. Do not set this parameter for the first request. For subsequent requests, set this parameter to the OffsetCreateTime value from the previous response. (This field is deprecated.)
	//
	// example:
	//
	// （本字段已废弃）
	OffsetCreateTime *string `json:"OffsetCreateTime,omitempty" xml:"OffsetCreateTime,omitempty"`
	// (This field is deprecated.)
	//
	// example:
	//
	// （本字段已废弃）
	OffsetCreateTimeDesc *string `json:"OffsetCreateTimeDesc,omitempty" xml:"OffsetCreateTimeDesc,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier for the request.
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of matching records.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 2
	TotalPages *int32                        `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Data       *GetTrackListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RcptClickCount       *string `json:"RcptClickCount,omitempty" xml:"RcptClickCount,omitempty"`
	RcptClickRate        *string `json:"RcptClickRate,omitempty" xml:"RcptClickRate,omitempty"`
	RcptOpenCount        *string `json:"RcptOpenCount,omitempty" xml:"RcptOpenCount,omitempty"`
	RcptOpenRate         *string `json:"RcptOpenRate,omitempty" xml:"RcptOpenRate,omitempty"`
	RcptUniqueClickCount *string `json:"RcptUniqueClickCount,omitempty" xml:"RcptUniqueClickCount,omitempty"`
	RcptUniqueClickRate  *string `json:"RcptUniqueClickRate,omitempty" xml:"RcptUniqueClickRate,omitempty"`
	RcptUniqueOpenCount  *string `json:"RcptUniqueOpenCount,omitempty" xml:"RcptUniqueOpenCount,omitempty"`
	RcptUniqueOpenRate   *string `json:"RcptUniqueOpenRate,omitempty" xml:"RcptUniqueOpenRate,omitempty"`
	TotalNumber          *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
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
