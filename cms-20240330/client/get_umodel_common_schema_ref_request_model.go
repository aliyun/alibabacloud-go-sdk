// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelCommonSchemaRefRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetUmodelCommonSchemaRefRequest struct {
}

func (s GetUmodelCommonSchemaRefRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelCommonSchemaRefRequest) GoString() string {
	return s.String()
}

func (s *GetUmodelCommonSchemaRefRequest) Validate() error {
	return dara.Validate(s)
}
