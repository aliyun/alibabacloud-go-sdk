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
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// f295a472-35ee-442a-9013-b13862505a1a
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The outbound instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b3dbfb82-1ae6-4e73-b717-f494727d2af3
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
