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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
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
