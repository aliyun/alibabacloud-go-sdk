// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertActionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAlertActionRequest struct {
}

func (s GetAlertActionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertActionRequest) GoString() string {
	return s.String()
}

func (s *GetAlertActionRequest) Validate() error {
	return dara.Validate(s)
}
