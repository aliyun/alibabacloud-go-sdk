// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceIspsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsps(v []*DescribePortViewSourceIspsResponseBodyIsps) *DescribePortViewSourceIspsResponseBody
	GetIsps() []*DescribePortViewSourceIspsResponseBodyIsps
	SetRequestId(v string) *DescribePortViewSourceIspsResponseBody
	GetRequestId() *string
}

type DescribePortViewSourceIspsResponseBody struct {
	// An array that consists of the details of the ISP.
	Isps []*DescribePortViewSourceIspsResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortViewSourceIspsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceIspsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsResponseBody) GetIsps() []*DescribePortViewSourceIspsResponseBodyIsps {
	return s.Isps
}

func (s *DescribePortViewSourceIspsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortViewSourceIspsResponseBody) SetIsps(v []*DescribePortViewSourceIspsResponseBodyIsps) *DescribePortViewSourceIspsResponseBody {
	s.Isps = v
	return s
}

func (s *DescribePortViewSourceIspsResponseBody) SetRequestId(v string) *DescribePortViewSourceIspsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortViewSourceIspsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePortViewSourceIspsResponseBodyIsps struct {
	// The total number of requests that are sent from the ISP.
	//
	// > This parameter does not indicate the accurate number of requests. You can use this parameter to calculate the proportion of requests from different ISPs.
	//
	// example:
	//
	// 3390671
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the ISP. For more information, see the ISP codes table.
	//
	// example:
	//
	// 100017
	IspId *string `json:"IspId,omitempty" xml:"IspId,omitempty"`
}

func (s DescribePortViewSourceIspsResponseBodyIsps) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceIspsResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsResponseBodyIsps) GetCount() *int64 {
	return s.Count
}

func (s *DescribePortViewSourceIspsResponseBodyIsps) GetIspId() *string {
	return s.IspId
}

func (s *DescribePortViewSourceIspsResponseBodyIsps) SetCount(v int64) *DescribePortViewSourceIspsResponseBodyIsps {
	s.Count = &v
	return s
}

func (s *DescribePortViewSourceIspsResponseBodyIsps) SetIspId(v string) *DescribePortViewSourceIspsResponseBodyIsps {
	s.IspId = &v
	return s
}

func (s *DescribePortViewSourceIspsResponseBodyIsps) Validate() error {
	return dara.Validate(s)
}
