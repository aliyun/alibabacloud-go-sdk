// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *AddCasesShrinkRequest
	GetCampaignId() *string
	SetCaseListShrink(v string) *AddCasesShrinkRequest
	GetCaseListShrink() *string
	SetInstanceId(v string) *AddCasesShrinkRequest
	GetInstanceId() *string
}

type AddCasesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c58b9719-3bc3-441d-a4d3-fc0309ef7066
	CampaignId     *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CaseListShrink *string `json:"CaseList,omitempty" xml:"CaseList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddCasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCasesShrinkRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *AddCasesShrinkRequest) GetCaseListShrink() *string {
	return s.CaseListShrink
}

func (s *AddCasesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddCasesShrinkRequest) SetCampaignId(v string) *AddCasesShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *AddCasesShrinkRequest) SetCaseListShrink(v string) *AddCasesShrinkRequest {
	s.CaseListShrink = &v
	return s
}

func (s *AddCasesShrinkRequest) SetInstanceId(v string) *AddCasesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AddCasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
