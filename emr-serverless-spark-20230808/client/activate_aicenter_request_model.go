// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateAICenterRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ActivateAICenterRequest struct {
}

func (s ActivateAICenterRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateAICenterRequest) GoString() string {
	return s.String()
}

func (s *ActivateAICenterRequest) Validate() error {
	return dara.Validate(s)
}
