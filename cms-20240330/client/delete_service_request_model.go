// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteServiceRequest struct {
}

func (s DeleteServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) Validate() error {
	return dara.Validate(s)
}
