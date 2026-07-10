// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteApiKeyRequest struct {
}

func (s DeleteApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
