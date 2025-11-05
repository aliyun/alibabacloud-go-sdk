// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopReferVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainTopReferVisitResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDomainTopReferVisitResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainTopReferVisitResponseBody
	GetStartTime() *string
	SetTopReferList(v *DescribeDomainTopReferVisitResponseBodyTopReferList) *DescribeDomainTopReferVisitResponseBody
	GetTopReferList() *DescribeDomainTopReferVisitResponseBodyTopReferList
}

type DescribeDomainTopReferVisitResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 95994621-8382-464B-8762-C708E73568D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2019-12-21T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The most frequently requested web pages.
	TopReferList *DescribeDomainTopReferVisitResponseBodyTopReferList `json:"TopReferList,omitempty" xml:"TopReferList,omitempty" type:"Struct"`
}

func (s DescribeDomainTopReferVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainTopReferVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopReferVisitResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainTopReferVisitResponseBody) GetTopReferList() *DescribeDomainTopReferVisitResponseBodyTopReferList {
	return s.TopReferList
}

func (s *DescribeDomainTopReferVisitResponseBody) SetDomainName(v string) *DescribeDomainTopReferVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBody) SetRequestId(v string) *DescribeDomainTopReferVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBody) SetStartTime(v string) *DescribeDomainTopReferVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBody) SetTopReferList(v *DescribeDomainTopReferVisitResponseBodyTopReferList) *DescribeDomainTopReferVisitResponseBody {
	s.TopReferList = v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBody) Validate() error {
	if s.TopReferList != nil {
		if err := s.TopReferList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainTopReferVisitResponseBodyTopReferList struct {
	ReferList []*DescribeDomainTopReferVisitResponseBodyTopReferListReferList `json:"ReferList,omitempty" xml:"ReferList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferList) GetReferList() []*DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	return s.ReferList
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferList) SetReferList(v []*DescribeDomainTopReferVisitResponseBodyTopReferListReferList) *DescribeDomainTopReferVisitResponseBodyTopReferList {
	s.ReferList = v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferList) Validate() error {
	if s.ReferList != nil {
		for _, item := range s.ReferList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopReferVisitResponseBodyTopReferListReferList struct {
	// The amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// 200
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The proportion of network traffic consumed to access the URL.
	//
	// example:
	//
	// 0.5
	FlowProportion *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	// The URLs to the most frequently requested web pages.
	//
	// example:
	//
	// learn.aliyundoc.com
	ReferDetail *string `json:"ReferDetail,omitempty" xml:"ReferDetail,omitempty"`
	// The number of visits to the URL.
	//
	// example:
	//
	// 3
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.5
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferListReferList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferListReferList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) GetReferDetail() *string {
	return s.ReferDetail
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetFlow(v string) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetFlowProportion(v float32) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetReferDetail(v string) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.ReferDetail = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitData(v string) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitProportion(v float32) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) Validate() error {
	return dara.Validate(s)
}
