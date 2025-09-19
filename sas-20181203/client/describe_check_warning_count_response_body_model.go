// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeCheckWarningCountResponseBody
	GetCount() *int32
	SetRequestId(v string) *DescribeCheckWarningCountResponseBody
	GetRequestId() *string
}

type DescribeCheckWarningCountResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9693CBA1-1EC4-5B5A-8D96-34010D9DXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCheckWarningCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningCountResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCheckWarningCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckWarningCountResponseBody) SetCount(v int32) *DescribeCheckWarningCountResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCheckWarningCountResponseBody) SetRequestId(v string) *DescribeCheckWarningCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningCountResponseBody) Validate() error {
	return dara.Validate(s)
}
