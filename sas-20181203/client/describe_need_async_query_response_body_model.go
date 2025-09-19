// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNeedAsyncQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DescribeNeedAsyncQueryResponseBody
	GetData() *bool
	SetRequestId(v string) *DescribeNeedAsyncQueryResponseBody
	GetRequestId() *string
}

type DescribeNeedAsyncQueryResponseBody struct {
	// Indicates whether the result is returned. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DFD6277-CC36-57F7-ACE6-F5952XXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNeedAsyncQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNeedAsyncQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNeedAsyncQueryResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeNeedAsyncQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNeedAsyncQueryResponseBody) SetData(v bool) *DescribeNeedAsyncQueryResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeNeedAsyncQueryResponseBody) SetRequestId(v string) *DescribeNeedAsyncQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNeedAsyncQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
