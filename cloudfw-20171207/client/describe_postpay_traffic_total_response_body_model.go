// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayTrafficTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePostpayTrafficTotalResponseBody
	GetRequestId() *string
	SetTotalAssets(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalAssets() *int64
	SetTotalBillTraffic(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalBillTraffic() *int64
	SetTotalInternetAssets(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalInternetAssets() *int64
	SetTotalInternetTraffic(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalInternetTraffic() *int64
	SetTotalNatAssets(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalNatAssets() *int64
	SetTotalNatTraffic(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalNatTraffic() *int64
	SetTotalSdlBillTraffic(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalSdlBillTraffic() *int64
	SetTotalSdlFreeTraffic(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalSdlFreeTraffic() *int64
	SetTotalTraffic(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalTraffic() *int64
	SetTotalVpcAssets(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalVpcAssets() *int64
	SetTotalVpcTraffic(v int64) *DescribePostpayTrafficTotalResponseBody
	GetTotalVpcTraffic() *int64
}

type DescribePostpayTrafficTotalResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 826B6280-9704-5643-97B1-6B47AC3F027A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of assets that are protected by border firewalls.
	//
	// example:
	//
	// 132
	TotalAssets *int64 `json:"TotalAssets,omitempty" xml:"TotalAssets,omitempty"`
	// For the subscription edition, this is the total billed elastic traffic after deductions are applied. Unit: bytes.
	//
	// example:
	//
	// 2320274874426
	TotalBillTraffic *int64 `json:"TotalBillTraffic,omitempty" xml:"TotalBillTraffic,omitempty"`
	// The total number of assets that are protected by Internet Border firewalls.
	//
	// example:
	//
	// 128
	TotalInternetAssets *int64 `json:"TotalInternetAssets,omitempty" xml:"TotalInternetAssets,omitempty"`
	// The total traffic of the Internet Border. For the subscription edition, this is the total elastic traffic of the Internet Border. Unit: bytes.
	//
	// example:
	//
	// 2320274874426
	TotalInternetTraffic *int64 `json:"TotalInternetTraffic,omitempty" xml:"TotalInternetTraffic,omitempty"`
	// The total number of assets that are protected by NAT border firewalls.
	//
	// example:
	//
	// 1
	TotalNatAssets *int64 `json:"TotalNatAssets,omitempty" xml:"TotalNatAssets,omitempty"`
	// The total traffic of the NAT border. For the subscription edition, this is the total elastic traffic of the NAT border. Unit: bytes.
	//
	// example:
	//
	// 560646279
	TotalNatTraffic *int64 `json:"TotalNatTraffic,omitempty" xml:"TotalNatTraffic,omitempty"`
	// The total billed traffic for data leakage detection.
	//
	// example:
	//
	// 0
	TotalSdlBillTraffic *int64 `json:"TotalSdlBillTraffic,omitempty" xml:"TotalSdlBillTraffic,omitempty"`
	// The total free traffic for data leakage detection.
	//
	// example:
	//
	// 0
	TotalSdlFreeTraffic *int64 `json:"TotalSdlFreeTraffic,omitempty" xml:"TotalSdlFreeTraffic,omitempty"`
	// The total traffic. For the subscription edition, this is the total elastic traffic. Unit: bytes.
	//
	// example:
	//
	// 2320274874426
	TotalTraffic *int64 `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	// The total number of assets that are protected by VPC border firewalls.
	//
	// example:
	//
	// 3
	TotalVpcAssets *int64 `json:"TotalVpcAssets,omitempty" xml:"TotalVpcAssets,omitempty"`
	// The total traffic of the VPC border. For the subscription edition, this is the total elastic traffic of the VPC border. Unit: bytes.
	//
	// example:
	//
	// 2320274874426
	TotalVpcTraffic *int64 `json:"TotalVpcTraffic,omitempty" xml:"TotalVpcTraffic,omitempty"`
}

func (s DescribePostpayTrafficTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayTrafficTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalAssets() *int64 {
	return s.TotalAssets
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalBillTraffic() *int64 {
	return s.TotalBillTraffic
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalInternetAssets() *int64 {
	return s.TotalInternetAssets
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalInternetTraffic() *int64 {
	return s.TotalInternetTraffic
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalNatAssets() *int64 {
	return s.TotalNatAssets
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalNatTraffic() *int64 {
	return s.TotalNatTraffic
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalSdlBillTraffic() *int64 {
	return s.TotalSdlBillTraffic
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalSdlFreeTraffic() *int64 {
	return s.TotalSdlFreeTraffic
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalTraffic() *int64 {
	return s.TotalTraffic
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalVpcAssets() *int64 {
	return s.TotalVpcAssets
}

func (s *DescribePostpayTrafficTotalResponseBody) GetTotalVpcTraffic() *int64 {
	return s.TotalVpcTraffic
}

func (s *DescribePostpayTrafficTotalResponseBody) SetRequestId(v string) *DescribePostpayTrafficTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalAssets(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalAssets = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalBillTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalBillTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalInternetAssets(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalInternetAssets = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalInternetTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalInternetTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalNatAssets(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalNatAssets = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalNatTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalNatTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalSdlBillTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalSdlBillTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalSdlFreeTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalSdlFreeTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalVpcAssets(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalVpcAssets = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalVpcTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalVpcTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) Validate() error {
	return dara.Validate(s)
}
