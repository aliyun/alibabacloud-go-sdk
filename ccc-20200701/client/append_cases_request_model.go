// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendCasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *AppendCasesRequest
	GetCampaignId() *string
	SetInstanceId(v string) *AppendCasesRequest
	GetInstanceId() *string
	SetBody(v []*AppendCasesRequestBody) *AppendCasesRequest
	GetBody() []*AppendCasesRequestBody
}

type AppendCasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 78cf6864-9a22-4ea8-a59d-5adc2d747b0e
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Body       []*AppendCasesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s AppendCasesRequest) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesRequest) GoString() string {
	return s.String()
}

func (s *AppendCasesRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *AppendCasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppendCasesRequest) GetBody() []*AppendCasesRequestBody {
	return s.Body
}

func (s *AppendCasesRequest) SetCampaignId(v string) *AppendCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *AppendCasesRequest) SetInstanceId(v string) *AppendCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *AppendCasesRequest) SetBody(v []*AppendCasesRequestBody) *AppendCasesRequest {
	s.Body = v
	return s
}

func (s *AppendCasesRequest) Validate() error {
	return dara.Validate(s)
}

type AppendCasesRequestBody struct {
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 01012345678
	Caller          *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// example:
	//
	// 071*****801
	MaskedCallee *string `json:"MaskedCallee,omitempty" xml:"MaskedCallee,omitempty"`
	// example:
	//
	// 188888****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 01
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s AppendCasesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesRequestBody) GoString() string {
	return s.String()
}

func (s *AppendCasesRequestBody) GetAgentId() *string {
	return s.AgentId
}

func (s *AppendCasesRequestBody) GetCaller() *string {
	return s.Caller
}

func (s *AppendCasesRequestBody) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *AppendCasesRequestBody) GetMaskedCallee() *string {
	return s.MaskedCallee
}

func (s *AppendCasesRequestBody) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AppendCasesRequestBody) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *AppendCasesRequestBody) SetAgentId(v string) *AppendCasesRequestBody {
	s.AgentId = &v
	return s
}

func (s *AppendCasesRequestBody) SetCaller(v string) *AppendCasesRequestBody {
	s.Caller = &v
	return s
}

func (s *AppendCasesRequestBody) SetCustomVariables(v string) *AppendCasesRequestBody {
	s.CustomVariables = &v
	return s
}

func (s *AppendCasesRequestBody) SetMaskedCallee(v string) *AppendCasesRequestBody {
	s.MaskedCallee = &v
	return s
}

func (s *AppendCasesRequestBody) SetPhoneNumber(v string) *AppendCasesRequestBody {
	s.PhoneNumber = &v
	return s
}

func (s *AppendCasesRequestBody) SetReferenceId(v string) *AppendCasesRequestBody {
	s.ReferenceId = &v
	return s
}

func (s *AppendCasesRequestBody) Validate() error {
	return dara.Validate(s)
}
