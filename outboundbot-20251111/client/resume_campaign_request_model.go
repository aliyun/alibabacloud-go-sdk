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
	// 活动ID
	//
	// This parameter is required.
	//
	// example:
	//
	// dcee2bca-0fa0-4826-89b8-1f693574023b
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// 外呼实例ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 33040b9a-b04b-452f-b554-cd6f3a15f850
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
