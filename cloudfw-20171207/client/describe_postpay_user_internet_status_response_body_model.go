// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserInternetStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePostpayUserInternetStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribePostpayUserInternetStatusResponseBody
	GetStatus() *string
	SetUnprotectedDate(v int64) *DescribePostpayUserInternetStatusResponseBody
	GetUnprotectedDate() *int64
}

type DescribePostpayUserInternetStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0DC783F1-B3A7-578D-8A63-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the Internet Border firewall. Valid values:
	//
	// - **open**: The firewall is enabled.
	//
	// - **init**: The firewall is being enabled.
	//
	// - **closed**: The firewall is disabled.
	//
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of days that the firewall was disabled. This parameter is returned only if the value of the Status parameter is open.
	//
	// example:
	//
	// 20
	UnprotectedDate *int64 `json:"UnprotectedDate,omitempty" xml:"UnprotectedDate,omitempty"`
}

func (s DescribePostpayUserInternetStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserInternetStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserInternetStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayUserInternetStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribePostpayUserInternetStatusResponseBody) GetUnprotectedDate() *int64 {
	return s.UnprotectedDate
}

func (s *DescribePostpayUserInternetStatusResponseBody) SetRequestId(v string) *DescribePostpayUserInternetStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayUserInternetStatusResponseBody) SetStatus(v string) *DescribePostpayUserInternetStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePostpayUserInternetStatusResponseBody) SetUnprotectedDate(v int64) *DescribePostpayUserInternetStatusResponseBody {
	s.UnprotectedDate = &v
	return s
}

func (s *DescribePostpayUserInternetStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
