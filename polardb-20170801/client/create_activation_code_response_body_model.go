// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActivationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivateAt(v string) *CreateActivationCodeResponseBody
	GetActivateAt() *string
	SetCertContentB64(v string) *CreateActivationCodeResponseBody
	GetCertContentB64() *string
	SetDescription(v string) *CreateActivationCodeResponseBody
	GetDescription() *string
	SetExpireAt(v string) *CreateActivationCodeResponseBody
	GetExpireAt() *string
	SetGmtCreated(v string) *CreateActivationCodeResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *CreateActivationCodeResponseBody
	GetGmtModified() *string
	SetId(v int32) *CreateActivationCodeResponseBody
	GetId() *int32
	SetMacAddress(v string) *CreateActivationCodeResponseBody
	GetMacAddress() *string
	SetName(v string) *CreateActivationCodeResponseBody
	GetName() *string
	SetRequestId(v string) *CreateActivationCodeResponseBody
	GetRequestId() *string
	SetSystemIdentifier(v string) *CreateActivationCodeResponseBody
	GetSystemIdentifier() *string
}

type CreateActivationCodeResponseBody struct {
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
	// The activation code ID.
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
	// 4CE6DF97-AEA4-484F-906F-C407EE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The system identifier of the database.
	//
	// example:
	//
	// 1234567890123456
	SystemIdentifier *string `json:"SystemIdentifier,omitempty" xml:"SystemIdentifier,omitempty"`
}

func (s CreateActivationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateActivationCodeResponseBody) GetActivateAt() *string {
	return s.ActivateAt
}

func (s *CreateActivationCodeResponseBody) GetCertContentB64() *string {
	return s.CertContentB64
}

func (s *CreateActivationCodeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateActivationCodeResponseBody) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *CreateActivationCodeResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateActivationCodeResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateActivationCodeResponseBody) GetId() *int32 {
	return s.Id
}

func (s *CreateActivationCodeResponseBody) GetMacAddress() *string {
	return s.MacAddress
}

func (s *CreateActivationCodeResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateActivationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateActivationCodeResponseBody) GetSystemIdentifier() *string {
	return s.SystemIdentifier
}

func (s *CreateActivationCodeResponseBody) SetActivateAt(v string) *CreateActivationCodeResponseBody {
	s.ActivateAt = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetCertContentB64(v string) *CreateActivationCodeResponseBody {
	s.CertContentB64 = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetDescription(v string) *CreateActivationCodeResponseBody {
	s.Description = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetExpireAt(v string) *CreateActivationCodeResponseBody {
	s.ExpireAt = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetGmtCreated(v string) *CreateActivationCodeResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetGmtModified(v string) *CreateActivationCodeResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetId(v int32) *CreateActivationCodeResponseBody {
	s.Id = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetMacAddress(v string) *CreateActivationCodeResponseBody {
	s.MacAddress = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetName(v string) *CreateActivationCodeResponseBody {
	s.Name = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetRequestId(v string) *CreateActivationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateActivationCodeResponseBody) SetSystemIdentifier(v string) *CreateActivationCodeResponseBody {
	s.SystemIdentifier = &v
	return s
}

func (s *CreateActivationCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
