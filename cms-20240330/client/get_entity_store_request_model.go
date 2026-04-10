// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetEntityStoreRequest struct {
}

func (s GetEntityStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreRequest) GoString() string {
	return s.String()
}

func (s *GetEntityStoreRequest) Validate() error {
	return dara.Validate(s)
}
