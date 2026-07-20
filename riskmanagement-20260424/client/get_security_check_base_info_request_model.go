// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityCheckBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSecurityCheckBaseInfoRequest struct {
}

func (s GetSecurityCheckBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
