// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSandboxRebuildOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceEndpointId(v string) *GetSandboxRebuildOperationRequest
	GetServiceEndpointId() *string
}

type GetSandboxRebuildOperationRequest struct {
	// The service registration endpoint ID. This parameter is optional. If specified, the operation targets the runtime instance corresponding to the endpoint under the agent. If omitted, the default latest runtime identity is used. Pass the endpoint ID instead of the endpoint name. Keep this parameter consistent when querying instances, committing rebuilds, and querying nodes.
	//
	// example:
	//
	// se-7f3a92c1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
}

func (s GetSandboxRebuildOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSandboxRebuildOperationRequest) GoString() string {
	return s.String()
}

func (s *GetSandboxRebuildOperationRequest) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *GetSandboxRebuildOperationRequest) SetServiceEndpointId(v string) *GetSandboxRebuildOperationRequest {
	s.ServiceEndpointId = &v
	return s
}

func (s *GetSandboxRebuildOperationRequest) Validate() error {
	return dara.Validate(s)
}
