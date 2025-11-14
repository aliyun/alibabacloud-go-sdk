// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCampaignNumbersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *ModifyCampaignNumbersShrinkRequest
	GetCampaignId() *string
	SetInstGroupId(v string) *ModifyCampaignNumbersShrinkRequest
	GetInstGroupId() *string
	SetInstanceId(v string) *ModifyCampaignNumbersShrinkRequest
	GetInstanceId() *string
	SetNumberListShrink(v string) *ModifyCampaignNumbersShrinkRequest
	GetNumberListShrink() *string
}

type ModifyCampaignNumbersShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6badb397-****-****-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// example:
	//
	// 3971876649-****-****-098763a382a09
	InstGroupId *string `json:"InstGroupId,omitempty" xml:"InstGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NumberListShrink *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s ModifyCampaignNumbersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCampaignNumbersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCampaignNumbersShrinkRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ModifyCampaignNumbersShrinkRequest) GetInstGroupId() *string {
	return s.InstGroupId
}

func (s *ModifyCampaignNumbersShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCampaignNumbersShrinkRequest) GetNumberListShrink() *string {
	return s.NumberListShrink
}

func (s *ModifyCampaignNumbersShrinkRequest) SetCampaignId(v string) *ModifyCampaignNumbersShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *ModifyCampaignNumbersShrinkRequest) SetInstGroupId(v string) *ModifyCampaignNumbersShrinkRequest {
	s.InstGroupId = &v
	return s
}

func (s *ModifyCampaignNumbersShrinkRequest) SetInstanceId(v string) *ModifyCampaignNumbersShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCampaignNumbersShrinkRequest) SetNumberListShrink(v string) *ModifyCampaignNumbersShrinkRequest {
	s.NumberListShrink = &v
	return s
}

func (s *ModifyCampaignNumbersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
