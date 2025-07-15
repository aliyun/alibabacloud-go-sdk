// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *AbortCampaignRequest
	GetCampaignId() *string
	SetInstanceId(v string) *AbortCampaignRequest
	GetInstanceId() *string
}

type AbortCampaignRequest struct {
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

func (s AbortCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortCampaignRequest) GoString() string {
	return s.String()
}

func (s *AbortCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *AbortCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AbortCampaignRequest) SetCampaignId(v string) *AbortCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCampaignRequest) SetInstanceId(v string) *AbortCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *AbortCampaignRequest) Validate() error {
	return dara.Validate(s)
}
