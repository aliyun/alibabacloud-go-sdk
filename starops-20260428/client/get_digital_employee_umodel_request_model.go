// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDigitalEmployeeUmodelRequest struct {
}

func (s GetDigitalEmployeeUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeUmodelRequest) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeUmodelRequest) Validate() error {
	return dara.Validate(s)
}
