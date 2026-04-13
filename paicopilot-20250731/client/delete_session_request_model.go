// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteSessionRequest struct {
}

func (s DeleteSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSessionRequest) Validate() error {
	return dara.Validate(s)
}
