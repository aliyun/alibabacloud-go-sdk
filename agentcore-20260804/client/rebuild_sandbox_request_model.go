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
	// The service registration endpoint ID. This parameter is optional. If specified, the operation targets the runtime instance corresponding to the service registration endpoint under the agent. If omitted, the default latest runtime identity is used. Pass the service registration endpoint ID, not the service registration endpoint name. Keep this parameter consistent when querying instances, committing the rebuild, and querying nodes.
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
