// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *AddCasesRequest
	GetCampaignId() *string
	SetCaseList(v []*AddCasesRequestCaseList) *AddCasesRequest
	GetCaseList() []*AddCasesRequestCaseList
	SetInstanceId(v string) *AddCasesRequest
	GetInstanceId() *string
}

type AddCasesRequest struct {
	// The predictive outbound call activity ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c58b9719-3bc3-441d-a4d3-fc0309ef7066
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The list of outbound call cases.
	CaseList []*AddCasesRequestCaseList `json:"CaseList,omitempty" xml:"CaseList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddCasesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasesRequest) GoString() string {
	return s.String()
}

func (s *AddCasesRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *AddCasesRequest) GetCaseList() []*AddCasesRequestCaseList {
	return s.CaseList
}

func (s *AddCasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddCasesRequest) SetCampaignId(v string) *AddCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *AddCasesRequest) SetCaseList(v []*AddCasesRequestCaseList) *AddCasesRequest {
	s.CaseList = v
	return s
}

func (s *AddCasesRequest) SetInstanceId(v string) *AddCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *AddCasesRequest) Validate() error {
	if s.CaseList != nil {
		for _, item := range s.CaseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddCasesRequestCaseList struct {
	// The caller number. If this field is specified, the corresponding contact will be called using this number instead of a randomly selected one.
	//
	// example:
	//
	// 185022xxxx
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// Customer-defined custom variables in JSON object format. The object can contain up to 10 properties, and both the name and value of each property are defined by the customer.
	//
	// example:
	//
	// {"name":"customer","客户标签":"tag"}
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// The desensitized callee number. If this field is not empty, it indicates that the callee number must be desensitized. The desensitization rule is defined by the customer, and only the desensitized callee number needs to be provided here. Using a desensitized callee number means that in certain scenarios, the displayed callee number will be the desensitized version, and the real callee number cannot be viewed.
	//
	// example:
	//
	// 166******66
	MaskedCallee *string `json:"MaskedCallee,omitempty" xml:"MaskedCallee,omitempty"`
	// The contact\\"s phone number.
	//
	// example:
	//
	// 188888****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The business ID, which is an identifier from the customer\\"s operational system used in integration scenarios.
	//
	// example:
	//
	// 01
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s AddCasesRequestCaseList) String() string {
	return dara.Prettify(s)
}

func (s AddCasesRequestCaseList) GoString() string {
	return s.String()
}

func (s *AddCasesRequestCaseList) GetCaller() *string {
	return s.Caller
}

func (s *AddCasesRequestCaseList) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *AddCasesRequestCaseList) GetMaskedCallee() *string {
	return s.MaskedCallee
}

func (s *AddCasesRequestCaseList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AddCasesRequestCaseList) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *AddCasesRequestCaseList) SetCaller(v string) *AddCasesRequestCaseList {
	s.Caller = &v
	return s
}

func (s *AddCasesRequestCaseList) SetCustomVariables(v string) *AddCasesRequestCaseList {
	s.CustomVariables = &v
	return s
}

func (s *AddCasesRequestCaseList) SetMaskedCallee(v string) *AddCasesRequestCaseList {
	s.MaskedCallee = &v
	return s
}

func (s *AddCasesRequestCaseList) SetPhoneNumber(v string) *AddCasesRequestCaseList {
	s.PhoneNumber = &v
	return s
}

func (s *AddCasesRequestCaseList) SetReferenceId(v string) *AddCasesRequestCaseList {
	s.ReferenceId = &v
	return s
}

func (s *AddCasesRequestCaseList) Validate() error {
	return dara.Validate(s)
}
