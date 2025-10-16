// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserNatStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePostpayUserNatStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribePostpayUserNatStatusResponseBody
	GetStatus() *string
	SetUnprotectedDate(v int64) *DescribePostpayUserNatStatusResponseBody
	GetUnprotectedDate() *int64
}

type DescribePostpayUserNatStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6B780BD6-282C-51A9-A8E6-59F636******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the NAT Firewall feature. Valid values:
	//
	// 	- **open**: enabled
	//
	// 	- **init**: being enabled
	//
	// 	- **closed**: disabled
	//
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of days during which no asset is added to the NAT Firewall feature for protection. This parameter is valid only when the value of Status is open.
	//
	// example:
	//
	// 20
	UnprotectedDate *int64 `json:"UnprotectedDate,omitempty" xml:"UnprotectedDate,omitempty"`
}

func (s DescribePostpayUserNatStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserNatStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserNatStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayUserNatStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribePostpayUserNatStatusResponseBody) GetUnprotectedDate() *int64 {
	return s.UnprotectedDate
}

func (s *DescribePostpayUserNatStatusResponseBody) SetRequestId(v string) *DescribePostpayUserNatStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayUserNatStatusResponseBody) SetStatus(v string) *DescribePostpayUserNatStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePostpayUserNatStatusResponseBody) SetUnprotectedDate(v int64) *DescribePostpayUserNatStatusResponseBody {
	s.UnprotectedDate = &v
	return s
}

func (s *DescribePostpayUserNatStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
