// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAliasRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAliasRequest struct {
}

func (s DeleteAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAliasRequest) GoString() string {
	return s.String()
}

func (s *DeleteAliasRequest) Validate() error {
	return dara.Validate(s)
}
