// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecMatchedHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecMatchedHostsResponseBodyData) *DescribeApisecMatchedHostsResponseBody
	GetData() []*DescribeApisecMatchedHostsResponseBodyData
	SetRequestId(v string) *DescribeApisecMatchedHostsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeApisecMatchedHostsResponseBody
	GetTotalCount() *string
}

type DescribeApisecMatchedHostsResponseBody struct {
	// The domain names.
	Data []*DescribeApisecMatchedHostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8D4CA088-F72B-5658-BD5B-ECE8B8F0C7BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecMatchedHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecMatchedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecMatchedHostsResponseBody) GetData() []*DescribeApisecMatchedHostsResponseBodyData {
	return s.Data
}

func (s *DescribeApisecMatchedHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecMatchedHostsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeApisecMatchedHostsResponseBody) SetData(v []*DescribeApisecMatchedHostsResponseBodyData) *DescribeApisecMatchedHostsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecMatchedHostsResponseBody) SetRequestId(v string) *DescribeApisecMatchedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecMatchedHostsResponseBody) SetTotalCount(v string) *DescribeApisecMatchedHostsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecMatchedHostsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecMatchedHostsResponseBodyData struct {
	// The number of APIs related to the domain name.
	//
	// example:
	//
	// 31
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name or IP address.
	//
	// example:
	//
	// bc.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
}

func (s DescribeApisecMatchedHostsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecMatchedHostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecMatchedHostsResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeApisecMatchedHostsResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecMatchedHostsResponseBodyData) SetCount(v int64) *DescribeApisecMatchedHostsResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeApisecMatchedHostsResponseBodyData) SetMatchedHost(v string) *DescribeApisecMatchedHostsResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecMatchedHostsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
