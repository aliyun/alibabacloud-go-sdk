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
	// This parameter is required.
	//
	// example:
	//
	// c58b9719-3bc3-441d-a4d3-fc0309ef7066
	CampaignId *string                    `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CaseList   []*AddCasesRequestCaseList `json:"CaseList,omitempty" xml:"CaseList,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type AddCasesRequestCaseList struct {
	Caller          *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	MaskedCallee    *string `json:"MaskedCallee,omitempty" xml:"MaskedCallee,omitempty"`
	// example:
	//
	// 188888****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
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
