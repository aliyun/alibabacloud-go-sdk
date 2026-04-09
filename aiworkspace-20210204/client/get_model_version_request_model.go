// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelVersionRequest struct {
}

func (s GetModelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelVersionRequest) GoString() string {
	return s.String()
}

func (s *GetModelVersionRequest) Validate() error {
	return dara.Validate(s)
}
