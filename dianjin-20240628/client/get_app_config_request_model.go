// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAppConfigRequest struct {
}

func (s GetAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAppConfigRequest) Validate() error {
	return dara.Validate(s)
}
