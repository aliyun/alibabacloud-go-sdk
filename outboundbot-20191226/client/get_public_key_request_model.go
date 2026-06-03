// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPublicKeyRequest struct {
}

func (s GetPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
