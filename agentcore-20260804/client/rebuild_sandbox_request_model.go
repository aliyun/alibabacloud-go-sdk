// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceEndpointId(v string) *RebuildSandboxRequest
	GetServiceEndpointId() *string
}

type RebuildSandboxRequest struct {
	// The ID of the service registration endpoint. This parameter is optional. If you specify this parameter, the operation targets the runtime instance that corresponds to the service registration endpoint under the agent. If you omit this parameter, the default latest runtime identity is used. Pass the service registration endpoint ID instead of the service registration endpoint name. You must keep this parameter consistent when you query instances, commit rebuilds, and query nodes.
	//
	// example:
	//
	// se-7f3a92c1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
}

func (s RebuildSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s RebuildSandboxRequest) GoString() string {
	return s.String()
}

func (s *RebuildSandboxRequest) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *RebuildSandboxRequest) SetServiceEndpointId(v string) *RebuildSandboxRequest {
	s.ServiceEndpointId = &v
	return s
}

func (s *RebuildSandboxRequest) Validate() error {
	return dara.Validate(s)
}
