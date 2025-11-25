// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallPrecheckDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsFound(v bool) *DescribeNatFirewallPrecheckDetailResponseBody
	GetIsFound() *bool
	SetPrecheckDetail(v *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) *DescribeNatFirewallPrecheckDetailResponseBody
	GetPrecheckDetail() *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail
	SetRequestId(v string) *DescribeNatFirewallPrecheckDetailResponseBody
	GetRequestId() *string
}

type DescribeNatFirewallPrecheckDetailResponseBody struct {
	// example:
	//
	// false
	IsFound        *bool                                                        `json:"IsFound,omitempty" xml:"IsFound,omitempty"`
	PrecheckDetail *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail `json:"PrecheckDetail,omitempty" xml:"PrecheckDetail,omitempty" type:"Struct"`
	// example:
	//
	// 99A65AA0-C5B5-5092-BFCF-8111B436****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatFirewallPrecheckDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPrecheckDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPrecheckDetailResponseBody) GetIsFound() *bool {
	return s.IsFound
}

func (s *DescribeNatFirewallPrecheckDetailResponseBody) GetPrecheckDetail() *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail {
	return s.PrecheckDetail
}

func (s *DescribeNatFirewallPrecheckDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallPrecheckDetailResponseBody) SetIsFound(v bool) *DescribeNatFirewallPrecheckDetailResponseBody {
	s.IsFound = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBody) SetPrecheckDetail(v *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) *DescribeNatFirewallPrecheckDetailResponseBody {
	s.PrecheckDetail = v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBody) SetRequestId(v string) *DescribeNatFirewallPrecheckDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBody) Validate() error {
	if s.PrecheckDetail != nil {
		if err := s.PrecheckDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail struct {
	// example:
	//
	// vfw-tr-7a9c8901ed394****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// example:
	//
	// vpc-m5emh0w6v2e15****
	NetworkInstanceId    *string                                                                            `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	PrecheckEntityGroups []*DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups `json:"PrecheckEntityGroups,omitempty" xml:"PrecheckEntityGroups,omitempty" type:"Repeated"`
	// example:
	//
	// failed
	PrecheckStatus *string `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	// example:
	//
	// 1715136000
	PrecheckTimestamp *string `json:"PrecheckTimestamp,omitempty" xml:"PrecheckTimestamp,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) GetPrecheckEntityGroups() []*DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	return s.PrecheckEntityGroups
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) GetPrecheckStatus() *string {
	return s.PrecheckStatus
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) GetPrecheckTimestamp() *string {
	return s.PrecheckTimestamp
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) SetFirewallId(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.FirewallId = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) SetNetworkInstanceId(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) SetPrecheckEntityGroups(v []*DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.PrecheckEntityGroups = v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) SetPrecheckStatus(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) SetPrecheckTimestamp(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.PrecheckTimestamp = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) SetRegionNo(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.RegionNo = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetail) Validate() error {
	if s.PrecheckEntityGroups != nil {
		for _, item := range s.PrecheckEntityGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups struct {
	// example:
	//
	// 9
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// test
	Name             *string                                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PrecheckEntities []*DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities `json:"PrecheckEntities,omitempty" xml:"PrecheckEntities,omitempty" type:"Repeated"`
	// example:
	//
	// running
	PrecheckEntityGroupStatus *string `json:"PrecheckEntityGroupStatus,omitempty" xml:"PrecheckEntityGroupStatus,omitempty"`
}

func (s DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetName() *string {
	return s.Name
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetPrecheckEntities() []*DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	return s.PrecheckEntities
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetPrecheckEntityGroupStatus() *string {
	return s.PrecheckEntityGroupStatus
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetFailedCount(v int32) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.FailedCount = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetName(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.Name = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetPrecheckEntities(v []*DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.PrecheckEntities = v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetPrecheckEntityGroupStatus(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.PrecheckEntityGroupStatus = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) Validate() error {
	if s.PrecheckEntities != nil {
		for _, item := range s.PrecheckEntities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities struct {
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// example:
	//
	// Precheck test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// passed
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetInfo() *string {
	return s.Info
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetName() *string {
	return s.Name
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetInfo(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Info = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetName(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Name = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetStatus(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Status = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetSuggestion(v string) *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Suggestion = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) Validate() error {
	return dara.Validate(s)
}
