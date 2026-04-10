// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceObservabilityRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateServiceObservabilityRequest struct {
}

func (s CreateServiceObservabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceObservabilityRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceObservabilityRequest) Validate() error {
	return dara.Validate(s)
}
