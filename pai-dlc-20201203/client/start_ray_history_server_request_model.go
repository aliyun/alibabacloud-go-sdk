// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRayHistoryServerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StartRayHistoryServerRequest struct {
}

func (s StartRayHistoryServerRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRayHistoryServerRequest) GoString() string {
	return s.String()
}

func (s *StartRayHistoryServerRequest) Validate() error {
	return dara.Validate(s)
}
