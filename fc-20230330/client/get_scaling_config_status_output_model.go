// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScalingConfigStatusOutput interface {
	dara.Model
	String() string
	GoString() string
}

type GetScalingConfigStatusOutput struct {
}

func (s GetScalingConfigStatusOutput) String() string {
	return dara.Prettify(s)
}

func (s GetScalingConfigStatusOutput) GoString() string {
	return s.String()
}

func (s *GetScalingConfigStatusOutput) Validate() error {
	return dara.Validate(s)
}
