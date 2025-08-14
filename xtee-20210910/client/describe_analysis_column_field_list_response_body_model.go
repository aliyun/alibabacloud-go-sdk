// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisColumnFieldListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAnalysisColumnFieldListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeAnalysisColumnFieldListResponseBody
	GetResultObject() *bool
}

type DescribeAnalysisColumnFieldListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeAnalysisColumnFieldListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisColumnFieldListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisColumnFieldListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnalysisColumnFieldListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeAnalysisColumnFieldListResponseBody) SetRequestId(v string) *DescribeAnalysisColumnFieldListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnalysisColumnFieldListResponseBody) SetResultObject(v bool) *DescribeAnalysisColumnFieldListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeAnalysisColumnFieldListResponseBody) Validate() error {
	return dara.Validate(s)
}
