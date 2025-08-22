// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTopUrlVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllUrlList(v *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetAllUrlList() *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList
	SetDomainName(v string) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetStartTime() *string
	SetUrl200List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetUrl200List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List
	SetUrl300List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetUrl300List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List
	SetUrl400List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetUrl400List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List
	SetUrl500List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) *DescribeDcdnDomainTopUrlVisitResponseBody
	GetUrl500List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List
}

type DescribeDcdnDomainTopUrlVisitResponseBody struct {
	// A list of frequently requested URLs.
	AllUrlList *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
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
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2018-10-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of URLs for which 2xx status codes were returned.
	Url200List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	// A list of URLs for which 3xx status codes were returned.
	Url300List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	// A list of URLs for which 4xx status codes were returned.
	Url400List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	// A list of URLs for which 5xx status codes were returned.
	Url500List *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetAllUrlList() *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList {
	return s.AllUrlList
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetUrl200List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List {
	return s.Url200List
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetUrl300List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List {
	return s.Url300List
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetUrl400List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List {
	return s.Url400List
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) GetUrl500List() *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List {
	return s.Url500List
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetAllUrlList(v *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetDomainName(v string) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetRequestId(v string) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetStartTime(v string) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl200List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl300List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl400List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) SetUrl500List(v *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) *DescribeDcdnDomainTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) GetUrlList() []*DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	return s.UrlList
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList struct {
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
	// The complete URL.
	//
	// example:
	//
	// http://example.com/nn_live/nn_x64/a0.m3u8
	UrlDetail *string `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	// The number of visits.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) GetUrlList() []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	return s.UrlList
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200List) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList struct {
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
	// The complete URL.
	//
	// example:
	//
	// http://example.com/nn_live/nn_x64/a0.m3u8
	UrlDetail *string `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	// The number of visits.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) GetUrlList() []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	return s.UrlList
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300List) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList struct {
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
	// The complete URL.
	//
	// example:
	//
	// http://example.com/nn_live/nn_x64/a0.m3u8
	UrlDetail *string `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	// The number of visits.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) GetUrlList() []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	return s.UrlList
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400List) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList struct {
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
	// The complete URL.
	//
	// example:
	//
	// http://example.com/nn_live/nn_x64/a0.m3u8
	UrlDetail *string `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	// The number of visits.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) GetUrlList() []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	return s.UrlList
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500List) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList struct {
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
	// The complete URL.
	//
	// example:
	//
	// http://example.com/nn_live/nn_x64/a0.m3u8
	UrlDetail *string `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	// The number of visits.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) Validate() error {
	return dara.Validate(s)
}
