// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetServiceNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInternetServiceNameListResponseBody
	GetRequestId() *string
	SetServiceNameList(v []*string) *DescribeInternetServiceNameListResponseBody
	GetServiceNameList() []*string
}

type DescribeInternetServiceNameListResponseBody struct {
	// example:
	//
	// 91B01BCD-DFB0-5CA8-9191-5B38C62****
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceNameList []*string `json:"ServiceNameList,omitempty" xml:"ServiceNameList,omitempty" type:"Repeated"`
}

func (s DescribeInternetServiceNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetServiceNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetServiceNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetServiceNameListResponseBody) GetServiceNameList() []*string {
	return s.ServiceNameList
}

func (s *DescribeInternetServiceNameListResponseBody) SetRequestId(v string) *DescribeInternetServiceNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetServiceNameListResponseBody) SetServiceNameList(v []*string) *DescribeInternetServiceNameListResponseBody {
	s.ServiceNameList = v
	return s
}

func (s *DescribeInternetServiceNameListResponseBody) Validate() error {
	return dara.Validate(s)
}
