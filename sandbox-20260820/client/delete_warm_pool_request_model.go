// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarmPoolRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteWarmPoolRequest struct {
}

func (s DeleteWarmPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarmPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarmPoolRequest) Validate() error {
	return dara.Validate(s)
}
