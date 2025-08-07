// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationCodeDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivateAt(v string) *DescribeActivationCodeDetailsResponseBody
	GetActivateAt() *string
	SetCertContentB64(v string) *DescribeActivationCodeDetailsResponseBody
	GetCertContentB64() *string
	SetDescription(v string) *DescribeActivationCodeDetailsResponseBody
	GetDescription() *string
	SetExpireAt(v string) *DescribeActivationCodeDetailsResponseBody
	GetExpireAt() *string
	SetGmtCreated(v string) *DescribeActivationCodeDetailsResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *DescribeActivationCodeDetailsResponseBody
	GetGmtModified() *string
	SetId(v int32) *DescribeActivationCodeDetailsResponseBody
	GetId() *int32
	SetMacAddress(v string) *DescribeActivationCodeDetailsResponseBody
	GetMacAddress() *string
	SetName(v string) *DescribeActivationCodeDetailsResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeActivationCodeDetailsResponseBody
	GetRequestId() *string
	SetSystemIdentifier(v string) *DescribeActivationCodeDetailsResponseBody
	GetSystemIdentifier() *string
}

type DescribeActivationCodeDetailsResponseBody struct {
	// The time when the activation code takes effect.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	ActivateAt *string `json:"ActivateAt,omitempty" xml:"ActivateAt,omitempty"`
	// The activation code in the base64 format. The activation code is decoded and stored into a file named license.lic. PolarDB can access and read the license.lic file upon startup to validate the license or perform related operations.
	//
	// example:
	//
	// AAEAA******AAA=
	CertContentB64 *string `json:"CertContentB64,omitempty" xml:"CertContentB64,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// testCode
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the activation code expires.
	//
	// example:
	//
	// 2054-10-09 16:46:20
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the activation code was last updated.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the activation code.
	//
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The MAC address.
	//
	// example:
	//
	// 12:34:56:78:98:00
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The name of the activation code.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2A9EFA7-915F-4572-8299-85A307******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The system identifier of the database.
	//
	// example:
	//
	// 1234567890123456
	SystemIdentifier *string `json:"SystemIdentifier,omitempty" xml:"SystemIdentifier,omitempty"`
}

func (s DescribeActivationCodeDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationCodeDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActivationCodeDetailsResponseBody) GetActivateAt() *string {
	return s.ActivateAt
}

func (s *DescribeActivationCodeDetailsResponseBody) GetCertContentB64() *string {
	return s.CertContentB64
}

func (s *DescribeActivationCodeDetailsResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeActivationCodeDetailsResponseBody) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *DescribeActivationCodeDetailsResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeActivationCodeDetailsResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeActivationCodeDetailsResponseBody) GetId() *int32 {
	return s.Id
}

func (s *DescribeActivationCodeDetailsResponseBody) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeActivationCodeDetailsResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeActivationCodeDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActivationCodeDetailsResponseBody) GetSystemIdentifier() *string {
	return s.SystemIdentifier
}

func (s *DescribeActivationCodeDetailsResponseBody) SetActivateAt(v string) *DescribeActivationCodeDetailsResponseBody {
	s.ActivateAt = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetCertContentB64(v string) *DescribeActivationCodeDetailsResponseBody {
	s.CertContentB64 = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetDescription(v string) *DescribeActivationCodeDetailsResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetExpireAt(v string) *DescribeActivationCodeDetailsResponseBody {
	s.ExpireAt = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetGmtCreated(v string) *DescribeActivationCodeDetailsResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetGmtModified(v string) *DescribeActivationCodeDetailsResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetId(v int32) *DescribeActivationCodeDetailsResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetMacAddress(v string) *DescribeActivationCodeDetailsResponseBody {
	s.MacAddress = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetName(v string) *DescribeActivationCodeDetailsResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetRequestId(v string) *DescribeActivationCodeDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) SetSystemIdentifier(v string) *DescribeActivationCodeDetailsResponseBody {
	s.SystemIdentifier = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
