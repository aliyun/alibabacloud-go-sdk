// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReportRecipientStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportRecipientList(v []*DescribeReportRecipientStatusResponseBodyReportRecipientList) *DescribeReportRecipientStatusResponseBody
	GetReportRecipientList() []*DescribeReportRecipientStatusResponseBodyReportRecipientList
	SetRequestId(v string) *DescribeReportRecipientStatusResponseBody
	GetRequestId() *string
}

type DescribeReportRecipientStatusResponseBody struct {
	// The report recipients.
	ReportRecipientList []*DescribeReportRecipientStatusResponseBodyReportRecipientList `json:"ReportRecipientList,omitempty" xml:"ReportRecipientList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0DC1F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeReportRecipientStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReportRecipientStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReportRecipientStatusResponseBody) GetReportRecipientList() []*DescribeReportRecipientStatusResponseBodyReportRecipientList {
	return s.ReportRecipientList
}

func (s *DescribeReportRecipientStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReportRecipientStatusResponseBody) SetReportRecipientList(v []*DescribeReportRecipientStatusResponseBodyReportRecipientList) *DescribeReportRecipientStatusResponseBody {
	s.ReportRecipientList = v
	return s
}

func (s *DescribeReportRecipientStatusResponseBody) SetRequestId(v string) *DescribeReportRecipientStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReportRecipientStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeReportRecipientStatusResponseBodyReportRecipientList struct {
	// Indicates whether the email address is verified. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// >  If no email is specified when you create a report, the value of this parameter is empty.
	//
	// example:
	//
	// 1
	IsVerify *int32 `json:"IsVerify,omitempty" xml:"IsVerify,omitempty"`
	// The email address of the report recipient.
	//
	// >  If no email is specified when you create a report, the value of this parameter is empty.
	//
	// example:
	//
	// username@example.com
	Recipient *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
}

func (s DescribeReportRecipientStatusResponseBodyReportRecipientList) String() string {
	return dara.Prettify(s)
}

func (s DescribeReportRecipientStatusResponseBodyReportRecipientList) GoString() string {
	return s.String()
}

func (s *DescribeReportRecipientStatusResponseBodyReportRecipientList) GetIsVerify() *int32 {
	return s.IsVerify
}

func (s *DescribeReportRecipientStatusResponseBodyReportRecipientList) GetRecipient() *string {
	return s.Recipient
}

func (s *DescribeReportRecipientStatusResponseBodyReportRecipientList) SetIsVerify(v int32) *DescribeReportRecipientStatusResponseBodyReportRecipientList {
	s.IsVerify = &v
	return s
}

func (s *DescribeReportRecipientStatusResponseBodyReportRecipientList) SetRecipient(v string) *DescribeReportRecipientStatusResponseBodyReportRecipientList {
	s.Recipient = &v
	return s
}

func (s *DescribeReportRecipientStatusResponseBodyReportRecipientList) Validate() error {
	return dara.Validate(s)
}
