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
	SetTotalAlertCnt(v int64) *DescribeACLProtectTrendResponseBody
	GetTotalAlertCnt() *int64
	SetTotalPassCnt(v int64) *DescribeACLProtectTrendResponseBody
	GetTotalPassCnt() *int64
	SetTotalProtectCnt(v int64) *DescribeACLProtectTrendResponseBody
	GetTotalProtectCnt() *int64
	SetTrendList(v []*DescribeACLProtectTrendResponseBodyTrendList) *DescribeACLProtectTrendResponseBody
	GetTrendList() []*DescribeACLProtectTrendResponseBodyTrendList
}

type DescribeACLProtectTrendResponseBody struct {
	// The number of inbound blocks by Internet access control.
	//
	// example:
	//
	// 100
	InProtectCnt *int64 `json:"InProtectCnt,omitempty" xml:"InProtectCnt,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated.
	//
	// example:
	//
	// 0
	InterVPCProtectCnt *int64 `json:"InterVPCProtectCnt,omitempty" xml:"InterVPCProtectCnt,omitempty"`
	// The step size of the returned data, in seconds. This indicates the interval at which a result is returned.
	//
	// example:
	//
	// 86400
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of outbound blocks by Internet access control.
	//
	// example:
	//
	// 200
	OutProtectCnt *int64 `json:"OutProtectCnt,omitempty" xml:"OutProtectCnt,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9063AB86-6FFA-5B2D-A16D-697C966DECA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of AlertCnt across all time points within the queried time range. This is the total number of sessions for which the monitor (alert) action is performed after an access control policy is matched during the entire time period.
	//
	// example:
	//
	// 100
	TotalAlertCnt *int64 `json:"TotalAlertCnt,omitempty" xml:"TotalAlertCnt,omitempty"`
	// The total number of PassCnt across all time points within the queried time range. This is the total number of sessions for which the allow action is performed after an access control policy is matched during the entire time period.
	//
	// example:
	//
	// 100
	TotalPassCnt *int64 `json:"TotalPassCnt,omitempty" xml:"TotalPassCnt,omitempty"`
	// The total number of blocks by Internet access control.
	//
	// example:
	//
	// 300
	TotalProtectCnt *int64 `json:"TotalProtectCnt,omitempty" xml:"TotalProtectCnt,omitempty"`
	// The list of trend data for Internet access control blocks.
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

func (s *DescribeACLProtectTrendResponseBody) GetTotalAlertCnt() *int64 {
	return s.TotalAlertCnt
}

func (s *DescribeACLProtectTrendResponseBody) GetTotalPassCnt() *int64 {
	return s.TotalPassCnt
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

func (s *DescribeACLProtectTrendResponseBody) SetTotalAlertCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.TotalAlertCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetTotalPassCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.TotalPassCnt = &v
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
	// The total number of sessions for which the monitor (alert) action is performed after an access control policy is matched within the time point.
	//
	// example:
	//
	// 20
	AlertCnt *int32 `json:"AlertCnt,omitempty" xml:"AlertCnt,omitempty"`
	// The total number of sessions for which the allow action is performed after an access control policy is matched within the time point.
	//
	// example:
	//
	// 10
	PassCnt *int32 `json:"PassCnt,omitempty" xml:"PassCnt,omitempty"`
	// The number of blocks by Internet access control on the day.
	//
	// example:
	//
	// 100
	ProtectCnt *int32 `json:"ProtectCnt,omitempty" xml:"ProtectCnt,omitempty"`
	// The UNIX timestamp of 00:00:00 each day, in seconds, which represents the date.
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

func (s *DescribeACLProtectTrendResponseBodyTrendList) GetAlertCnt() *int32 {
	return s.AlertCnt
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) GetPassCnt() *int32 {
	return s.PassCnt
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) GetProtectCnt() *int32 {
	return s.ProtectCnt
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) SetAlertCnt(v int32) *DescribeACLProtectTrendResponseBodyTrendList {
	s.AlertCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) SetPassCnt(v int32) *DescribeACLProtectTrendResponseBodyTrendList {
	s.PassCnt = &v
	return s
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
