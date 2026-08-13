// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *AbortCasesRequest
	GetCampaignId() *string
	SetInstanceId(v string) *AbortCasesRequest
	GetInstanceId() *string
	SetPhoneNumbers(v []*string) *AbortCasesRequest
	GetPhoneNumbers() []*string
}

type AbortCasesRequest struct {
	// The ID of the predictive outbound call campaign.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2180afb0-83a9-4a13-9f19-467d63041dbf
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1dcb09c5-d5db-4397-bf65-db854463beea
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of contact phone numbers to cancel.
	PhoneNumbers []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
}

func (s AbortCasesRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortCasesRequest) GoString() string {
	return s.String()
}

func (s *AbortCasesRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *AbortCasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AbortCasesRequest) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *AbortCasesRequest) SetCampaignId(v string) *AbortCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCasesRequest) SetInstanceId(v string) *AbortCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *AbortCasesRequest) SetPhoneNumbers(v []*string) *AbortCasesRequest {
	s.PhoneNumbers = v
	return s
}

func (s *AbortCasesRequest) Validate() error {
	return dara.Validate(s)
}
