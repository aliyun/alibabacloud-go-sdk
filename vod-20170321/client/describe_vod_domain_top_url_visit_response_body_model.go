// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTopUrlVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllUrlList(v *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList) *DescribeVodDomainTopUrlVisitResponseBody
	GetAllUrlList() *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList
	SetDomainName(v string) *DescribeVodDomainTopUrlVisitResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeVodDomainTopUrlVisitResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainTopUrlVisitResponseBody
	GetStartTime() *string
	SetUrl200List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl200List) *DescribeVodDomainTopUrlVisitResponseBody
	GetUrl200List() *DescribeVodDomainTopUrlVisitResponseBodyUrl200List
	SetUrl300List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl300List) *DescribeVodDomainTopUrlVisitResponseBody
	GetUrl300List() *DescribeVodDomainTopUrlVisitResponseBodyUrl300List
	SetUrl400List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl400List) *DescribeVodDomainTopUrlVisitResponseBody
	GetUrl400List() *DescribeVodDomainTopUrlVisitResponseBodyUrl400List
	SetUrl500List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl500List) *DescribeVodDomainTopUrlVisitResponseBody
	GetUrl500List() *DescribeVodDomainTopUrlVisitResponseBodyUrl500List
}

type DescribeVodDomainTopUrlVisitResponseBody struct {
	AllUrlList *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
	DomainName *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime  *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Url200List *DescribeVodDomainTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	Url300List *DescribeVodDomainTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	Url400List *DescribeVodDomainTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	Url500List *DescribeVodDomainTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeVodDomainTopUrlVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetAllUrlList() *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList {
	return s.AllUrlList
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetUrl200List() *DescribeVodDomainTopUrlVisitResponseBodyUrl200List {
	return s.Url200List
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetUrl300List() *DescribeVodDomainTopUrlVisitResponseBodyUrl300List {
	return s.Url300List
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetUrl400List() *DescribeVodDomainTopUrlVisitResponseBodyUrl400List {
	return s.Url400List
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) GetUrl500List() *DescribeVodDomainTopUrlVisitResponseBodyUrl500List {
	return s.Url500List
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetAllUrlList(v *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList) *DescribeVodDomainTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetDomainName(v string) *DescribeVodDomainTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetRequestId(v string) *DescribeVodDomainTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetStartTime(v string) *DescribeVodDomainTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetUrl200List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl200List) *DescribeVodDomainTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetUrl300List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl300List) *DescribeVodDomainTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetUrl400List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl400List) *DescribeVodDomainTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) SetUrl500List(v *DescribeVodDomainTopUrlVisitResponseBodyUrl500List) *DescribeVodDomainTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyAllUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList) GetUrlList() []*DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	return s.UrlList
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyAllUrlListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl200List) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200List) GetUrlList() []*DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	return s.UrlList
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeVodDomainTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200List) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl200ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl300List) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300List) GetUrlList() []*DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	return s.UrlList
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeVodDomainTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300List) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl300ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl400List) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400List) GetUrlList() []*DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	return s.UrlList
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeVodDomainTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400List) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl400ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl500List) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500List) GetUrlList() []*DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	return s.UrlList
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeVodDomainTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500List) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponseBodyUrl500ListUrlList) Validate() error {
	return dara.Validate(s)
}
