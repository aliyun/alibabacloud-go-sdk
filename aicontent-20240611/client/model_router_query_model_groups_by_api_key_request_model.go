// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupsByApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterQueryModelGroupsByApiKeyRequest struct {
}

func (s ModelRouterQueryModelGroupsByApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupsByApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupsByApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
