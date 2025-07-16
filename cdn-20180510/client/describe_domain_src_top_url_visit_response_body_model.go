// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcTopUrlVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllUrlList(v *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) *DescribeDomainSrcTopUrlVisitResponseBody
	GetAllUrlList() *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList
	SetDomainName(v string) *DescribeDomainSrcTopUrlVisitResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDomainSrcTopUrlVisitResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainSrcTopUrlVisitResponseBody
	GetStartTime() *string
	SetUrl200List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) *DescribeDomainSrcTopUrlVisitResponseBody
	GetUrl200List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List
	SetUrl300List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) *DescribeDomainSrcTopUrlVisitResponseBody
	GetUrl300List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List
	SetUrl400List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) *DescribeDomainSrcTopUrlVisitResponseBody
	GetUrl400List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List
	SetUrl500List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) *DescribeDomainSrcTopUrlVisitResponseBody
	GetUrl500List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List
}

type DescribeDomainSrcTopUrlVisitResponseBody struct {
	// A list of frequently requested URLs.
	AllUrlList *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
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
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2018-10-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of URLs for which 2xx status codes were returned.
	Url200List *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	// A list of URLs for which 3xx status codes were returned.
	Url300List *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	// A list of URLs for which 4xx status codes were returned.
	Url400List *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	// A list of URLs for which 5xx status codes were returned.
	Url500List *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetAllUrlList() *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList {
	return s.AllUrlList
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetUrl200List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List {
	return s.Url200List
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetUrl300List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List {
	return s.Url300List
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetUrl400List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List {
	return s.Url400List
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) GetUrl500List() *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List {
	return s.Url500List
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetAllUrlList(v *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetDomainName(v string) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetRequestId(v string) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetStartTime(v string) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl200List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl300List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl400List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl500List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) GetUrlList() []*DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	return s.UrlList
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList struct {
	// The amount of network traffic. Unit: bytes.
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
	// The number of visits to the URL.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) GetUrlList() []*DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList struct {
	// The amount of network traffic. Unit: bytes.
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
	// The number of visits to the URL.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) GetUrlList() []*DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList struct {
	// The amount of network traffic. Unit: bytes.
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
	// The number of visits to the URL.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) GetUrlList() []*DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList struct {
	// The amount of network traffic. Unit: bytes.
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
	// The number of visits to the URL.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) GetUrlList() []*DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList struct {
	// The amount of network traffic. Unit: bytes.
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
	// The number of visits to the URL.
	//
	// example:
	//
	// 161673
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) Validate() error {
	return dara.Validate(s)
}
