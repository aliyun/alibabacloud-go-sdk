// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendCasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *AppendCasesShrinkRequest
	GetCampaignId() *string
	SetInstanceId(v string) *AppendCasesShrinkRequest
	GetInstanceId() *string
	SetBodyShrink(v string) *AppendCasesShrinkRequest
	GetBodyShrink() *string
}

type AppendCasesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 78cf6864-9a22-4ea8-a59d-5adc2d747b0e
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendCasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AppendCasesShrinkRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *AppendCasesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppendCasesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *AppendCasesShrinkRequest) SetCampaignId(v string) *AppendCasesShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *AppendCasesShrinkRequest) SetInstanceId(v string) *AppendCasesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AppendCasesShrinkRequest) SetBodyShrink(v string) *AppendCasesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *AppendCasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
