// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateFunctionTaskRequest struct {
}

func (s CreateFunctionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionTaskRequest) Validate() error {
	return dara.Validate(s)
}
