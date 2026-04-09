// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCodeSourceRequest struct {
}

func (s GetCodeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *GetCodeSourceRequest) Validate() error {
	return dara.Validate(s)
}
