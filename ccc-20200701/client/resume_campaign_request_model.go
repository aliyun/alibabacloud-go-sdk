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
