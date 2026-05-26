// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteConfigRequest struct {
}

func (s DeleteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigRequest) Validate() error {
	return dara.Validate(s)
}
