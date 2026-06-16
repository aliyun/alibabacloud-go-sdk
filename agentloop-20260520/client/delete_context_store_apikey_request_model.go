// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextStoreAPIKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteContextStoreAPIKeyRequest struct {
}

func (s DeleteContextStoreAPIKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextStoreAPIKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextStoreAPIKeyRequest) Validate() error {
	return dara.Validate(s)
}
