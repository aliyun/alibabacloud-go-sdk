// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRiskCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateRiskCheckTaskRequest struct {
}

func (s CreateRiskCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRiskCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRiskCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
