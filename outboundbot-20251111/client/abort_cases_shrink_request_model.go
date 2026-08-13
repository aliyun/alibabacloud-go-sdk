// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *AbortCasesShrinkRequest
	GetCampaignId() *string
	SetInstanceId(v string) *AbortCasesShrinkRequest
	GetInstanceId() *string
	SetPhoneNumbersShrink(v string) *AbortCasesShrinkRequest
	GetPhoneNumbersShrink() *string
}

type AbortCasesShrinkRequest struct {
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
	PhoneNumbersShrink *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
}

func (s AbortCasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortCasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AbortCasesShrinkRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *AbortCasesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AbortCasesShrinkRequest) GetPhoneNumbersShrink() *string {
	return s.PhoneNumbersShrink
}

func (s *AbortCasesShrinkRequest) SetCampaignId(v string) *AbortCasesShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCasesShrinkRequest) SetInstanceId(v string) *AbortCasesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AbortCasesShrinkRequest) SetPhoneNumbersShrink(v string) *AbortCasesShrinkRequest {
	s.PhoneNumbersShrink = &v
	return s
}

func (s *AbortCasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
