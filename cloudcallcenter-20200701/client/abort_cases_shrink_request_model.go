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
	SetPhoneNumberListShrink(v string) *AbortCasesShrinkRequest
	GetPhoneNumberListShrink() *string
}

type AbortCasesShrinkRequest struct {
	CampaignId            *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneNumberListShrink *string `json:"PhoneNumberList,omitempty" xml:"PhoneNumberList,omitempty"`
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

func (s *AbortCasesShrinkRequest) GetPhoneNumberListShrink() *string {
	return s.PhoneNumberListShrink
}

func (s *AbortCasesShrinkRequest) SetCampaignId(v string) *AbortCasesShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCasesShrinkRequest) SetInstanceId(v string) *AbortCasesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AbortCasesShrinkRequest) SetPhoneNumberListShrink(v string) *AbortCasesShrinkRequest {
	s.PhoneNumberListShrink = &v
	return s
}

func (s *AbortCasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
