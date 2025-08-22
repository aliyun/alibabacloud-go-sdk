// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTopReferVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainTopReferVisitResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDcdnDomainTopReferVisitResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainTopReferVisitResponseBody
	GetStartTime() *string
	SetTopReferList(v *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) *DescribeDcdnDomainTopReferVisitResponseBody
	GetTopReferList() *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList
}

type DescribeDcdnDomainTopReferVisitResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 95994621-8382-464B-8762-C708E73568D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2018-10-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of frequently referenced URLs returned.
	TopReferList *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList `json:"TopReferList,omitempty" xml:"TopReferList,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainTopReferVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) GetTopReferList() *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList {
	return s.TopReferList
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetDomainName(v string) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetRequestId(v string) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetStartTime(v string) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) SetTopReferList(v *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) *DescribeDcdnDomainTopReferVisitResponseBody {
	s.TopReferList = v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopReferVisitResponseBodyTopReferList struct {
	ReferList []*DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList `json:"ReferList,omitempty" xml:"ReferList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) GetReferList() []*DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	return s.ReferList
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) SetReferList(v []*DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList {
	s.ReferList = v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList struct {
	// The amount of network traffic. (Unit: bytes)
	//
	// example:
	//
	// 460486880
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The proportion of network traffic consumed to access the URL.
	//
	// example:
	//
	// 0.35
	FlowProportion *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	// The complete referenced URL.
	//
	// example:
	//
	// 192.168.0.1
	ReferDetail *string `json:"ReferDetail,omitempty" xml:"ReferDetail,omitempty"`
	// The number of visits.
	//
	// example:
	//
	// 229567
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) GetReferDetail() *string {
	return s.ReferDetail
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetFlow(v string) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetFlowProportion(v float32) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetReferDetail(v string) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.ReferDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitData(v string) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitProportion(v float32) *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponseBodyTopReferListReferList) Validate() error {
	return dara.Validate(s)
}
