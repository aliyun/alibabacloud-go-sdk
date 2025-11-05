// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSubListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeCdnSubListResponseBodyContent) *DescribeCdnSubListResponseBody
	GetContent() *DescribeCdnSubListResponseBodyContent
	SetRequestId(v string) *DescribeCdnSubListResponseBody
	GetRequestId() *string
}

type DescribeCdnSubListResponseBody struct {
	// The information about the custom report task.
	//
	// example:
	//
	// {"RequestId":"3250A51D-C11D-46BA-B6B3-95348EEDE652","Description":"Successful","Content":{"data":[{"subId":5,"reportId":[1,2,3],"createTime":"2020-09-25T09:39:33Z","domains"["www.example.com","www.example.com"],"effectiveFrom":"2020-09-17T00:00:00Z","effectiveEnd":"2020-11-17T00:00:00Z","status":"enable"}]}}
	Content *DescribeCdnSubListResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3250A51D-C11D-46BA-B6B3-95348EEDE652
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnSubListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSubListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnSubListResponseBody) GetContent() *DescribeCdnSubListResponseBodyContent {
	return s.Content
}

func (s *DescribeCdnSubListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnSubListResponseBody) SetContent(v *DescribeCdnSubListResponseBodyContent) *DescribeCdnSubListResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnSubListResponseBody) SetRequestId(v string) *DescribeCdnSubListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnSubListResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnSubListResponseBodyContent struct {
	Data []*DescribeCdnSubListResponseBodyContentData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeCdnSubListResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSubListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCdnSubListResponseBodyContent) GetData() []*DescribeCdnSubListResponseBodyContentData {
	return s.Data
}

func (s *DescribeCdnSubListResponseBodyContent) SetData(v []*DescribeCdnSubListResponseBodyContentData) *DescribeCdnSubListResponseBodyContent {
	s.Data = v
	return s
}

func (s *DescribeCdnSubListResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnSubListResponseBodyContentData struct {
	// example:
	//
	// 2024-05-16T09:43:38Z
	CreateTime *string   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Domains    []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-05-16T09:43:38Z
	EffectiveEnd *string `json:"effectiveEnd,omitempty" xml:"effectiveEnd,omitempty"`
	// example:
	//
	// 2024-05-16T09:43:38Z
	EffectiveFrom *string  `json:"effectiveFrom,omitempty" xml:"effectiveFrom,omitempty"`
	ReportId      []*int64 `json:"reportId,omitempty" xml:"reportId,omitempty" type:"Repeated"`
	// example:
	//
	// enable
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	SubId *int64 `json:"subId,omitempty" xml:"subId,omitempty"`
}

func (s DescribeCdnSubListResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSubListResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *DescribeCdnSubListResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCdnSubListResponseBodyContentData) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeCdnSubListResponseBodyContentData) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *DescribeCdnSubListResponseBodyContentData) GetEffectiveFrom() *string {
	return s.EffectiveFrom
}

func (s *DescribeCdnSubListResponseBodyContentData) GetReportId() []*int64 {
	return s.ReportId
}

func (s *DescribeCdnSubListResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnSubListResponseBodyContentData) GetSubId() *int64 {
	return s.SubId
}

func (s *DescribeCdnSubListResponseBodyContentData) SetCreateTime(v string) *DescribeCdnSubListResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *DescribeCdnSubListResponseBodyContentData) SetDomains(v []*string) *DescribeCdnSubListResponseBodyContentData {
	s.Domains = v
	return s
}

func (s *DescribeCdnSubListResponseBodyContentData) SetEffectiveEnd(v string) *DescribeCdnSubListResponseBodyContentData {
	s.EffectiveEnd = &v
	return s
}

func (s *DescribeCdnSubListResponseBodyContentData) SetEffectiveFrom(v string) *DescribeCdnSubListResponseBodyContentData {
	s.EffectiveFrom = &v
	return s
}

func (s *DescribeCdnSubListResponseBodyContentData) SetReportId(v []*int64) *DescribeCdnSubListResponseBodyContentData {
	s.ReportId = v
	return s
}

func (s *DescribeCdnSubListResponseBodyContentData) SetStatus(v string) *DescribeCdnSubListResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *DescribeCdnSubListResponseBodyContentData) SetSubId(v int64) *DescribeCdnSubListResponseBodyContentData {
	s.SubId = &v
	return s
}

func (s *DescribeCdnSubListResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
