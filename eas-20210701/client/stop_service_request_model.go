// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopServiceRequest struct {
}

func (s StopServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopServiceRequest) GoString() string {
	return s.String()
}

func (s *StopServiceRequest) Validate() error {
	return dara.Validate(s)
}
