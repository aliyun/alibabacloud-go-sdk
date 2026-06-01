// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRayHistoryServerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRayHistoryServerRequest struct {
}

func (s DeleteRayHistoryServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRayHistoryServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteRayHistoryServerRequest) Validate() error {
	return dara.Validate(s)
}
