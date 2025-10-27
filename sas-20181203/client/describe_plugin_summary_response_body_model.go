// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribePluginSummaryResponseBodyData) *DescribePluginSummaryResponseBody
	GetData() *DescribePluginSummaryResponseBodyData
	SetRequestId(v string) *DescribePluginSummaryResponseBody
	GetRequestId() *string
}

type DescribePluginSummaryResponseBody struct {
	// The details of the plug-in data.
	Data *DescribePluginSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePluginSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginSummaryResponseBody) GetData() *DescribePluginSummaryResponseBodyData {
	return s.Data
}

func (s *DescribePluginSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginSummaryResponseBody) SetData(v *DescribePluginSummaryResponseBodyData) *DescribePluginSummaryResponseBody {
	s.Data = v
	return s
}

func (s *DescribePluginSummaryResponseBody) SetRequestId(v string) *DescribePluginSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePluginSummaryResponseBodyData struct {
	// The number of hosts on which the plug-in failed to be installed.
	//
	// example:
	//
	// 3
	FailedCnt *int32 `json:"FailedCnt,omitempty" xml:"FailedCnt,omitempty"`
	// The causes of installation failures.
	FailedReasons []*DescribePluginSummaryResponseBodyDataFailedReasons `json:"FailedReasons,omitempty" xml:"FailedReasons,omitempty" type:"Repeated"`
	// The number of hosts on which the plug-in is offline.
	//
	// example:
	//
	// 2
	OfflineCnt *int32 `json:"OfflineCnt,omitempty" xml:"OfflineCnt,omitempty"`
	// The number of hosts on which the plug-in is online.
	//
	// example:
	//
	// 10
	OnlineCnt *int32 `json:"OnlineCnt,omitempty" xml:"OnlineCnt,omitempty"`
	// The number of hosts for which the plug-in is not enabled.
	//
	// example:
	//
	// 5
	SwitchOffCnt *int32 `json:"SwitchOffCnt,omitempty" xml:"SwitchOffCnt,omitempty"`
	// The total number of hosts.
	//
	// example:
	//
	// 20
	TotalCnt *int32 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribePluginSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePluginSummaryResponseBodyData) GetFailedCnt() *int32 {
	return s.FailedCnt
}

func (s *DescribePluginSummaryResponseBodyData) GetFailedReasons() []*DescribePluginSummaryResponseBodyDataFailedReasons {
	return s.FailedReasons
}

func (s *DescribePluginSummaryResponseBodyData) GetOfflineCnt() *int32 {
	return s.OfflineCnt
}

func (s *DescribePluginSummaryResponseBodyData) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *DescribePluginSummaryResponseBodyData) GetSwitchOffCnt() *int32 {
	return s.SwitchOffCnt
}

func (s *DescribePluginSummaryResponseBodyData) GetTotalCnt() *int32 {
	return s.TotalCnt
}

func (s *DescribePluginSummaryResponseBodyData) SetFailedCnt(v int32) *DescribePluginSummaryResponseBodyData {
	s.FailedCnt = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyData) SetFailedReasons(v []*DescribePluginSummaryResponseBodyDataFailedReasons) *DescribePluginSummaryResponseBodyData {
	s.FailedReasons = v
	return s
}

func (s *DescribePluginSummaryResponseBodyData) SetOfflineCnt(v int32) *DescribePluginSummaryResponseBodyData {
	s.OfflineCnt = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyData) SetOnlineCnt(v int32) *DescribePluginSummaryResponseBodyData {
	s.OnlineCnt = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyData) SetSwitchOffCnt(v int32) *DescribePluginSummaryResponseBodyData {
	s.SwitchOffCnt = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyData) SetTotalCnt(v int32) *DescribePluginSummaryResponseBodyData {
	s.TotalCnt = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyData) Validate() error {
	if s.FailedReasons != nil {
		for _, item := range s.FailedReasons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePluginSummaryResponseBodyDataFailedReasons struct {
	// The error code for the installation failure.
	//
	// example:
	//
	// -1
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of hosts on which the installation failed for this reason.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The cause of the installation failure.
	//
	// example:
	//
	// Other
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribePluginSummaryResponseBodyDataFailedReasons) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSummaryResponseBodyDataFailedReasons) GoString() string {
	return s.String()
}

func (s *DescribePluginSummaryResponseBodyDataFailedReasons) GetCode() *string {
	return s.Code
}

func (s *DescribePluginSummaryResponseBodyDataFailedReasons) GetCount() *int32 {
	return s.Count
}

func (s *DescribePluginSummaryResponseBodyDataFailedReasons) GetReason() *string {
	return s.Reason
}

func (s *DescribePluginSummaryResponseBodyDataFailedReasons) SetCode(v string) *DescribePluginSummaryResponseBodyDataFailedReasons {
	s.Code = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyDataFailedReasons) SetCount(v int32) *DescribePluginSummaryResponseBodyDataFailedReasons {
	s.Count = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyDataFailedReasons) SetReason(v string) *DescribePluginSummaryResponseBodyDataFailedReasons {
	s.Reason = &v
	return s
}

func (s *DescribePluginSummaryResponseBodyDataFailedReasons) Validate() error {
	return dara.Validate(s)
}
