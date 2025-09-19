// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutScalingConfigOutput interface {
	dara.Model
	String() string
	GoString() string
}

type PutScalingConfigOutput struct {
}

func (s PutScalingConfigOutput) String() string {
	return dara.Prettify(s)
}

func (s PutScalingConfigOutput) GoString() string {
	return s.String()
}

func (s *PutScalingConfigOutput) Validate() error {
	return dara.Validate(s)
}
