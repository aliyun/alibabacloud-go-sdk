// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCampaignNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *ModifyCampaignNumbersRequest
	GetCampaignId() *string
	SetInstGroupId(v string) *ModifyCampaignNumbersRequest
	GetInstGroupId() *string
	SetInstanceId(v string) *ModifyCampaignNumbersRequest
	GetInstanceId() *string
	SetNumberList(v []*string) *ModifyCampaignNumbersRequest
	GetNumberList() []*string
}

type ModifyCampaignNumbersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6badb397-****-****-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// example:
	//
	// 3971876649-****-****-098763a382a09
	InstGroupId *string `json:"InstGroupId,omitempty" xml:"InstGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NumberList []*string `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
}

func (s ModifyCampaignNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCampaignNumbersRequest) GoString() string {
	return s.String()
}

func (s *ModifyCampaignNumbersRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ModifyCampaignNumbersRequest) GetInstGroupId() *string {
	return s.InstGroupId
}

func (s *ModifyCampaignNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCampaignNumbersRequest) GetNumberList() []*string {
	return s.NumberList
}

func (s *ModifyCampaignNumbersRequest) SetCampaignId(v string) *ModifyCampaignNumbersRequest {
	s.CampaignId = &v
	return s
}

func (s *ModifyCampaignNumbersRequest) SetInstGroupId(v string) *ModifyCampaignNumbersRequest {
	s.InstGroupId = &v
	return s
}

func (s *ModifyCampaignNumbersRequest) SetInstanceId(v string) *ModifyCampaignNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCampaignNumbersRequest) SetNumberList(v []*string) *ModifyCampaignNumbersRequest {
	s.NumberList = v
	return s
}

func (s *ModifyCampaignNumbersRequest) Validate() error {
	return dara.Validate(s)
}
