// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayHistoryServerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetRayHistoryServerRequest struct {
}

func (s GetRayHistoryServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRayHistoryServerRequest) GoString() string {
	return s.String()
}

func (s *GetRayHistoryServerRequest) Validate() error {
	return dara.Validate(s)
}
