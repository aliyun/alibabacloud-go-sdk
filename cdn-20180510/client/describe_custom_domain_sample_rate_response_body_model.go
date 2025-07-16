// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomDomainSampleRateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeCustomDomainSampleRateResponseBodyContent) *DescribeCustomDomainSampleRateResponseBody
	GetContent() *DescribeCustomDomainSampleRateResponseBodyContent
	SetRequestId(v string) *DescribeCustomDomainSampleRateResponseBody
	GetRequestId() *string
}

type DescribeCustomDomainSampleRateResponseBody struct {
	Content   *DescribeCustomDomainSampleRateResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomDomainSampleRateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomDomainSampleRateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomDomainSampleRateResponseBody) GetContent() *DescribeCustomDomainSampleRateResponseBodyContent {
	return s.Content
}

func (s *DescribeCustomDomainSampleRateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomDomainSampleRateResponseBody) SetContent(v *DescribeCustomDomainSampleRateResponseBodyContent) *DescribeCustomDomainSampleRateResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBody) SetRequestId(v string) *DescribeCustomDomainSampleRateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomDomainSampleRateResponseBodyContent struct {
	DomainContent []*DescribeCustomDomainSampleRateResponseBodyContentDomainContent `json:"DomainContent,omitempty" xml:"DomainContent,omitempty" type:"Repeated"`
	PageNumber    *int64                                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount    *int64                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCustomDomainSampleRateResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomDomainSampleRateResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) GetDomainContent() []*DescribeCustomDomainSampleRateResponseBodyContentDomainContent {
	return s.DomainContent
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) SetDomainContent(v []*DescribeCustomDomainSampleRateResponseBodyContentDomainContent) *DescribeCustomDomainSampleRateResponseBodyContent {
	s.DomainContent = v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) SetPageNumber(v int64) *DescribeCustomDomainSampleRateResponseBodyContent {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) SetPageSize(v int64) *DescribeCustomDomainSampleRateResponseBodyContent {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) SetTotalCount(v int64) *DescribeCustomDomainSampleRateResponseBodyContent {
	s.TotalCount = &v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomDomainSampleRateResponseBodyContentDomainContent struct {
	DomainName *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SampleRate *float32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s DescribeCustomDomainSampleRateResponseBodyContentDomainContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomDomainSampleRateResponseBodyContentDomainContent) GoString() string {
	return s.String()
}

func (s *DescribeCustomDomainSampleRateResponseBodyContentDomainContent) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCustomDomainSampleRateResponseBodyContentDomainContent) GetSampleRate() *float32 {
	return s.SampleRate
}

func (s *DescribeCustomDomainSampleRateResponseBodyContentDomainContent) SetDomainName(v string) *DescribeCustomDomainSampleRateResponseBodyContentDomainContent {
	s.DomainName = &v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBodyContentDomainContent) SetSampleRate(v float32) *DescribeCustomDomainSampleRateResponseBodyContentDomainContent {
	s.SampleRate = &v
	return s
}

func (s *DescribeCustomDomainSampleRateResponseBodyContentDomainContent) Validate() error {
	return dara.Validate(s)
}
