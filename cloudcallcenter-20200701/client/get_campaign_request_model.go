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
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
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
