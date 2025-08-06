// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTopReferVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainTopReferVisitResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeVodDomainTopReferVisitResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainTopReferVisitResponseBody
	GetStartTime() *string
	SetTopReferList(v *DescribeVodDomainTopReferVisitResponseBodyTopReferList) *DescribeVodDomainTopReferVisitResponseBody
	GetTopReferList() *DescribeVodDomainTopReferVisitResponseBodyTopReferList
}

type DescribeVodDomainTopReferVisitResponseBody struct {
	DomainName   *string                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopReferList *DescribeVodDomainTopReferVisitResponseBodyTopReferList `json:"TopReferList,omitempty" xml:"TopReferList,omitempty" type:"Struct"`
}

func (s DescribeVodDomainTopReferVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopReferVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopReferVisitResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainTopReferVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainTopReferVisitResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainTopReferVisitResponseBody) GetTopReferList() *DescribeVodDomainTopReferVisitResponseBodyTopReferList {
	return s.TopReferList
}

func (s *DescribeVodDomainTopReferVisitResponseBody) SetDomainName(v string) *DescribeVodDomainTopReferVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBody) SetRequestId(v string) *DescribeVodDomainTopReferVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBody) SetStartTime(v string) *DescribeVodDomainTopReferVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBody) SetTopReferList(v *DescribeVodDomainTopReferVisitResponseBodyTopReferList) *DescribeVodDomainTopReferVisitResponseBody {
	s.TopReferList = v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopReferVisitResponseBodyTopReferList struct {
	ReferList []*DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList `json:"ReferList,omitempty" xml:"ReferList,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTopReferVisitResponseBodyTopReferList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopReferVisitResponseBodyTopReferList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferList) GetReferList() []*DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList {
	return s.ReferList
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferList) SetReferList(v []*DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) *DescribeVodDomainTopReferVisitResponseBodyTopReferList {
	s.ReferList = v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	ReferDetail     *string  `json:"ReferDetail,omitempty" xml:"ReferDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) GetReferDetail() *string {
	return s.ReferDetail
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) SetFlow(v string) *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList {
	s.Flow = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) SetFlowProportion(v float32) *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) SetReferDetail(v string) *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList {
	s.ReferDetail = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitData(v string) *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitData = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitProportion(v float32) *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponseBodyTopReferListReferList) Validate() error {
	return dara.Validate(s)
}
