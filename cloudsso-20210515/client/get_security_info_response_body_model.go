// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSecurityInfoResponseBody
	GetRequestId() *string
	SetSecurityEmail(v *GetSecurityInfoResponseBodySecurityEmail) *GetSecurityInfoResponseBody
	GetSecurityEmail() *GetSecurityInfoResponseBodySecurityEmail
}

type GetSecurityInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EE42D2C4-A30A-59B7-ACEB-6D22FB44174A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security email information of the user.
	//
	// > If the user has not configured a security email address, this field is not returned.
	SecurityEmail *GetSecurityInfoResponseBodySecurityEmail `json:"SecurityEmail,omitempty" xml:"SecurityEmail,omitempty" type:"Struct"`
}

func (s GetSecurityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityInfoResponseBody) GetSecurityEmail() *GetSecurityInfoResponseBodySecurityEmail {
	return s.SecurityEmail
}

func (s *GetSecurityInfoResponseBody) SetRequestId(v string) *GetSecurityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityInfoResponseBody) SetSecurityEmail(v *GetSecurityInfoResponseBodySecurityEmail) *GetSecurityInfoResponseBody {
	s.SecurityEmail = v
	return s
}

func (s *GetSecurityInfoResponseBody) Validate() error {
	if s.SecurityEmail != nil {
		if err := s.SecurityEmail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityInfoResponseBodySecurityEmail struct {
	// The security email address of the user.
	//
	// example:
	//
	// Alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The verification status of the security email address. Valid values:
	//
	// - active: Verified.
	//
	// - pending: Pending verification.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSecurityInfoResponseBodySecurityEmail) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityInfoResponseBodySecurityEmail) GoString() string {
	return s.String()
}

func (s *GetSecurityInfoResponseBodySecurityEmail) GetEmail() *string {
	return s.Email
}

func (s *GetSecurityInfoResponseBodySecurityEmail) GetStatus() *string {
	return s.Status
}

func (s *GetSecurityInfoResponseBodySecurityEmail) SetEmail(v string) *GetSecurityInfoResponseBodySecurityEmail {
	s.Email = &v
	return s
}

func (s *GetSecurityInfoResponseBodySecurityEmail) SetStatus(v string) *GetSecurityInfoResponseBodySecurityEmail {
	s.Status = &v
	return s
}

func (s *GetSecurityInfoResponseBodySecurityEmail) Validate() error {
	return dara.Validate(s)
}
