// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliasRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAliasRequest struct {
}

func (s GetAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliasRequest) GoString() string {
	return s.String()
}

func (s *GetAliasRequest) Validate() error {
	return dara.Validate(s)
}
