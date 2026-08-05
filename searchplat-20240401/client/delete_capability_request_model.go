// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCapabilityRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteCapabilityRequest struct {
}

func (s DeleteCapabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCapabilityRequest) GoString() string {
	return s.String()
}

func (s *DeleteCapabilityRequest) Validate() error {
	return dara.Validate(s)
}
