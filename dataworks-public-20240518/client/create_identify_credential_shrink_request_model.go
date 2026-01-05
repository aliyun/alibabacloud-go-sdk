// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentifyCredentialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyCredentialShrink(v string) *CreateIdentifyCredentialShrinkRequest
	GetIdentifyCredentialShrink() *string
}

type CreateIdentifyCredentialShrinkRequest struct {
	// The user credential object.
	IdentifyCredentialShrink *string `json:"IdentifyCredential,omitempty" xml:"IdentifyCredential,omitempty"`
}

func (s CreateIdentifyCredentialShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentifyCredentialShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentifyCredentialShrinkRequest) GetIdentifyCredentialShrink() *string {
	return s.IdentifyCredentialShrink
}

func (s *CreateIdentifyCredentialShrinkRequest) SetIdentifyCredentialShrink(v string) *CreateIdentifyCredentialShrinkRequest {
	s.IdentifyCredentialShrink = &v
	return s
}

func (s *CreateIdentifyCredentialShrinkRequest) Validate() error {
	return dara.Validate(s)
}
