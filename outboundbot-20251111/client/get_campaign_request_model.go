// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *GetCampaignRequest
	GetCampaignId() *string
	SetInstanceId(v string) *GetCampaignRequest
	GetInstanceId() *string
}

type GetCampaignRequest struct {
	// The ID of the outbound task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6ac878ab-115b-4170-a5d8-547481273364
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The outbound instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 024f8cf0-c842-4c01-b74b-c8667e4579c7
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCampaignRequest) GoString() string {
	return s.String()
}

func (s *GetCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *GetCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCampaignRequest) SetCampaignId(v string) *GetCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *GetCampaignRequest) SetInstanceId(v string) *GetCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCampaignRequest) Validate() error {
	return dara.Validate(s)
}
