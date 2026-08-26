// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteWorkspaceRequest struct {
}

func (s DeleteWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
