// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyFlowLogAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyFlowLogAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyFlowLogAttributeRequest
	GetDryRun() *bool
	SetEcrId(v string) *ModifyFlowLogAttributeRequest
	GetEcrId() *string
	SetFlowLogId(v string) *ModifyFlowLogAttributeRequest
	GetFlowLogId() *string
	SetFlowLogName(v string) *ModifyFlowLogAttributeRequest
	GetFlowLogName() *string
	SetInterval(v int32) *ModifyFlowLogAttributeRequest
	GetInterval() *int32
	SetSamplingRate(v string) *ModifyFlowLogAttributeRequest
	GetSamplingRate() *string
	SetVersion(v string) *ModifyFlowLogAttributeRequest
	GetVersion() *string
}

type ModifyFlowLogAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the token on your client. Make sure that the token is unique across different requests. The client token supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the flow log. The description can be empty or 0 to 256 characters in length.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - **true**: sends a dry run request. No EPS resource is modified.
	//
	// - **false**: sends a normal request. This is the default value.
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
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The instance ID of the flow log instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-flog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The new name of the flow log. The name must be 0 to 128 characters in length.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The duration of the capture window for the flow log to capture traffic. Unit: seconds. Valid values:
	//
	// - **60*	-
	//
	// - **600**
	//
	// Default value: **600**.
	//
	// example:
	//
	// 600
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The sampling rate. Valid values:
	//
	// - **1:4096**
	//
	// - **1:2048**
	//
	// - **1:1024**
	//
	// Default value: **1:4096**.
	//
	// example:
	//
	// 1:4096
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ModifyFlowLogAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowLogAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyFlowLogAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFlowLogAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyFlowLogAttributeRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ModifyFlowLogAttributeRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *ModifyFlowLogAttributeRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *ModifyFlowLogAttributeRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyFlowLogAttributeRequest) GetSamplingRate() *string {
	return s.SamplingRate
}

func (s *ModifyFlowLogAttributeRequest) GetVersion() *string {
	return s.Version
}

func (s *ModifyFlowLogAttributeRequest) SetClientToken(v string) *ModifyFlowLogAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDescription(v string) *ModifyFlowLogAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDryRun(v bool) *ModifyFlowLogAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetEcrId(v string) *ModifyFlowLogAttributeRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogId(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogName(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogName = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetInterval(v int32) *ModifyFlowLogAttributeRequest {
	s.Interval = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetSamplingRate(v string) *ModifyFlowLogAttributeRequest {
	s.SamplingRate = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetVersion(v string) *ModifyFlowLogAttributeRequest {
	s.Version = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) Validate() error {
	return dara.Validate(s)
}
