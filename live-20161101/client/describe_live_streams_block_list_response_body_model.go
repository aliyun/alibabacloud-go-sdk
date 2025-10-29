// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsBlockListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamsBlockListResponseBody
	GetDomainName() *string
	SetPageNum(v int32) *DescribeLiveStreamsBlockListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamsBlockListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveStreamsBlockListResponseBody
	GetRequestId() *string
	SetStreamUrls(v *DescribeLiveStreamsBlockListResponseBodyStreamUrls) *DescribeLiveStreamsBlockListResponseBody
	GetStreamUrls() *DescribeLiveStreamsBlockListResponseBodyStreamUrls
	SetTotalNum(v int32) *DescribeLiveStreamsBlockListResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveStreamsBlockListResponseBody
	GetTotalPage() *int32
}

type DescribeLiveStreamsBlockListResponseBody struct {
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9D855EC8-CF71-4615-B164-F7DFCB23B41D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The complete URL of each live stream.
	StreamUrls *DescribeLiveStreamsBlockListResponseBodyStreamUrls `json:"StreamUrls,omitempty" xml:"StreamUrls,omitempty" type:"Struct"`
	// The total number of live stream URLs that meet the specified conditions.
	//
	// example:
	//
	// 11
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveStreamsBlockListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsBlockListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsBlockListResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsBlockListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamsBlockListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsBlockListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamsBlockListResponseBody) GetStreamUrls() *DescribeLiveStreamsBlockListResponseBodyStreamUrls {
	return s.StreamUrls
}

func (s *DescribeLiveStreamsBlockListResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveStreamsBlockListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveStreamsBlockListResponseBody) SetDomainName(v string) *DescribeLiveStreamsBlockListResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBody) SetPageNum(v int32) *DescribeLiveStreamsBlockListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBody) SetPageSize(v int32) *DescribeLiveStreamsBlockListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBody) SetRequestId(v string) *DescribeLiveStreamsBlockListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBody) SetStreamUrls(v *DescribeLiveStreamsBlockListResponseBodyStreamUrls) *DescribeLiveStreamsBlockListResponseBody {
	s.StreamUrls = v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBody) SetTotalNum(v int32) *DescribeLiveStreamsBlockListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBody) SetTotalPage(v int32) *DescribeLiveStreamsBlockListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBody) Validate() error {
	if s.StreamUrls != nil {
		if err := s.StreamUrls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamsBlockListResponseBodyStreamUrls struct {
	StreamUrl []*string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamsBlockListResponseBodyStreamUrls) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsBlockListResponseBodyStreamUrls) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsBlockListResponseBodyStreamUrls) GetStreamUrl() []*string {
	return s.StreamUrl
}

func (s *DescribeLiveStreamsBlockListResponseBodyStreamUrls) SetStreamUrl(v []*string) *DescribeLiveStreamsBlockListResponseBodyStreamUrls {
	s.StreamUrl = v
	return s
}

func (s *DescribeLiveStreamsBlockListResponseBodyStreamUrls) Validate() error {
	return dara.Validate(s)
}
