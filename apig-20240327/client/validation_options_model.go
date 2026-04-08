// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidationOptions interface {
	dara.Model
	String() string
	GoString() string
}

type ValidationOptions struct {
}

func (s ValidationOptions) String() string {
	return dara.Prettify(s)
}

func (s ValidationOptions) GoString() string {
	return s.String()
}

func (s *ValidationOptions) Validate() error {
	return dara.Validate(s)
}
