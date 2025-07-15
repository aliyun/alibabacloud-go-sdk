// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeCampaignStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRealtimeCampaignStatsRequest
	GetInstanceId() *string
	SetQueueId(v string) *GetRealtimeCampaignStatsRequest
	GetQueueId() *string
}

type GetRealtimeCampaignStatsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
}

func (s GetRealtimeCampaignStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeCampaignStatsRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealtimeCampaignStatsRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *GetRealtimeCampaignStatsRequest) SetInstanceId(v string) *GetRealtimeCampaignStatsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeCampaignStatsRequest) SetQueueId(v string) *GetRealtimeCampaignStatsRequest {
	s.QueueId = &v
	return s
}

func (s *GetRealtimeCampaignStatsRequest) Validate() error {
	return dara.Validate(s)
}
