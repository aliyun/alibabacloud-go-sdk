// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetWorkspaceRequest struct {
}

func (s GetWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
