// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowlogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFlowlogRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteFlowlogRequest
	GetDryRun() *bool
	SetEcrId(v string) *DeleteFlowlogRequest
	GetEcrId() *string
	SetFlowLogId(v string) *DeleteFlowlogRequest
	GetFlowLogId() *string
	SetVersion(v string) *DeleteFlowlogRequest
	GetVersion() *string
}

type DeleteFlowlogRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client and ensure that the value is unique across different requests. The client token supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for this request. Valid values:
	//
	// - **true**: sends a dry run request without modifying EPS resources.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The flow log instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-flog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DeleteFlowlogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowlogRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFlowlogRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteFlowlogRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DeleteFlowlogRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DeleteFlowlogRequest) GetVersion() *string {
	return s.Version
}

func (s *DeleteFlowlogRequest) SetClientToken(v string) *DeleteFlowlogRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFlowlogRequest) SetDryRun(v bool) *DeleteFlowlogRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteFlowlogRequest) SetEcrId(v string) *DeleteFlowlogRequest {
	s.EcrId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetFlowLogId(v string) *DeleteFlowlogRequest {
	s.FlowLogId = &v
	return s
}

func (s *DeleteFlowlogRequest) SetVersion(v string) *DeleteFlowlogRequest {
	s.Version = &v
	return s
}

func (s *DeleteFlowlogRequest) Validate() error {
	return dara.Validate(s)
}
