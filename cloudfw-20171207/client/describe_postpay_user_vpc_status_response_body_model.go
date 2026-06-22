// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserVpcStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePostpayUserVpcStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribePostpayUserVpcStatusResponseBody
	GetStatus() *string
	SetUnprotectedDate(v int64) *DescribePostpayUserVpcStatusResponseBody
	GetUnprotectedDate() *int64
}

type DescribePostpayUserVpcStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7447795A-39AB-52CB-8F92-128DF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the Cloud Firewall VPC firewall switch module. Valid values:
	//
	// - **open**: The module is enabled.
	//
	// - **init**: The module is being enabled.
	//
	// - **closed**: The module is not enabled.
	//
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of days during which protection is not enabled. This parameter is valid only when the module is enabled.
	//
	// example:
	//
	// 20
	UnprotectedDate *int64 `json:"UnprotectedDate,omitempty" xml:"UnprotectedDate,omitempty"`
}

func (s DescribePostpayUserVpcStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserVpcStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserVpcStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayUserVpcStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribePostpayUserVpcStatusResponseBody) GetUnprotectedDate() *int64 {
	return s.UnprotectedDate
}

func (s *DescribePostpayUserVpcStatusResponseBody) SetRequestId(v string) *DescribePostpayUserVpcStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayUserVpcStatusResponseBody) SetStatus(v string) *DescribePostpayUserVpcStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePostpayUserVpcStatusResponseBody) SetUnprotectedDate(v int64) *DescribePostpayUserVpcStatusResponseBody {
	s.UnprotectedDate = &v
	return s
}

func (s *DescribePostpayUserVpcStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
