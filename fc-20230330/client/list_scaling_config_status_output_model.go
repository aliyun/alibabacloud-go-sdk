// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScalingConfigStatusOutput interface {
	dara.Model
	String() string
	GoString() string
}

type ListScalingConfigStatusOutput struct {
}

func (s ListScalingConfigStatusOutput) String() string {
	return dara.Prettify(s)
}

func (s ListScalingConfigStatusOutput) GoString() string {
	return s.String()
}

func (s *ListScalingConfigStatusOutput) Validate() error {
	return dara.Validate(s)
}
