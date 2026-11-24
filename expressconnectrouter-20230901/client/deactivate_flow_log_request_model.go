// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeactivateFlowLogRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeactivateFlowLogRequest
	GetDryRun() *bool
	SetEcrId(v string) *DeactivateFlowLogRequest
	GetEcrId() *string
	SetFlowLogId(v string) *DeactivateFlowLogRequest
	GetFlowLogId() *string
	SetVersion(v string) *DeactivateFlowLogRequest
	GetVersion() *string
}

type DeactivateFlowLogRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that the value is unique among different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether this is a dry run request. Valid values:
	//
	// - **true**: Sends a check request and does not modify the resources.
	//
	// - **false**: Sends a normal request. Default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Express Connect Router (ECR) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the flow log instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-flog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DeactivateFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactivateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *DeactivateFlowLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeactivateFlowLogRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeactivateFlowLogRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DeactivateFlowLogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DeactivateFlowLogRequest) GetVersion() *string {
	return s.Version
}

func (s *DeactivateFlowLogRequest) SetClientToken(v string) *DeactivateFlowLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeactivateFlowLogRequest) SetDryRun(v bool) *DeactivateFlowLogRequest {
	s.DryRun = &v
	return s
}

func (s *DeactivateFlowLogRequest) SetEcrId(v string) *DeactivateFlowLogRequest {
	s.EcrId = &v
	return s
}

func (s *DeactivateFlowLogRequest) SetFlowLogId(v string) *DeactivateFlowLogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeactivateFlowLogRequest) SetVersion(v string) *DeactivateFlowLogRequest {
	s.Version = &v
	return s
}

func (s *DeactivateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
