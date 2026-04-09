// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetUserInfoRequest struct {
}

func (s GetUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
