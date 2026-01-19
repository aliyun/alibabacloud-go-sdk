// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *ResumeCampaignRequest
	GetCampaignId() *string
	SetInstanceId(v string) *ResumeCampaignRequest
	GetInstanceId() *string
}

type ResumeCampaignRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ResumeCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeCampaignRequest) GoString() string {
	return s.String()
}

func (s *ResumeCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ResumeCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResumeCampaignRequest) SetCampaignId(v string) *ResumeCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *ResumeCampaignRequest) SetInstanceId(v string) *ResumeCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *ResumeCampaignRequest) Validate() error {
	return dara.Validate(s)
}
