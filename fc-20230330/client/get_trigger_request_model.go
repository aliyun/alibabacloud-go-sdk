// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetTriggerRequest struct {
}

func (s GetTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTriggerRequest) GoString() string {
	return s.String()
}

func (s *GetTriggerRequest) Validate() error {
	return dara.Validate(s)
}
