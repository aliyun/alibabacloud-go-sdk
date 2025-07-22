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
}

type ActivateFlowLogRequest struct {
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
	// true
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

func (s *ActivateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
