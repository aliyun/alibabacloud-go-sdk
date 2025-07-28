// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserApiRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserApiRequestResponseBody
	GetRequestId() *string
	SetRequests(v []*DescribeUserApiRequestResponseBodyRequests) *DescribeUserApiRequestResponseBody
	GetRequests() []*DescribeUserApiRequestResponseBodyRequests
}

type DescribeUserApiRequestResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D13E4540-4432-5AD7-B216-6369512514F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics.
	Requests []*DescribeUserApiRequestResponseBodyRequests `json:"Requests,omitempty" xml:"Requests,omitempty" type:"Repeated"`
}

func (s DescribeUserApiRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserApiRequestResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserApiRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserApiRequestResponseBody) GetRequests() []*DescribeUserApiRequestResponseBodyRequests {
	return s.Requests
}

func (s *DescribeUserApiRequestResponseBody) SetRequestId(v string) *DescribeUserApiRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserApiRequestResponseBody) SetRequests(v []*DescribeUserApiRequestResponseBodyRequests) *DescribeUserApiRequestResponseBody {
	s.Requests = v
	return s
}

func (s *DescribeUserApiRequestResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserApiRequestResponseBodyRequests struct {
	// The number of entries returned.
	//
	// example:
	//
	// 76
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of the statistics. Valid values:
	//
	// 	- **client_list**: client
	//
	// 	- **ip**: IP address
	//
	// 	- **region_id*	- region
	//
	// 	- **country_id**: country
	//
	// example:
	//
	// {
	//
	//     "client_list": [
	//
	//         "Unknown"
	//
	//     ],
	//
	//     "ip": "47.92.113.***",
	//
	//     "region_id": "110000",
	//
	//     "country_id": "CN"
	//
	// }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeUserApiRequestResponseBodyRequests) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserApiRequestResponseBodyRequests) GoString() string {
	return s.String()
}

func (s *DescribeUserApiRequestResponseBodyRequests) GetCount() *int64 {
	return s.Count
}

func (s *DescribeUserApiRequestResponseBodyRequests) GetValue() *string {
	return s.Value
}

func (s *DescribeUserApiRequestResponseBodyRequests) SetCount(v int64) *DescribeUserApiRequestResponseBodyRequests {
	s.Count = &v
	return s
}

func (s *DescribeUserApiRequestResponseBodyRequests) SetValue(v string) *DescribeUserApiRequestResponseBodyRequests {
	s.Value = &v
	return s
}

func (s *DescribeUserApiRequestResponseBodyRequests) Validate() error {
	return dara.Validate(s)
}
