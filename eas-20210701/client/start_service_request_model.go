// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StartServiceRequest struct {
}

func (s StartServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartServiceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceRequest) Validate() error {
	return dara.Validate(s)
}
