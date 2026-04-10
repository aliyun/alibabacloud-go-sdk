// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceObservabilityRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetServiceObservabilityRequest struct {
}

func (s GetServiceObservabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceObservabilityRequest) GoString() string {
	return s.String()
}

func (s *GetServiceObservabilityRequest) Validate() error {
	return dara.Validate(s)
}
