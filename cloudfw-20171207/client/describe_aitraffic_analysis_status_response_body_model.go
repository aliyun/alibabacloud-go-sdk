// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAITrafficAnalysisStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAITrafficAnalysisStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAITrafficAnalysisStatusResponseBody
	GetStatus() *string
}

type DescribeAITrafficAnalysisStatusResponseBody struct {
	// example:
	//
	// 4E7F94C7-781F-5192-86CF-DB085****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAITrafficAnalysisStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAITrafficAnalysisStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAITrafficAnalysisStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAITrafficAnalysisStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAITrafficAnalysisStatusResponseBody) SetRequestId(v string) *DescribeAITrafficAnalysisStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAITrafficAnalysisStatusResponseBody) SetStatus(v string) *DescribeAITrafficAnalysisStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAITrafficAnalysisStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
