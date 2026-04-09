// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopJobRequest struct {
}

func (s StopJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopJobRequest) GoString() string {
	return s.String()
}

func (s *StopJobRequest) Validate() error {
	return dara.Validate(s)
}
