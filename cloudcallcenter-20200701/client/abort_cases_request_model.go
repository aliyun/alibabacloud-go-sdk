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
	SetPhoneNumberList(v []*string) *AbortCasesRequest
	GetPhoneNumberList() []*string
}

type AbortCasesRequest struct {
	CampaignId      *string   `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneNumberList []*string `json:"PhoneNumberList,omitempty" xml:"PhoneNumberList,omitempty" type:"Repeated"`
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

func (s *AbortCasesRequest) GetPhoneNumberList() []*string {
	return s.PhoneNumberList
}

func (s *AbortCasesRequest) SetCampaignId(v string) *AbortCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCasesRequest) SetInstanceId(v string) *AbortCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *AbortCasesRequest) SetPhoneNumberList(v []*string) *AbortCasesRequest {
	s.PhoneNumberList = v
	return s
}

func (s *AbortCasesRequest) Validate() error {
	return dara.Validate(s)
}
