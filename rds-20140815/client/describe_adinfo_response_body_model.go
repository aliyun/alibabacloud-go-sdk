// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeADInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetADDNS(v string) *DescribeADInfoResponseBody
	GetADDNS() *string
	SetADServerIpAddress(v string) *DescribeADInfoResponseBody
	GetADServerIpAddress() *string
	SetADStatus(v string) *DescribeADInfoResponseBody
	GetADStatus() *string
	SetAbnormalReason(v string) *DescribeADInfoResponseBody
	GetAbnormalReason() *string
	SetRequestId(v string) *DescribeADInfoResponseBody
	GetRequestId() *string
	SetUserName(v string) *DescribeADInfoResponseBody
	GetUserName() *string
}

type DescribeADInfoResponseBody struct {
	// The DNS information about the AD domain.
	//
	// example:
	//
	// 100.100.XX.XX
	ADDNS *string `json:"ADDNS,omitempty" xml:"ADDNS,omitempty"`
	// The service IP address of the AD domain.
	//
	// example:
	//
	// 192.168.XX.XX
	ADServerIpAddress *string `json:"ADServerIpAddress,omitempty" xml:"ADServerIpAddress,omitempty"`
	// The status of the AD domain. Valid values:
	//
	// 	- **-1**: The instance is being added to the AD domain.
	//
	// 	- **0**: The instance fails to be added to the AD domain.
	//
	// 	- **1**: The instance is added to the AD domain.
	//
	// example:
	//
	// 1
	ADStatus *string `json:"ADStatus,omitempty" xml:"ADStatus,omitempty"`
	// The cause of the error.
	//
	// example:
	//
	// XXXX
	AbnormalReason *string `json:"AbnormalReason,omitempty" xml:"AbnormalReason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The username of the AD domain.
	//
	// example:
	//
	// test_01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeADInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeADInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeADInfoResponseBody) GetADDNS() *string {
	return s.ADDNS
}

func (s *DescribeADInfoResponseBody) GetADServerIpAddress() *string {
	return s.ADServerIpAddress
}

func (s *DescribeADInfoResponseBody) GetADStatus() *string {
	return s.ADStatus
}

func (s *DescribeADInfoResponseBody) GetAbnormalReason() *string {
	return s.AbnormalReason
}

func (s *DescribeADInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeADInfoResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *DescribeADInfoResponseBody) SetADDNS(v string) *DescribeADInfoResponseBody {
	s.ADDNS = &v
	return s
}

func (s *DescribeADInfoResponseBody) SetADServerIpAddress(v string) *DescribeADInfoResponseBody {
	s.ADServerIpAddress = &v
	return s
}

func (s *DescribeADInfoResponseBody) SetADStatus(v string) *DescribeADInfoResponseBody {
	s.ADStatus = &v
	return s
}

func (s *DescribeADInfoResponseBody) SetAbnormalReason(v string) *DescribeADInfoResponseBody {
	s.AbnormalReason = &v
	return s
}

func (s *DescribeADInfoResponseBody) SetRequestId(v string) *DescribeADInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeADInfoResponseBody) SetUserName(v string) *DescribeADInfoResponseBody {
	s.UserName = &v
	return s
}

func (s *DescribeADInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
