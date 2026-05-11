// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRolloutRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteServiceRolloutRequest struct {
}

func (s DeleteServiceRolloutRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRolloutRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRolloutRequest) Validate() error {
	return dara.Validate(s)
}
