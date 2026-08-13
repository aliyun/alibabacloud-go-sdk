// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *PauseCampaignRequest
	GetCampaignId() *string
	SetInstanceId(v string) *PauseCampaignRequest
	GetInstanceId() *string
}

type PauseCampaignRequest struct {
	// The ID of the campaign.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6ac878ab-115b-4170-a5d8-547481273364
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The instance ID of the outbound call instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PauseCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseCampaignRequest) GoString() string {
	return s.String()
}

func (s *PauseCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *PauseCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PauseCampaignRequest) SetCampaignId(v string) *PauseCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *PauseCampaignRequest) SetInstanceId(v string) *PauseCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *PauseCampaignRequest) Validate() error {
	return dara.Validate(s)
}
