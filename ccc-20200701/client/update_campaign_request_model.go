// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableTime(v string) *UpdateCampaignRequest
	GetCallableTime() *string
	SetCampaignId(v string) *UpdateCampaignRequest
	GetCampaignId() *string
	SetContactFlowId(v string) *UpdateCampaignRequest
	GetContactFlowId() *string
	SetEndTime(v string) *UpdateCampaignRequest
	GetEndTime() *string
	SetInstanceId(v string) *UpdateCampaignRequest
	GetInstanceId() *string
	SetName(v string) *UpdateCampaignRequest
	GetName() *string
	SetStartTime(v string) *UpdateCampaignRequest
	GetStartTime() *string
	SetStrategyParameters(v string) *UpdateCampaignRequest
	GetStrategyParameters() *string
}

type UpdateCampaignRequest struct {
	// example:
	//
	// [
	//
	//       {
	//
	//             "beginTime": "09:00:00",
	//
	//             "endTime": "12:00:00"
	//
	//       }
	//
	// ]
	CallableTime *string `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ab11908b-6ebc-4b0c-b51e-3e17c7060
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// example:
	//
	// 3a310f56-4d30-4081-ba24-5d87a3b7262e
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// example:
	//
	// 1689933600000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1689901200000
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StrategyParameters *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
}

func (s UpdateCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCampaignRequest) GoString() string {
	return s.String()
}

func (s *UpdateCampaignRequest) GetCallableTime() *string {
	return s.CallableTime
}

func (s *UpdateCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *UpdateCampaignRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *UpdateCampaignRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateCampaignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCampaignRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCampaignRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateCampaignRequest) GetStrategyParameters() *string {
	return s.StrategyParameters
}

func (s *UpdateCampaignRequest) SetCallableTime(v string) *UpdateCampaignRequest {
	s.CallableTime = &v
	return s
}

func (s *UpdateCampaignRequest) SetCampaignId(v string) *UpdateCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *UpdateCampaignRequest) SetContactFlowId(v string) *UpdateCampaignRequest {
	s.ContactFlowId = &v
	return s
}

func (s *UpdateCampaignRequest) SetEndTime(v string) *UpdateCampaignRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateCampaignRequest) SetInstanceId(v string) *UpdateCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCampaignRequest) SetName(v string) *UpdateCampaignRequest {
	s.Name = &v
	return s
}

func (s *UpdateCampaignRequest) SetStartTime(v string) *UpdateCampaignRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateCampaignRequest) SetStrategyParameters(v string) *UpdateCampaignRequest {
	s.StrategyParameters = &v
	return s
}

func (s *UpdateCampaignRequest) Validate() error {
	return dara.Validate(s)
}
