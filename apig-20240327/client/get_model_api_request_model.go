// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelApiRequest struct {
}

func (s GetModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelApiRequest) GoString() string {
	return s.String()
}

func (s *GetModelApiRequest) Validate() error {
	return dara.Validate(s)
}
