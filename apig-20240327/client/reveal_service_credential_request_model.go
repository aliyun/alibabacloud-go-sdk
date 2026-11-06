// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevealServiceCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RevealServiceCredentialRequest struct {
}

func (s RevealServiceCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s RevealServiceCredentialRequest) GoString() string {
	return s.String()
}

func (s *RevealServiceCredentialRequest) Validate() error {
	return dara.Validate(s)
}
