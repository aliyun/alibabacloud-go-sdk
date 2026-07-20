// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityCheckResultBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSecurityCheckResultBaseInfoRequest struct {
}

func (s GetSecurityCheckResultBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckResultBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckResultBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
