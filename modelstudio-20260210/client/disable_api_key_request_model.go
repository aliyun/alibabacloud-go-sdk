// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DisableApiKeyRequest struct {
}

func (s DisableApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DisableApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
