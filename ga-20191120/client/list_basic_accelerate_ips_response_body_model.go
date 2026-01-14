// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAccelerateIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIps(v []*ListBasicAccelerateIpsResponseBodyAccelerateIps) *ListBasicAccelerateIpsResponseBody
	GetAccelerateIps() []*ListBasicAccelerateIpsResponseBodyAccelerateIps
	SetMaxResults(v int32) *ListBasicAccelerateIpsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListBasicAccelerateIpsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBasicAccelerateIpsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBasicAccelerateIpsResponseBody
	GetTotalCount() *int32
}

type ListBasicAccelerateIpsResponseBody struct {
	// The accelerated IP addresses of the basic GA instance.
	AccelerateIps []*ListBasicAccelerateIpsResponseBodyAccelerateIps `json:"AccelerateIps,omitempty" xml:"AccelerateIps,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If **NextToken*	- was not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBasicAccelerateIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsResponseBody) GetAccelerateIps() []*ListBasicAccelerateIpsResponseBodyAccelerateIps {
	return s.AccelerateIps
}

func (s *ListBasicAccelerateIpsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBasicAccelerateIpsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBasicAccelerateIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBasicAccelerateIpsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBasicAccelerateIpsResponseBody) SetAccelerateIps(v []*ListBasicAccelerateIpsResponseBodyAccelerateIps) *ListBasicAccelerateIpsResponseBody {
	s.AccelerateIps = v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetMaxResults(v int32) *ListBasicAccelerateIpsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetNextToken(v string) *ListBasicAccelerateIpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetRequestId(v string) *ListBasicAccelerateIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetTotalCount(v int32) *ListBasicAccelerateIpsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) Validate() error {
	if s.AccelerateIps != nil {
		for _, item := range s.AccelerateIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBasicAccelerateIpsResponseBodyAccelerateIps struct {
	// The accelerated IP address of the basic GA instance.
	//
	// example:
	//
	// 116.132.XX.XX
	AccelerateIpAddress *string `json:"AccelerateIpAddress,omitempty" xml:"AccelerateIpAddress,omitempty"`
	// The ID of the accelerated IP address of the basic GA instance.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the acceleration region.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The status of the accelerated IP address. Valid values:
	//
	// 	- **active**: The accelerated IP address is available.
	//
	// 	- **binding**: The accelerated IP address is being associated.
	//
	// 	- **bound**: The accelerated IP address is associated.
	//
	// 	- **unbinding**: The accelerated IP address is being disassociated.
	//
	// 	- **deleting**: The accelerated IP address is being deleted.
	//
	// >  This parameter is unavailable when the accelerated IP address is being created.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListBasicAccelerateIpsResponseBodyAccelerateIps) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpsResponseBodyAccelerateIps) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) GetAccelerateIpAddress() *string {
	return s.AccelerateIpAddress
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) GetIpSetId() *string {
	return s.IpSetId
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) GetState() *string {
	return s.State
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetAccelerateIpAddress(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.AccelerateIpAddress = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetAccelerateIpId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetAcceleratorId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetIpSetId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.IpSetId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetState(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.State = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) Validate() error {
	return dara.Validate(s)
}
