// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RestartServiceRequest struct {
}

func (s RestartServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartServiceRequest) GoString() string {
	return s.String()
}

func (s *RestartServiceRequest) Validate() error {
	return dara.Validate(s)
}
