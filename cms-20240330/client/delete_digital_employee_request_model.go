// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDigitalEmployeeRequest struct {
}

func (s DeleteDigitalEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeRequest) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeRequest) Validate() error {
	return dara.Validate(s)
}
