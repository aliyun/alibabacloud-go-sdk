// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSubListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeDcdnSubListResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeDcdnSubListResponseBody
	GetRequestId() *string
}

type DescribeDcdnSubListResponseBody struct {
	// The information about the custom report.
	//
	// example:
	//
	// "data": [{"subId": 336,"reportId": [6,8],"createTime": "2021-07-05T06:18:46Z","domains": ["example.com"],"effectiveFrom": "2021-07-05T06:18:46Z","effectiveEnd": "2021-10-05T06:18:46Z","status": "enable"}]
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3250A51D-C11D-46BA-B6B3-95348EEDE652
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSubListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSubListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSubListResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnSubListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSubListResponseBody) SetContent(v string) *DescribeDcdnSubListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnSubListResponseBody) SetRequestId(v string) *DescribeDcdnSubListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSubListResponseBody) Validate() error {
	return dara.Validate(s)
}
