// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateQwenEduApiRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ActivateQwenEduApiRequest struct {
}

func (s ActivateQwenEduApiRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateQwenEduApiRequest) GoString() string {
	return s.String()
}

func (s *ActivateQwenEduApiRequest) Validate() error {
	return dara.Validate(s)
}
