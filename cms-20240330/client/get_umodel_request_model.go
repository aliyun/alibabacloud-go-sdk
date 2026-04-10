// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetUmodelRequest struct {
}

func (s GetUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelRequest) GoString() string {
	return s.String()
}

func (s *GetUmodelRequest) Validate() error {
	return dara.Validate(s)
}
