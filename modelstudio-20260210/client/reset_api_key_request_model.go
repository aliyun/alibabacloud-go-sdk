// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ResetApiKeyRequest struct {
}

func (s ResetApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ResetApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
