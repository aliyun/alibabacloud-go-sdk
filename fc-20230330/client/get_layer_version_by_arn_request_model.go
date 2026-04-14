// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayerVersionByArnRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetLayerVersionByArnRequest struct {
}

func (s GetLayerVersionByArnRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLayerVersionByArnRequest) GoString() string {
	return s.String()
}

func (s *GetLayerVersionByArnRequest) Validate() error {
	return dara.Validate(s)
}
