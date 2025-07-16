// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSubListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeCdnSubListResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeCdnSubListResponseBody
	GetRequestId() *string
}

type DescribeCdnSubListResponseBody struct {
	// The information about the custom report task.
	//
	// example:
	//
	// {"RequestId":"3250A51D-C11D-46BA-B6B3-95348EEDE652","Description":"Successful","Content":{"data":[{"subId":5,"reportId":[1,2,3],"createTime":"2020-09-25T09:39:33Z","domains"["www.example.com","www.example.com"],"effectiveFrom":"2020-09-17T00:00:00Z","effectiveEnd":"2020-11-17T00:00:00Z","status":"enable"}]}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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

func (s *DescribeCdnSubListResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeCdnSubListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnSubListResponseBody) SetContent(v string) *DescribeCdnSubListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCdnSubListResponseBody) SetRequestId(v string) *DescribeCdnSubListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnSubListResponseBody) Validate() error {
	return dara.Validate(s)
}
