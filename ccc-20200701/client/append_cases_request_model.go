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
	// The predictive outbound campaign ID.
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
	// The list of outbound call cases in the request body.
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
	// The agent ID of the specified agent to which the call is transferred. If this field is not empty, the system transfers the call to the specified agent. If this field is empty, the system assigns the call to an idle agent in the skill group.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The caller number. If this field is not empty, the outbound call system preferentially uses the provided number as the caller to initiate the call. If this field is empty, the system automatically selects a caller number.
	//
	// example:
	//
	// 01012345678
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The custom variables defined by the customer. The value is a JSON object that contains up to 10 properties. The name and value of each property are defined by the customer.
	//
	// example:
	//
	// {
	//
	//       "name": "customer",
	//
	//       "Customer tag": "tag"
	//
	// }
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// The masked callee number. If this field is not empty, the callee number is masked based on custom rules defined by the customer. You only need to enter the masked callee number. If a masked callee number is used, the masked number is displayed in certain scenarios, and the actual callee number cannot be viewed.
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
	// The business ID, which is the identifier in the customer\\"s business system and is used for integration scenarios.
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
