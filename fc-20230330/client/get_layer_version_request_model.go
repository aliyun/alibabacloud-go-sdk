// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayerVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetLayerVersionRequest struct {
}

func (s GetLayerVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLayerVersionRequest) GoString() string {
	return s.String()
}

func (s *GetLayerVersionRequest) Validate() error {
	return dara.Validate(s)
}
