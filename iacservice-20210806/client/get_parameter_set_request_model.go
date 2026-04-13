// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterSetRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetParameterSetRequest struct {
}

func (s GetParameterSetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParameterSetRequest) GoString() string {
	return s.String()
}

func (s *GetParameterSetRequest) Validate() error {
	return dara.Validate(s)
}
