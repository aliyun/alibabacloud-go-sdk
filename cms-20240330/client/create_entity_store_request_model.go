// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEntityStoreRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateEntityStoreRequest struct {
}

func (s CreateEntityStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEntityStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityStoreRequest) Validate() error {
	return dara.Validate(s)
}
