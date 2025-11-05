// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopUrlVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllUrlList(v *DescribeDomainTopUrlVisitResponseBodyAllUrlList) *DescribeDomainTopUrlVisitResponseBody
	GetAllUrlList() *DescribeDomainTopUrlVisitResponseBodyAllUrlList
	SetDomainName(v string) *DescribeDomainTopUrlVisitResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDomainTopUrlVisitResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainTopUrlVisitResponseBody
	GetStartTime() *string
	SetUrl200List(v *DescribeDomainTopUrlVisitResponseBodyUrl200List) *DescribeDomainTopUrlVisitResponseBody
	GetUrl200List() *DescribeDomainTopUrlVisitResponseBodyUrl200List
	SetUrl300List(v *DescribeDomainTopUrlVisitResponseBodyUrl300List) *DescribeDomainTopUrlVisitResponseBody
	GetUrl300List() *DescribeDomainTopUrlVisitResponseBodyUrl300List
	SetUrl400List(v *DescribeDomainTopUrlVisitResponseBodyUrl400List) *DescribeDomainTopUrlVisitResponseBody
	GetUrl400List() *DescribeDomainTopUrlVisitResponseBodyUrl400List
	SetUrl500List(v *DescribeDomainTopUrlVisitResponseBodyUrl500List) *DescribeDomainTopUrlVisitResponseBody
	GetUrl500List() *DescribeDomainTopUrlVisitResponseBodyUrl500List
}

type DescribeDomainTopUrlVisitResponseBody struct {
	// A list of frequently requested URLs.
	AllUrlList *DescribeDomainTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
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
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-10-03T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of URLs for which 2xx status codes were returned.
	Url200List *DescribeDomainTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	// A list of URLs for which 3xx status codes were returned.
	Url300List *DescribeDomainTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	// A list of URLs for which 4xx status codes were returned.
	Url400List *DescribeDomainTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	// A list of URLs for which 5xx status codes were returned.
	Url500List *DescribeDomainTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeDomainTopUrlVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetAllUrlList() *DescribeDomainTopUrlVisitResponseBodyAllUrlList {
	return s.AllUrlList
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetUrl200List() *DescribeDomainTopUrlVisitResponseBodyUrl200List {
	return s.Url200List
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetUrl300List() *DescribeDomainTopUrlVisitResponseBodyUrl300List {
	return s.Url300List
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetUrl400List() *DescribeDomainTopUrlVisitResponseBodyUrl400List {
	return s.Url400List
}

func (s *DescribeDomainTopUrlVisitResponseBody) GetUrl500List() *DescribeDomainTopUrlVisitResponseBodyUrl500List {
	return s.Url500List
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetAllUrlList(v *DescribeDomainTopUrlVisitResponseBodyAllUrlList) *DescribeDomainTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetDomainName(v string) *DescribeDomainTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetRequestId(v string) *DescribeDomainTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetStartTime(v string) *DescribeDomainTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl200List(v *DescribeDomainTopUrlVisitResponseBodyUrl200List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl300List(v *DescribeDomainTopUrlVisitResponseBodyUrl300List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl400List(v *DescribeDomainTopUrlVisitResponseBodyUrl400List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl500List(v *DescribeDomainTopUrlVisitResponseBodyUrl500List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) Validate() error {
	if s.AllUrlList != nil {
		if err := s.AllUrlList.Validate(); err != nil {
			return err
		}
	}
	if s.Url200List != nil {
		if err := s.Url200List.Validate(); err != nil {
			return err
		}
	}
	if s.Url300List != nil {
		if err := s.Url300List.Validate(); err != nil {
			return err
		}
	}
	if s.Url400List != nil {
		if err := s.Url400List.Validate(); err != nil {
			return err
		}
	}
	if s.Url500List != nil {
		if err := s.Url500List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlList) GetUrlList() []*DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	return s.UrlList
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeDomainTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlList) Validate() error {
	if s.UrlList != nil {
		for _, item := range s.UrlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList struct {
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

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl200List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200List) GetUrlList() []*DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200List) Validate() error {
	if s.UrlList != nil {
		for _, item := range s.UrlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList struct {
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
	// http://example.com/nn_live/nn_x64/aWQ9SE5KU0bGxfcGNfbGl2ZQ,,/HNJSMPP360.m3u8
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

func (s DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl300List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300List) GetUrlList() []*DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300List) Validate() error {
	if s.UrlList != nil {
		for _, item := range s.UrlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList struct {
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

func (s DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400List) GetUrlList() []*DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400List) Validate() error {
	if s.UrlList != nil {
		for _, item := range s.UrlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList struct {
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
	// http://example.com/nn_live/nn_x64/aWQ9SE5KU01QUhbGxfcGNfbGl2ZQ,,/HNJSMPP360.m3u8
	UrlDetail *string `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	// The number of visits to the URL.
	//
	// example:
	//
	// 1884
	VisitData *string `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	// The proportion of visits to the URL.
	//
	// example:
	//
	// 0.35
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl500List) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500List) GetUrlList() []*DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	return s.UrlList
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500List) Validate() error {
	if s.UrlList != nil {
		for _, item := range s.UrlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList struct {
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
	// http://example.com/nn_live/nn_x64/aWQ9SE5KU0GNfbGl2ZQ,,/HNJSMPP360.m3u8
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

func (s DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetFlowProportion() *float32 {
	return s.FlowProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetUrlDetail() *string {
	return s.UrlDetail
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitData() *string {
	return s.VisitData
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) GetVisitProportion() *float32 {
	return s.VisitProportion
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) Validate() error {
	return dara.Validate(s)
}
