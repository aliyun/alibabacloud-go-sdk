// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatusCodeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainStatusCodeListResponseBody
	GetRequestId() *string
	SetStatusCodeList(v []*DescribeDomainStatusCodeListResponseBodyStatusCodeList) *DescribeDomainStatusCodeListResponseBody
	GetStatusCodeList() []*DescribeDomainStatusCodeListResponseBodyStatusCodeList
}

type DescribeDomainStatusCodeListResponseBody struct {
	// example:
	//
	// 3B63C0DD-8AC5-44B2-95D6-064CA9296B9C
	RequestId      *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusCodeList []*DescribeDomainStatusCodeListResponseBodyStatusCodeList `json:"StatusCodeList,omitempty" xml:"StatusCodeList,omitempty" type:"Repeated"`
}

func (s DescribeDomainStatusCodeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainStatusCodeListResponseBody) GetStatusCodeList() []*DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	return s.StatusCodeList
}

func (s *DescribeDomainStatusCodeListResponseBody) SetRequestId(v string) *DescribeDomainStatusCodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBody) SetStatusCodeList(v []*DescribeDomainStatusCodeListResponseBodyStatusCodeList) *DescribeDomainStatusCodeListResponseBody {
	s.StatusCodeList = v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainStatusCodeListResponseBodyStatusCodeList struct {
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 320
	Status200 *int64 `json:"Status200,omitempty" xml:"Status200,omitempty"`
	// example:
	//
	// 5776
	Status2XX *int64 `json:"Status2XX,omitempty" xml:"Status2XX,omitempty"`
	// example:
	//
	// 0
	Status3XX *int64 `json:"Status3XX,omitempty" xml:"Status3XX,omitempty"`
	// example:
	//
	// 0
	Status403 *int64 `json:"Status403,omitempty" xml:"Status403,omitempty"`
	// example:
	//
	// 34
	Status404 *int64 `json:"Status404,omitempty" xml:"Status404,omitempty"`
	// example:
	//
	// 11
	Status405 *int64 `json:"Status405,omitempty" xml:"Status405,omitempty"`
	Status410 *int64 `json:"Status410,omitempty" xml:"Status410,omitempty"`
	Status499 *int64 `json:"Status499,omitempty" xml:"Status499,omitempty"`
	// example:
	//
	// 168
	Status4XX *int64 `json:"Status4XX,omitempty" xml:"Status4XX,omitempty"`
	// example:
	//
	// 0
	Status501 *int64 `json:"Status501,omitempty" xml:"Status501,omitempty"`
	// example:
	//
	// 3
	Status502 *int64 `json:"Status502,omitempty" xml:"Status502,omitempty"`
	// example:
	//
	// 0
	Status503 *int64 `json:"Status503,omitempty" xml:"Status503,omitempty"`
	// example:
	//
	// 0
	Status504 *int64 `json:"Status504,omitempty" xml:"Status504,omitempty"`
	// example:
	//
	// 7
	Status5XX *int64 `json:"Status5XX,omitempty" xml:"Status5XX,omitempty"`
}

func (s DescribeDomainStatusCodeListResponseBodyStatusCodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponseBodyStatusCodeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus200() *int64 {
	return s.Status200
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus2XX() *int64 {
	return s.Status2XX
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus3XX() *int64 {
	return s.Status3XX
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus403() *int64 {
	return s.Status403
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus404() *int64 {
	return s.Status404
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus405() *int64 {
	return s.Status405
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus410() *int64 {
	return s.Status410
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus499() *int64 {
	return s.Status499
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus4XX() *int64 {
	return s.Status4XX
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus501() *int64 {
	return s.Status501
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus502() *int64 {
	return s.Status502
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus503() *int64 {
	return s.Status503
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus504() *int64 {
	return s.Status504
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) GetStatus5XX() *int64 {
	return s.Status5XX
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetIndex(v int32) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Index = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus200(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status200 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus2XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status2XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus3XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status3XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus403(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status403 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus404(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status404 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus405(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status405 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus410(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status410 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus499(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status499 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus4XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status4XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus501(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status501 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus502(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status502 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus503(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status503 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus504(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status504 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus5XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status5XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) Validate() error {
	return dara.Validate(s)
}
