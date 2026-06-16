// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreAPIKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetContextStoreAPIKeyRequest struct {
}

func (s GetContextStoreAPIKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreAPIKeyRequest) GoString() string {
	return s.String()
}

func (s *GetContextStoreAPIKeyRequest) Validate() error {
	return dara.Validate(s)
}
