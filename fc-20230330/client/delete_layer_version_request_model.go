// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayerVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteLayerVersionRequest struct {
}

func (s DeleteLayerVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayerVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayerVersionRequest) Validate() error {
	return dara.Validate(s)
}
