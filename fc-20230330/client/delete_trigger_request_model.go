// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteTriggerRequest struct {
}

func (s DeleteTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteTriggerRequest) Validate() error {
	return dara.Validate(s)
}
