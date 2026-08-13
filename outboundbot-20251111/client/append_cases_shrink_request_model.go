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
	SetCasesShrink(v string) *AppendCasesShrinkRequest
	GetCasesShrink() *string
	SetInstanceId(v string) *AppendCasesShrinkRequest
	GetInstanceId() *string
}

type AppendCasesShrinkRequest struct {
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
	CasesShrink *string `json:"Cases,omitempty" xml:"Cases,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *AppendCasesShrinkRequest) GetCasesShrink() *string {
	return s.CasesShrink
}

func (s *AppendCasesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppendCasesShrinkRequest) SetCampaignId(v string) *AppendCasesShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *AppendCasesShrinkRequest) SetCasesShrink(v string) *AppendCasesShrinkRequest {
	s.CasesShrink = &v
	return s
}

func (s *AppendCasesShrinkRequest) SetInstanceId(v string) *AppendCasesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AppendCasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
