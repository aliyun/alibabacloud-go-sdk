// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAccountAttributesRequest struct {
}

func (s GetAccountAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetAccountAttributesRequest) Validate() error {
	return dara.Validate(s)
}
