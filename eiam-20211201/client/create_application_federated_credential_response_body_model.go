// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationFederatedCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialId(v string) *CreateApplicationFederatedCredentialResponseBody
	GetApplicationFederatedCredentialId() *string
	SetRequestId(v string) *CreateApplicationFederatedCredentialResponseBody
	GetRequestId() *string
}

type CreateApplicationFederatedCredentialResponseBody struct {
	// example:
	//
	// afc_asd123daxxxx
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationFederatedCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialResponseBody) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *CreateApplicationFederatedCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationFederatedCredentialResponseBody) SetApplicationFederatedCredentialId(v string) *CreateApplicationFederatedCredentialResponseBody {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialResponseBody) SetRequestId(v string) *CreateApplicationFederatedCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationFederatedCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
