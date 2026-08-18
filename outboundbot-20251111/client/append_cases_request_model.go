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
	SetCases(v []*AppendCasesRequestCases) *AppendCasesRequest
	GetCases() []*AppendCasesRequestCases
	SetInstanceId(v string) *AppendCasesRequest
	GetInstanceId() *string
}

type AppendCasesRequest struct {
	// The outbound call task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40ea7fc2-c9d4-47e3-af1e-216bf7f79a44
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The list of contacts.
	//
	// This parameter is required.
	Cases []*AppendCasesRequestCases `json:"Cases,omitempty" xml:"Cases,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *AppendCasesRequest) GetCases() []*AppendCasesRequestCases {
	return s.Cases
}

func (s *AppendCasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppendCasesRequest) SetCampaignId(v string) *AppendCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *AppendCasesRequest) SetCases(v []*AppendCasesRequestCases) *AppendCasesRequest {
	s.Cases = v
	return s
}

func (s *AppendCasesRequest) SetInstanceId(v string) *AppendCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *AppendCasesRequest) Validate() error {
	if s.Cases != nil {
		for _, item := range s.Cases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppendCasesRequestCases struct {
	// Custom variables defined by the customer. The value is a JSON object that contains up to 10 properties. The name and value of each property are defined by the customer.
	//
	// example:
	//
	// ["key1":"value1"]
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 1331234****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The business system ID of the contact.
	//
	// example:
	//
	// business3-watermark-2704-1776997551
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s AppendCasesRequestCases) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesRequestCases) GoString() string {
	return s.String()
}

func (s *AppendCasesRequestCases) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *AppendCasesRequestCases) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AppendCasesRequestCases) GetPriority() *int32 {
	return s.Priority
}

func (s *AppendCasesRequestCases) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *AppendCasesRequestCases) SetCustomVariables(v string) *AppendCasesRequestCases {
	s.CustomVariables = &v
	return s
}

func (s *AppendCasesRequestCases) SetPhoneNumber(v string) *AppendCasesRequestCases {
	s.PhoneNumber = &v
	return s
}

func (s *AppendCasesRequestCases) SetPriority(v int32) *AppendCasesRequestCases {
	s.Priority = &v
	return s
}

func (s *AppendCasesRequestCases) SetReferenceId(v string) *AppendCasesRequestCases {
	s.ReferenceId = &v
	return s
}

func (s *AppendCasesRequestCases) Validate() error {
	return dara.Validate(s)
}
