// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductRefRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetProductRefRequest struct {
}

func (s GetProductRefRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductRefRequest) GoString() string {
	return s.String()
}

func (s *GetProductRefRequest) Validate() error {
	return dara.Validate(s)
}
