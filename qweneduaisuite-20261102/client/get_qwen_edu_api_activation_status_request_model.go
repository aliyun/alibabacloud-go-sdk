// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQwenEduApiActivationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetQwenEduApiActivationStatusRequest struct {
}

func (s GetQwenEduApiActivationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQwenEduApiActivationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetQwenEduApiActivationStatusRequest) Validate() error {
	return dara.Validate(s)
}
