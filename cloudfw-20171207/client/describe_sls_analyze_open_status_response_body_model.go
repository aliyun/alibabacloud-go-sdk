// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAnalyzeOpenStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpenStatus(v string) *DescribeSlsAnalyzeOpenStatusResponseBody
	GetOpenStatus() *string
	SetRequestId(v string) *DescribeSlsAnalyzeOpenStatusResponseBody
	GetRequestId() *string
}

type DescribeSlsAnalyzeOpenStatusResponseBody struct {
	// example:
	//
	// false
	OpenStatus *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// example:
	//
	// 6CC01A2B-92FB-535C-9415-9A951C20****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlsAnalyzeOpenStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAnalyzeOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsAnalyzeOpenStatusResponseBody) GetOpenStatus() *string {
	return s.OpenStatus
}

func (s *DescribeSlsAnalyzeOpenStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsAnalyzeOpenStatusResponseBody) SetOpenStatus(v string) *DescribeSlsAnalyzeOpenStatusResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *DescribeSlsAnalyzeOpenStatusResponseBody) SetRequestId(v string) *DescribeSlsAnalyzeOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsAnalyzeOpenStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
