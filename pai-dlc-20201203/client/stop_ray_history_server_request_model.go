// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRayHistoryServerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopRayHistoryServerRequest struct {
}

func (s StopRayHistoryServerRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRayHistoryServerRequest) GoString() string {
	return s.String()
}

func (s *StopRayHistoryServerRequest) Validate() error {
	return dara.Validate(s)
}
