// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetConfigRequest struct {
}

func (s GetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRequest) Validate() error {
	return dara.Validate(s)
}
