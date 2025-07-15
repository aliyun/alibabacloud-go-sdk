// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalCampaignReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *GetHistoricalCampaignReportRequest
	GetCampaignId() *string
	SetInstanceId(v string) *GetHistoricalCampaignReportRequest
	GetInstanceId() *string
}

type GetHistoricalCampaignReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4fc85829-f6fc-476e-9c0d-c350184e36f1
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHistoricalCampaignReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCampaignReportRequest) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *GetHistoricalCampaignReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHistoricalCampaignReportRequest) SetCampaignId(v string) *GetHistoricalCampaignReportRequest {
	s.CampaignId = &v
	return s
}

func (s *GetHistoricalCampaignReportRequest) SetInstanceId(v string) *GetHistoricalCampaignReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHistoricalCampaignReportRequest) Validate() error {
	return dara.Validate(s)
}
