// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetServiceRequest struct {
}

func (s GetServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) Validate() error {
	return dara.Validate(s)
}
