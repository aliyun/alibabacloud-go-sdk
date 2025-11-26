// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentifyCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyCredential(v *IdentifyCredential) *CreateIdentifyCredentialRequest
	GetIdentifyCredential() *IdentifyCredential
}

type CreateIdentifyCredentialRequest struct {
	IdentifyCredential *IdentifyCredential `json:"IdentifyCredential,omitempty" xml:"IdentifyCredential,omitempty"`
}

func (s CreateIdentifyCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentifyCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentifyCredentialRequest) GetIdentifyCredential() *IdentifyCredential {
	return s.IdentifyCredential
}

func (s *CreateIdentifyCredentialRequest) SetIdentifyCredential(v *IdentifyCredential) *CreateIdentifyCredentialRequest {
	s.IdentifyCredential = v
	return s
}

func (s *CreateIdentifyCredentialRequest) Validate() error {
	if s.IdentifyCredential != nil {
		if err := s.IdentifyCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}
