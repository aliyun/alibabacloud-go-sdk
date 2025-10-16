// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLProtectTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInProtectCnt(v int64) *DescribeACLProtectTrendResponseBody
	GetInProtectCnt() *int64
	SetInterVPCProtectCnt(v int64) *DescribeACLProtectTrendResponseBody
	GetInterVPCProtectCnt() *int64
	SetInterval(v int32) *DescribeACLProtectTrendResponseBody
	GetInterval() *int32
	SetOutProtectCnt(v int64) *DescribeACLProtectTrendResponseBody
	GetOutProtectCnt() *int64
	SetRequestId(v string) *DescribeACLProtectTrendResponseBody
	GetRequestId() *string
	SetTotalProtectCnt(v int64) *DescribeACLProtectTrendResponseBody
	GetTotalProtectCnt() *int64
	SetTrendList(v []*DescribeACLProtectTrendResponseBodyTrendList) *DescribeACLProtectTrendResponseBody
	GetTrendList() []*DescribeACLProtectTrendResponseBodyTrendList
}

type DescribeACLProtectTrendResponseBody struct {
	// The number of internal requests that are blocked by the ACL feature.
	//
	// example:
	//
	// 100
	InProtectCnt *int64 `json:"InProtectCnt,omitempty" xml:"InProtectCnt,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 0
	InterVPCProtectCnt *int64 `json:"InterVPCProtectCnt,omitempty" xml:"InterVPCProtectCnt,omitempty"`
	// The interval for returning data. Unit: seconds.
	//
	// example:
	//
	// 86400
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of external requests that are blocked by the ACL feature.
	//
	// example:
	//
	// 200
	OutProtectCnt *int64 `json:"OutProtectCnt,omitempty" xml:"OutProtectCnt,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9063AB86-6FFA-5B2D-A16D-697C966DECA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of requests that are blocked by the ACL feature.
	//
	// example:
	//
	// 300
	TotalProtectCnt *int64 `json:"TotalProtectCnt,omitempty" xml:"TotalProtectCnt,omitempty"`
	// The statistics on the requests that are blocked by the ACL feature.
	TrendList []*DescribeACLProtectTrendResponseBodyTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s DescribeACLProtectTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLProtectTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendResponseBody) GetInProtectCnt() *int64 {
	return s.InProtectCnt
}

func (s *DescribeACLProtectTrendResponseBody) GetInterVPCProtectCnt() *int64 {
	return s.InterVPCProtectCnt
}

func (s *DescribeACLProtectTrendResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeACLProtectTrendResponseBody) GetOutProtectCnt() *int64 {
	return s.OutProtectCnt
}

func (s *DescribeACLProtectTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeACLProtectTrendResponseBody) GetTotalProtectCnt() *int64 {
	return s.TotalProtectCnt
}

func (s *DescribeACLProtectTrendResponseBody) GetTrendList() []*DescribeACLProtectTrendResponseBodyTrendList {
	return s.TrendList
}

func (s *DescribeACLProtectTrendResponseBody) SetInProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.InProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetInterVPCProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.InterVPCProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetInterval(v int32) *DescribeACLProtectTrendResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetOutProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.OutProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetRequestId(v string) *DescribeACLProtectTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetTotalProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.TotalProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetTrendList(v []*DescribeACLProtectTrendResponseBodyTrendList) *DescribeACLProtectTrendResponseBody {
	s.TrendList = v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) Validate() error {
	if s.TrendList != nil {
		for _, item := range s.TrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeACLProtectTrendResponseBodyTrendList struct {
	// The number of requests that are blocked by ACL on the current day.
	//
	// example:
	//
	// 100
	ProtectCnt *int32 `json:"ProtectCnt,omitempty" xml:"ProtectCnt,omitempty"`
	// The UNIX timestamp at midnight (00:00:00) of each day, which indicates the date of the current day. Unit: seconds.
	//
	// example:
	//
	// 1697299200
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeACLProtectTrendResponseBodyTrendList) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLProtectTrendResponseBodyTrendList) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) GetProtectCnt() *int32 {
	return s.ProtectCnt
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) SetProtectCnt(v int32) *DescribeACLProtectTrendResponseBodyTrendList {
	s.ProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) SetTime(v int64) *DescribeACLProtectTrendResponseBodyTrendList {
	s.Time = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) Validate() error {
	return dara.Validate(s)
}
