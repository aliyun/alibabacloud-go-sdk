// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgreementStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgreementCode(v string) *DescribeAgreementStatusResponseBody
	GetAgreementCode() *string
	SetRequestId(v string) *DescribeAgreementStatusResponseBody
	GetRequestId() *string
	SetStatus(v int32) *DescribeAgreementStatusResponseBody
	GetStatus() *int32
	SetUserId(v string) *DescribeAgreementStatusResponseBody
	GetUserId() *string
}

type DescribeAgreementStatusResponseBody struct {
	// example:
	//
	// 10aa40008e081ad7b1fb50bffc3a70b1
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
	// example:
	//
	// 6254E13A-A79F-5786-9C75-7590727342C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1219541161213057
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeAgreementStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgreementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusResponseBody) GetAgreementCode() *string {
	return s.AgreementCode
}

func (s *DescribeAgreementStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgreementStatusResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAgreementStatusResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *DescribeAgreementStatusResponseBody) SetAgreementCode(v string) *DescribeAgreementStatusResponseBody {
	s.AgreementCode = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetRequestId(v string) *DescribeAgreementStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetStatus(v int32) *DescribeAgreementStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetUserId(v string) *DescribeAgreementStatusResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
