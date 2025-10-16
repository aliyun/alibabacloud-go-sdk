// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfficeSiteOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListOfficeSiteOverviewResponseBody
	GetNextToken() *string
	SetOfficeSiteOverviewResults(v []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) *ListOfficeSiteOverviewResponseBody
	GetOfficeSiteOverviewResults() []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults
	SetRequestId(v string) *ListOfficeSiteOverviewResponseBody
	GetRequestId() *string
}

type ListOfficeSiteOverviewResponseBody struct {
	// The token that is used to start the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network information.
	OfficeSiteOverviewResults []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults `json:"OfficeSiteOverviewResults,omitempty" xml:"OfficeSiteOverviewResults,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOfficeSiteOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOfficeSiteOverviewResponseBody) GetOfficeSiteOverviewResults() []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	return s.OfficeSiteOverviewResults
}

func (s *ListOfficeSiteOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOfficeSiteOverviewResponseBody) SetNextToken(v string) *ListOfficeSiteOverviewResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBody) SetOfficeSiteOverviewResults(v []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) *ListOfficeSiteOverviewResponseBody {
	s.OfficeSiteOverviewResults = v
	return s
}

func (s *ListOfficeSiteOverviewResponseBody) SetRequestId(v string) *ListOfficeSiteOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBody) Validate() error {
	if s.OfficeSiteOverviewResults != nil {
		for _, item := range s.OfficeSiteOverviewResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults struct {
	// The number of expired cloud computers in the office network.
	//
	// example:
	//
	// 0
	HasExpiredEdsCount *int32 `json:"HasExpiredEdsCount,omitempty" xml:"HasExpiredEdsCount,omitempty"`
	// The number of expired cloud computers in the cloud computer pool.
	//
	// example:
	//
	// 0
	HasExpiredEdsCountForGroup *int32 `json:"HasExpiredEdsCountForGroup,omitempty" xml:"HasExpiredEdsCountForGroup,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The office network status.
	//
	// Default values:
	//
	// 	- CONFIGUSERFAILED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- REGISTERING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- REGISTERED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NEEDCONFIGTRUST
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CONFIGUSERING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CONFIGTRUSTFAILED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ERROR
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CONFIGTRUSTING
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NEEDCONFIGUSER
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// REGISTERED
	OfficeSiteStatus *string `json:"OfficeSiteStatus,omitempty" xml:"OfficeSiteStatus,omitempty"`
	// The ID of the region where the office network resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of cloud computers that are running in the office network.
	//
	// example:
	//
	// 1
	RunningEdsCount *int32 `json:"RunningEdsCount,omitempty" xml:"RunningEdsCount,omitempty"`
	// The number of running cloud computers in the cloud computer pool.
	//
	// example:
	//
	// 1
	RunningEdsCountForGroup *int32 `json:"RunningEdsCountForGroup,omitempty" xml:"RunningEdsCountForGroup,omitempty"`
	// The total number of cloud computers in the office network.
	//
	// example:
	//
	// 1
	TotalEdsCount *int32 `json:"TotalEdsCount,omitempty" xml:"TotalEdsCount,omitempty"`
	// The total number of cloud computers in the cloud computer pool.
	//
	// example:
	//
	// 1
	TotalEdsCountForGroup *int32 `json:"TotalEdsCountForGroup,omitempty" xml:"TotalEdsCountForGroup,omitempty"`
	// The office network type and its suitable VPC type.
	//
	// Valid values:
	//
	// 	- standard (default): standard, exclusive VPC
	//
	// 	- customized: custom, user VPC
	//
	// 	- basic: basic, shared VPC
	//
	// example:
	//
	// standard
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	// The number of cloud computers that are about to expire in the office network.
	//
	// example:
	//
	// 0
	WillExpiredEdsCount *int32 `json:"WillExpiredEdsCount,omitempty" xml:"WillExpiredEdsCount,omitempty"`
	// The number of cloud computers that are about to expire in the cloud computer pool.
	//
	// example:
	//
	// 0
	WillExpiredEdsCountForGroup *int32 `json:"WillExpiredEdsCountForGroup,omitempty" xml:"WillExpiredEdsCountForGroup,omitempty"`
}

func (s ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) String() string {
	return dara.Prettify(s)
}

func (s ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetHasExpiredEdsCount() *int32 {
	return s.HasExpiredEdsCount
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetHasExpiredEdsCountForGroup() *int32 {
	return s.HasExpiredEdsCountForGroup
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetOfficeSiteStatus() *string {
	return s.OfficeSiteStatus
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetRunningEdsCount() *int32 {
	return s.RunningEdsCount
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetRunningEdsCountForGroup() *int32 {
	return s.RunningEdsCountForGroup
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetTotalEdsCount() *int32 {
	return s.TotalEdsCount
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetTotalEdsCountForGroup() *int32 {
	return s.TotalEdsCountForGroup
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetVpcType() *string {
	return s.VpcType
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetWillExpiredEdsCount() *int32 {
	return s.WillExpiredEdsCount
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GetWillExpiredEdsCountForGroup() *int32 {
	return s.WillExpiredEdsCountForGroup
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetHasExpiredEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.HasExpiredEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetHasExpiredEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.HasExpiredEdsCountForGroup = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteId(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteName(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteName = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteStatus(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteStatus = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRegionId(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRunningEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RunningEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRunningEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RunningEdsCountForGroup = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetTotalEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.TotalEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetTotalEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.TotalEdsCountForGroup = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetVpcType(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.VpcType = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetWillExpiredEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.WillExpiredEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetWillExpiredEdsCountForGroup(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.WillExpiredEdsCountForGroup = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) Validate() error {
	return dara.Validate(s)
}
