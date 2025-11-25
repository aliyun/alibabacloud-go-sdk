// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInternetIpNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetIpNum() *int32
	SetInternetPortNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetPortNum() *int32
	SetInternetRiskIpNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetRiskIpNum() *int32
	SetInternetRiskPortNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetRiskPortNum() *int32
	SetInternetRiskServiceNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetRiskServiceNum() *int32
	SetInternetServiceNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetServiceNum() *int32
	SetInternetSlbIpNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetSlbIpNum() *int32
	SetInternetSlbIpPortNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetSlbIpPortNum() *int32
	SetInternetUnprotectedPortNum(v int32) *DescribeInternetOpenStatisticResponseBody
	GetInternetUnprotectedPortNum() *int32
	SetRequestId(v string) *DescribeInternetOpenStatisticResponseBody
	GetRequestId() *string
}

type DescribeInternetOpenStatisticResponseBody struct {
	// example:
	//
	// 9
	InternetIpNum *int32 `json:"InternetIpNum,omitempty" xml:"InternetIpNum,omitempty"`
	// example:
	//
	// 38
	InternetPortNum *int32 `json:"InternetPortNum,omitempty" xml:"InternetPortNum,omitempty"`
	// example:
	//
	// 8
	InternetRiskIpNum *int32 `json:"InternetRiskIpNum,omitempty" xml:"InternetRiskIpNum,omitempty"`
	// example:
	//
	// 5
	InternetRiskPortNum *int32 `json:"InternetRiskPortNum,omitempty" xml:"InternetRiskPortNum,omitempty"`
	// example:
	//
	// 3
	InternetRiskServiceNum *int32 `json:"InternetRiskServiceNum,omitempty" xml:"InternetRiskServiceNum,omitempty"`
	// example:
	//
	// 15
	InternetServiceNum *int32 `json:"InternetServiceNum,omitempty" xml:"InternetServiceNum,omitempty"`
	// example:
	//
	// 10
	InternetSlbIpNum *int32 `json:"InternetSlbIpNum,omitempty" xml:"InternetSlbIpNum,omitempty"`
	// example:
	//
	// 16
	InternetSlbIpPortNum *int32 `json:"InternetSlbIpPortNum,omitempty" xml:"InternetSlbIpPortNum,omitempty"`
	// example:
	//
	// 6
	InternetUnprotectedPortNum *int32 `json:"InternetUnprotectedPortNum,omitempty" xml:"InternetUnprotectedPortNum,omitempty"`
	// example:
	//
	// 6AB7822C-0D73-5D1D-81FD-45D4FB7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetOpenStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetIpNum() *int32 {
	return s.InternetIpNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetPortNum() *int32 {
	return s.InternetPortNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetRiskIpNum() *int32 {
	return s.InternetRiskIpNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetRiskPortNum() *int32 {
	return s.InternetRiskPortNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetRiskServiceNum() *int32 {
	return s.InternetRiskServiceNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetServiceNum() *int32 {
	return s.InternetServiceNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetSlbIpNum() *int32 {
	return s.InternetSlbIpNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetSlbIpPortNum() *int32 {
	return s.InternetSlbIpPortNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetInternetUnprotectedPortNum() *int32 {
	return s.InternetUnprotectedPortNum
}

func (s *DescribeInternetOpenStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetIpNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetIpNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetPortNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetPortNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetRiskIpNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetRiskIpNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetRiskPortNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetRiskPortNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetRiskServiceNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetRiskServiceNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetServiceNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetServiceNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetSlbIpNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetSlbIpNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetSlbIpPortNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetSlbIpPortNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetInternetUnprotectedPortNum(v int32) *DescribeInternetOpenStatisticResponseBody {
	s.InternetUnprotectedPortNum = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) SetRequestId(v string) *DescribeInternetOpenStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
