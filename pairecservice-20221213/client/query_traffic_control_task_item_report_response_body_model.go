// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTaskItemReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryTrafficControlTaskItemReportResponseBody
	GetRequestId() *string
	SetTrafficControlTaskItemReports(v []*QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) *QueryTrafficControlTaskItemReportResponseBody
	GetTrafficControlTaskItemReports() []*QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports
}

type QueryTrafficControlTaskItemReportResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId                     *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlTaskItemReports []*QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports `json:"TrafficControlTaskItemReports,omitempty" xml:"TrafficControlTaskItemReports,omitempty" type:"Repeated"`
}

func (s QueryTrafficControlTaskItemReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTaskItemReportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTaskItemReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTrafficControlTaskItemReportResponseBody) GetTrafficControlTaskItemReports() []*QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	return s.TrafficControlTaskItemReports
}

func (s *QueryTrafficControlTaskItemReportResponseBody) SetRequestId(v string) *QueryTrafficControlTaskItemReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBody) SetTrafficControlTaskItemReports(v []*QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) *QueryTrafficControlTaskItemReportResponseBody {
	s.TrafficControlTaskItemReports = v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBody) Validate() error {
	if s.TrafficControlTaskItemReports != nil {
		for _, item := range s.TrafficControlTaskItemReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports struct {
	// example:
	//
	// 4
	ActualItemControlNum *int64 `json:"ActualItemControlNum,omitempty" xml:"ActualItemControlNum,omitempty"`
	// example:
	//
	// 500
	ActualItemControlTraffic *int64 `json:"ActualItemControlTraffic,omitempty" xml:"ActualItemControlTraffic,omitempty"`
	// example:
	//
	// 6
	DoneItemControlNum *int64 `json:"DoneItemControlNum,omitempty" xml:"DoneItemControlNum,omitempty"`
	// example:
	//
	// 20%
	DoneItemControlPercentage *string `json:"DoneItemControlPercentage,omitempty" xml:"DoneItemControlPercentage,omitempty"`
	// example:
	//
	// 10%
	ItemControlNumPercentage *string `json:"ItemControlNumPercentage,omitempty" xml:"ItemControlNumPercentage,omitempty"`
	// example:
	//
	// 20%
	ItemControlTrafficPercentage *string `json:"ItemControlTrafficPercentage,omitempty" xml:"ItemControlTrafficPercentage,omitempty"`
	// example:
	//
	// 10
	OughtItemControlNum *int64 `json:"OughtItemControlNum,omitempty" xml:"OughtItemControlNum,omitempty"`
	// example:
	//
	// 1000
	OughtItemControlTraffic *int64 `json:"OughtItemControlTraffic,omitempty" xml:"OughtItemControlTraffic,omitempty"`
	// example:
	//
	// 3
	TrafficControlTargetId *string `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	// example:
	//
	// item-1
	TrafficControlTargetName *string `json:"TrafficControlTargetName,omitempty" xml:"TrafficControlTargetName,omitempty"`
}

func (s QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetActualItemControlNum() *int64 {
	return s.ActualItemControlNum
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetActualItemControlTraffic() *int64 {
	return s.ActualItemControlTraffic
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetDoneItemControlNum() *int64 {
	return s.DoneItemControlNum
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetDoneItemControlPercentage() *string {
	return s.DoneItemControlPercentage
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetItemControlNumPercentage() *string {
	return s.ItemControlNumPercentage
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetItemControlTrafficPercentage() *string {
	return s.ItemControlTrafficPercentage
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetOughtItemControlNum() *int64 {
	return s.OughtItemControlNum
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetOughtItemControlTraffic() *int64 {
	return s.OughtItemControlTraffic
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetTrafficControlTargetId() *string {
	return s.TrafficControlTargetId
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) GetTrafficControlTargetName() *string {
	return s.TrafficControlTargetName
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetActualItemControlNum(v int64) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.ActualItemControlNum = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetActualItemControlTraffic(v int64) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.ActualItemControlTraffic = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetDoneItemControlNum(v int64) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.DoneItemControlNum = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetDoneItemControlPercentage(v string) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.DoneItemControlPercentage = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetItemControlNumPercentage(v string) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.ItemControlNumPercentage = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetItemControlTrafficPercentage(v string) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.ItemControlTrafficPercentage = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetOughtItemControlNum(v int64) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.OughtItemControlNum = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetOughtItemControlTraffic(v int64) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.OughtItemControlTraffic = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetTrafficControlTargetId(v string) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.TrafficControlTargetId = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) SetTrafficControlTargetName(v string) *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports {
	s.TrafficControlTargetName = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportResponseBodyTrafficControlTaskItemReports) Validate() error {
	return dara.Validate(s)
}
