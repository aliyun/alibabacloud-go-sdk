// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeInterpreterRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCodeInterpreterRequest struct {
}

func (s GetCodeInterpreterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCodeInterpreterRequest) GoString() string {
	return s.String()
}

func (s *GetCodeInterpreterRequest) Validate() error {
	return dara.Validate(s)
}
