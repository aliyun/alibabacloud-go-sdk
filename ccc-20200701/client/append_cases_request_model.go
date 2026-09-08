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
	// The predictive campaign ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 78cf6864-9a22-4ea8-a59d-5adc2d747b0e
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of cases to be added.
	Body []*AppendCasesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
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
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppendCasesRequestBody struct {
	// The agent ID. If you specify this parameter, the system routes the call to the specified agent. If you leave this parameter empty, the system routes the call to an idle agent in the skill group.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The caller number. If you specify this parameter, the system preferentially uses the specified number to initiate a call. If you leave this parameter empty, the system automatically selects a number to initiate a call.
	//
	// example:
	//
	// 01012345678
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// Custom variables in the format of a JSON object. The object can contain up to 10 properties, and the name and value of each property are custom.
	//
	// example:
	//
	// {
	//
	//       "name": "customer",
	//
	//       "客户标签": "tag"
	//
	// }
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// The masked callee number. If this parameter is not empty, the callee number will be masked. You can define the masking rule and specify the masked callee number. In some cases, you can only view the masked callee number instead of the real one.
	//
	// example:
	//
	// 071*****801
	MaskedCallee *string `json:"MaskedCallee,omitempty" xml:"MaskedCallee,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 188888****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The business ID, which is a custom ID from your business system, used for integration purposes.
	//
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
