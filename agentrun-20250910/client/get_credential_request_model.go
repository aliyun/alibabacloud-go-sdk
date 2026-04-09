// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCredentialRequest struct {
}

func (s GetCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialRequest) Validate() error {
	return dara.Validate(s)
}
