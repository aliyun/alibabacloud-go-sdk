// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPluginWorkspaceRequest struct {
}

func (s GetPluginWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPluginWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetPluginWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
