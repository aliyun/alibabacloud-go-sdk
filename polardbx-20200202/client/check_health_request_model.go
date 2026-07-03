// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHealthRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CheckHealthRequest struct {
}

func (s CheckHealthRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckHealthRequest) GoString() string {
	return s.String()
}

func (s *CheckHealthRequest) Validate() error {
	return dara.Validate(s)
}
