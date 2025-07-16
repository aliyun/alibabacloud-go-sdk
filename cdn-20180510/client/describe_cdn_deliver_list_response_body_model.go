// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDeliverListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeCdnDeliverListResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeCdnDeliverListResponseBody
	GetRequestId() *string
}

type DescribeCdnDeliverListResponseBody struct {
	// The information about the tracking task.
	//
	// example:
	//
	// "data": [{"deliverId": 1,"status": "enable","createTime": "2020-10-14T11:19:26Z","crontab": "0 0 0 \\	- \\	- ?","frequency": "d","name": "The name of the tracking task","dmList": ["www.example.com"],"reports": [{"reportId": 1,"conditions": [{"op": "in","field": "prov","value": ["Heilongjiang","Beijing"]}} },{"reportId": 2}],"deliver": {"email": {"subject": "subject","to": ["example@alibaba-inc.com","example@alibaba-inc.com"]}}}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 12345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDeliverListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeCdnDeliverListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDeliverListResponseBody) SetContent(v string) *DescribeCdnDeliverListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBody) SetRequestId(v string) *DescribeCdnDeliverListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBody) Validate() error {
	return dara.Validate(s)
}
