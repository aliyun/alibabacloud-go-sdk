// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeVulDesktopsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVulDesktopsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeVulDesktopsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVulDesktopsResponseBody
	GetTotalCount() *int32
	SetVulDesktops(v []*DescribeVulDesktopsResponseBodyVulDesktops) *DescribeVulDesktopsResponseBody
	GetVulDesktops() []*DescribeVulDesktopsResponseBodyVulDesktops
}

type DescribeVulDesktopsResponseBody struct {
	// The number of entries per page in a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. An empty value indicates that no more results exist.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kRxd1mKkNnHlUy14zdjl/I
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of cloud computers affected by the vulnerability.
	VulDesktops []*DescribeVulDesktopsResponseBodyVulDesktops `json:"VulDesktops,omitempty" xml:"VulDesktops,omitempty" type:"Repeated"`
}

func (s DescribeVulDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulDesktopsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVulDesktopsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVulDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulDesktopsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulDesktopsResponseBody) GetVulDesktops() []*DescribeVulDesktopsResponseBodyVulDesktops {
	return s.VulDesktops
}

func (s *DescribeVulDesktopsResponseBody) SetMaxResults(v int32) *DescribeVulDesktopsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeVulDesktopsResponseBody) SetNextToken(v string) *DescribeVulDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVulDesktopsResponseBody) SetRequestId(v string) *DescribeVulDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulDesktopsResponseBody) SetTotalCount(v int32) *DescribeVulDesktopsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulDesktopsResponseBody) SetVulDesktops(v []*DescribeVulDesktopsResponseBodyVulDesktops) *DescribeVulDesktopsResponseBody {
	s.VulDesktops = v
	return s
}

func (s *DescribeVulDesktopsResponseBody) Validate() error {
	if s.VulDesktops != nil {
		for _, item := range s.VulDesktops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVulDesktopsResponseBodyVulDesktops struct {
	// The configuration task ID.
	//
	// example:
	//
	// ccg-0bbay4w7bwbxd****
	ConfigGroupId *string `json:"ConfigGroupId,omitempty" xml:"ConfigGroupId,omitempty"`
	// The number of vulnerabilities.
	//
	// example:
	//
	// 60
	CveCount *int32 `json:"CveCount,omitempty" xml:"CveCount,omitempty"`
	// The list of vulnerability details.
	Cves []*DescribeVulDesktopsResponseBodyVulDesktopsCves `json:"Cves,omitempty" xml:"Cves,omitempty" type:"Repeated"`
	// The ID of the cloud computer affected by the vulnerability.
	//
	// example:
	//
	// ecd-0jtd4z5binubxe32e
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// Indicates whether the activation code is disabled.
	//
	// example:
	//
	// False
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The time when the vulnerability was first discovered.
	//
	// example:
	//
	// 2026-08-05 00:00:00
	FirstFoundTime *string `json:"FirstFoundTime,omitempty" xml:"FirstFoundTime,omitempty"`
	// The list of fix records for the cloud computer.
	FixRecords []*DescribeVulDesktopsResponseBodyVulDesktopsFixRecords `json:"FixRecords,omitempty" xml:"FixRecords,omitempty" type:"Repeated"`
	// The list of patch IDs.
	PatchIds []*string `json:"PatchIds,omitempty" xml:"PatchIds,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by WUYING Workspace.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The enterprise resource group ID.
	//
	// example:
	//
	// rg-acfm2tswogr****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The fix status of the patch.
	//
	// example:
	//
	// Fixed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The patch level.
	//
	// example:
	//
	// high
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeVulDesktopsResponseBodyVulDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDesktopsResponseBodyVulDesktops) GoString() string {
	return s.String()
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetConfigGroupId() *string {
	return s.ConfigGroupId
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetCveCount() *int32 {
	return s.CveCount
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetCves() []*DescribeVulDesktopsResponseBodyVulDesktopsCves {
	return s.Cves
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetFirstFoundTime() *string {
	return s.FirstFoundTime
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetFixRecords() []*DescribeVulDesktopsResponseBodyVulDesktopsFixRecords {
	return s.FixRecords
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetPatchIds() []*string {
	return s.PatchIds
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetStatus() *string {
	return s.Status
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) GetVulLevel() *string {
	return s.VulLevel
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetConfigGroupId(v string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.ConfigGroupId = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetCveCount(v int32) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.CveCount = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetCves(v []*DescribeVulDesktopsResponseBodyVulDesktopsCves) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.Cves = v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetDesktopId(v string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetDisabled(v bool) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.Disabled = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetFirstFoundTime(v string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.FirstFoundTime = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetFixRecords(v []*DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.FixRecords = v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetPatchIds(v []*string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.PatchIds = v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetRegionId(v string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.RegionId = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetResourceGroupId(v string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetStatus(v string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.Status = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) SetVulLevel(v string) *DescribeVulDesktopsResponseBodyVulDesktops {
	s.VulLevel = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktops) Validate() error {
	if s.Cves != nil {
		for _, item := range s.Cves {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FixRecords != nil {
		for _, item := range s.FixRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVulDesktopsResponseBodyVulDesktopsCves struct {
	// The CVE ID.
	//
	// example:
	//
	// CVE-2026-62690
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The vulnerability level.
	//
	// example:
	//
	// low
	CveLevel *string `json:"CveLevel,omitempty" xml:"CveLevel,omitempty"`
	// The vulnerability name.
	//
	// example:
	//
	// Windows Push Notification Elevation of Privilege Vulnerability
	CveTitle *string `json:"CveTitle,omitempty" xml:"CveTitle,omitempty"`
	// The CVE URL.
	//
	// example:
	//
	// https://avd.aliyun.com/detail/CVE-2026-62690
	CveUrl *string `json:"CveUrl,omitempty" xml:"CveUrl,omitempty"`
	// The vulnerability score.
	//
	// example:
	//
	// 7.0
	ImpactScore *string `json:"ImpactScore,omitempty" xml:"ImpactScore,omitempty"`
	// The reference URL.
	//
	// example:
	//
	// https://msrc.microsoft.com/update-guide/vulnerability/CVE-2026-62690
	ReferenceUrl *string `json:"ReferenceUrl,omitempty" xml:"ReferenceUrl,omitempty"`
	// The release time. The time follows the ISO 8601 standard in UTC: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-08-11 07:00:00
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
}

func (s DescribeVulDesktopsResponseBodyVulDesktopsCves) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDesktopsResponseBodyVulDesktopsCves) GoString() string {
	return s.String()
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) GetCveId() *string {
	return s.CveId
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) GetCveLevel() *string {
	return s.CveLevel
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) GetCveTitle() *string {
	return s.CveTitle
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) GetCveUrl() *string {
	return s.CveUrl
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) GetImpactScore() *string {
	return s.ImpactScore
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) GetReferenceUrl() *string {
	return s.ReferenceUrl
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) SetCveId(v string) *DescribeVulDesktopsResponseBodyVulDesktopsCves {
	s.CveId = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) SetCveLevel(v string) *DescribeVulDesktopsResponseBodyVulDesktopsCves {
	s.CveLevel = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) SetCveTitle(v string) *DescribeVulDesktopsResponseBodyVulDesktopsCves {
	s.CveTitle = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) SetCveUrl(v string) *DescribeVulDesktopsResponseBodyVulDesktopsCves {
	s.CveUrl = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) SetImpactScore(v string) *DescribeVulDesktopsResponseBodyVulDesktopsCves {
	s.ImpactScore = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) SetReferenceUrl(v string) *DescribeVulDesktopsResponseBodyVulDesktopsCves {
	s.ReferenceUrl = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) SetReleaseTime(v string) *DescribeVulDesktopsResponseBodyVulDesktopsCves {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsCves) Validate() error {
	return dara.Validate(s)
}

type DescribeVulDesktopsResponseBodyVulDesktopsFixRecords struct {
	// The batch ID of the scheduled task execution.
	//
	// example:
	//
	// d7f3d7bc-b98b-4da8-95ae-fea21b604b34
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// Update failed
	FixFailureReason *string `json:"FixFailureReason,omitempty" xml:"FixFailureReason,omitempty"`
	// The fix result.
	//
	// example:
	//
	// SUCCEED
	FixResult *string `json:"FixResult,omitempty" xml:"FixResult,omitempty"`
	// The timestamp when the fix task ended, in milliseconds.
	//
	// example:
	//
	// 2026-08-05 13:57:31
	FixTime *string `json:"FixTime,omitempty" xml:"FixTime,omitempty"`
	// The fix type.
	//
	// example:
	//
	// AutoTask
	FixType *string `json:"FixType,omitempty" xml:"FixType,omitempty"`
}

func (s DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) GoString() string {
	return s.String()
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) GetBatchId() *string {
	return s.BatchId
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) GetFixFailureReason() *string {
	return s.FixFailureReason
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) GetFixResult() *string {
	return s.FixResult
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) GetFixTime() *string {
	return s.FixTime
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) GetFixType() *string {
	return s.FixType
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) SetBatchId(v string) *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords {
	s.BatchId = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) SetFixFailureReason(v string) *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords {
	s.FixFailureReason = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) SetFixResult(v string) *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords {
	s.FixResult = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) SetFixTime(v string) *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords {
	s.FixTime = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) SetFixType(v string) *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords {
	s.FixType = &v
	return s
}

func (s *DescribeVulDesktopsResponseBodyVulDesktopsFixRecords) Validate() error {
	return dara.Validate(s)
}
