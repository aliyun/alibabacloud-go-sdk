// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteCredentialRequest struct {
}

func (s DeleteCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteCredentialRequest) Validate() error {
	return dara.Validate(s)
}
