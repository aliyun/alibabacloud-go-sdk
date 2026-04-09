// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteModelServiceRequest struct {
}

func (s DeleteModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
