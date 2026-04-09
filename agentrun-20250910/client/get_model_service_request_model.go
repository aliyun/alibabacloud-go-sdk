// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelServiceRequest struct {
}

func (s GetModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelServiceRequest) GoString() string {
	return s.String()
}

func (s *GetModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
