// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceDiscoveryEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetWorkspaceDiscoveryEndpointsRequest struct {
}

func (s GetWorkspaceDiscoveryEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceDiscoveryEndpointsRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceDiscoveryEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
