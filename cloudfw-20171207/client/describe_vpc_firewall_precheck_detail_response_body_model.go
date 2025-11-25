// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallPrecheckDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsFound(v bool) *DescribeVpcFirewallPrecheckDetailResponseBody
	GetIsFound() *bool
	SetPrecheckDetail(v *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) *DescribeVpcFirewallPrecheckDetailResponseBody
	GetPrecheckDetail() *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail
	SetRequestId(v string) *DescribeVpcFirewallPrecheckDetailResponseBody
	GetRequestId() *string
}

type DescribeVpcFirewallPrecheckDetailResponseBody struct {
	// example:
	//
	// false
	IsFound        *bool                                                        `json:"IsFound,omitempty" xml:"IsFound,omitempty"`
	PrecheckDetail *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail `json:"PrecheckDetail,omitempty" xml:"PrecheckDetail,omitempty" type:"Struct"`
	// example:
	//
	// 4FB718F0-CC04-5A12-B17B-188CFC3F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpcFirewallPrecheckDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPrecheckDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBody) GetIsFound() *bool {
	return s.IsFound
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBody) GetPrecheckDetail() *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail {
	return s.PrecheckDetail
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBody) SetIsFound(v bool) *DescribeVpcFirewallPrecheckDetailResponseBody {
	s.IsFound = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBody) SetPrecheckDetail(v *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) *DescribeVpcFirewallPrecheckDetailResponseBody {
	s.PrecheckDetail = v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallPrecheckDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBody) Validate() error {
	if s.PrecheckDetail != nil {
		if err := s.PrecheckDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail struct {
	// example:
	//
	// vfw-tr-7a9c8901ed394****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// example:
	//
	// vpc-m5emh0w6v2e15****
	NetworkInstanceId    *string                                                                            `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	PrecheckEntityGroups []*DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups `json:"PrecheckEntityGroups,omitempty" xml:"PrecheckEntityGroups,omitempty" type:"Repeated"`
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
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) GetPrecheckEntityGroups() []*DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	return s.PrecheckEntityGroups
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) GetPrecheckStatus() *string {
	return s.PrecheckStatus
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) GetPrecheckTimestamp() *string {
	return s.PrecheckTimestamp
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) SetFirewallId(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.FirewallId = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) SetNetworkInstanceId(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) SetPrecheckEntityGroups(v []*DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.PrecheckEntityGroups = v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) SetPrecheckStatus(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) SetPrecheckTimestamp(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.PrecheckTimestamp = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) SetRegionNo(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetail) Validate() error {
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

type DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups struct {
	// example:
	//
	// 3
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// test
	Name             *string                                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PrecheckEntities []*DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities `json:"PrecheckEntities,omitempty" xml:"PrecheckEntities,omitempty" type:"Repeated"`
	// example:
	//
	// running
	PrecheckEntityGroupStatus *string `json:"PrecheckEntityGroupStatus,omitempty" xml:"PrecheckEntityGroupStatus,omitempty"`
}

func (s DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetName() *string {
	return s.Name
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetPrecheckEntities() []*DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	return s.PrecheckEntities
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) GetPrecheckEntityGroupStatus() *string {
	return s.PrecheckEntityGroupStatus
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetFailedCount(v int32) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.FailedCount = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetName(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.Name = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetPrecheckEntities(v []*DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.PrecheckEntities = v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) SetPrecheckEntityGroupStatus(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups {
	s.PrecheckEntityGroupStatus = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroups) Validate() error {
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

type DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities struct {
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

func (s DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetInfo() *string {
	return s.Info
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetName() *string {
	return s.Name
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetInfo(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Info = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetName(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Name = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetStatus(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Status = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) SetSuggestion(v string) *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities {
	s.Suggestion = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailResponseBodyPrecheckDetailPrecheckEntityGroupsPrecheckEntities) Validate() error {
	return dara.Validate(s)
}
