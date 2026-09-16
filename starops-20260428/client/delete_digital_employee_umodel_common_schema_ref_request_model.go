// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeUmodelCommonSchemaRefRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDigitalEmployeeUmodelCommonSchemaRefRequest struct {
}

func (s DeleteDigitalEmployeeUmodelCommonSchemaRefRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeUmodelCommonSchemaRefRequest) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefRequest) Validate() error {
	return dara.Validate(s)
}
