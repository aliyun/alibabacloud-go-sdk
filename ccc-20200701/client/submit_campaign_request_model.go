// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *SubmitCampaignRequest
	GetCampaignId() *string
	SetInstanceId(v string) *SubmitCampaignRequest
	GetInstanceId() *string
}

type SubmitCampaignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6badb397-a8b5-40b6-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SubmitCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCampaignRequest) GoString() string {
	return s.String()
}

func (s *SubmitCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *SubmitCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SubmitCampaignRequest) SetCampaignId(v string) *SubmitCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *SubmitCampaignRequest) SetInstanceId(v string) *SubmitCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *SubmitCampaignRequest) Validate() error {
	return dara.Validate(s)
}
