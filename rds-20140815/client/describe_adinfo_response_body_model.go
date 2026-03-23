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
	ADDNS             *string `json:"ADDNS,omitempty" xml:"ADDNS,omitempty"`
	ADServerIpAddress *string `json:"ADServerIpAddress,omitempty" xml:"ADServerIpAddress,omitempty"`
	ADStatus          *string `json:"ADStatus,omitempty" xml:"ADStatus,omitempty"`
	AbnormalReason    *string `json:"AbnormalReason,omitempty" xml:"AbnormalReason,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserName          *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
