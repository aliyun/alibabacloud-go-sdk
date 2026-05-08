// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetContextStoreRequest struct {
}

func (s GetContextStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreRequest) GoString() string {
	return s.String()
}

func (s *GetContextStoreRequest) Validate() error {
	return dara.Validate(s)
}
