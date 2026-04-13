// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetProjectRequest struct {
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
