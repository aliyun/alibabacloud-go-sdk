// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterSetRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteParameterSetRequest struct {
}

func (s DeleteParameterSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterSetRequest) Validate() error {
	return dara.Validate(s)
}
