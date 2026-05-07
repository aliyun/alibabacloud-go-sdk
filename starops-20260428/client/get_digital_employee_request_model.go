// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDigitalEmployeeRequest struct {
}

func (s GetDigitalEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeRequest) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeRequest) Validate() error {
	return dara.Validate(s)
}
