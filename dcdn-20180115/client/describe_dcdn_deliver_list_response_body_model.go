// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDeliverListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeDcdnDeliverListResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeDcdnDeliverListResponseBody
	GetRequestId() *string
}

type DescribeDcdnDeliverListResponseBody struct {
	// The information about the tracking task.
	//
	// example:
	//
	// "data": [{"deliverId": 1,"status": "enable","createTime": "2021-06-14T11:19:26Z","crontab": "0 0 0 \\	- \\	- ?","frequency": "d","name": "Domain name report","dmList": ["www.example.com"],"reports": [{"reportId": 1,"conditions": [{"op": "in","field": "prov","value": ["Heilongjiang","Beijing"]}]},{"reportId": 2}],"deliver": {"email": {"subject": "subject","to": ["username@example.com","username@example.org"],"copy":["username@example.com","username@example.org"]}}}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDeliverListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeliverListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeliverListResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnDeliverListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDeliverListResponseBody) SetContent(v string) *DescribeDcdnDeliverListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnDeliverListResponseBody) SetRequestId(v string) *DescribeDcdnDeliverListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDeliverListResponseBody) Validate() error {
	return dara.Validate(s)
}
