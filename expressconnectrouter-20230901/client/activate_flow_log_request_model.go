// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ActivateFlowLogRequest
	GetClientToken() *string
	SetDryRun(v bool) *ActivateFlowLogRequest
	GetDryRun() *bool
	SetEcrId(v string) *ActivateFlowLogRequest
	GetEcrId() *string
	SetFlowLogId(v string) *ActivateFlowLogRequest
	GetFlowLogId() *string
	SetVersion(v string) *ActivateFlowLogRequest
	GetVersion() *string
}

type ActivateFlowLogRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client to ensure that the value is unique among different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether it is a dry run for this request. Valid values:
	//
	// - **true**: sends a check request without modifying any EPS resource.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The Express Connect Router (ECR) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The flow log instance ID.
	//
	// > The flow log must be in the DEACTIVATE state. After you invoke the DeactivateFlowLog operation, the flow log enters the DEACTIVATING transitional state. You must invoke the DescribeFlowLogs operation for polling and wait until the state changes to DEACTIVATE before you invoke this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-flog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ActivateFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *ActivateFlowLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ActivateFlowLogRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ActivateFlowLogRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ActivateFlowLogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *ActivateFlowLogRequest) GetVersion() *string {
	return s.Version
}

func (s *ActivateFlowLogRequest) SetClientToken(v string) *ActivateFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *ActivateFlowLogRequest) SetDryRun(v bool) *ActivateFlowLogRequest {
	s.DryRun = &v
	return s
}

func (s *ActivateFlowLogRequest) SetEcrId(v string) *ActivateFlowLogRequest {
	s.EcrId = &v
	return s
}

func (s *ActivateFlowLogRequest) SetFlowLogId(v string) *ActivateFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *ActivateFlowLogRequest) SetVersion(v string) *ActivateFlowLogRequest {
	s.Version = &v
	return s
}

func (s *ActivateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
