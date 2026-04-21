// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetNormalServiceConfigRequest struct {
}

func (s GetNormalServiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNormalServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNormalServiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
