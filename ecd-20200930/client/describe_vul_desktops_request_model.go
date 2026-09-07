// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCveId(v string) *DescribeVulDesktopsRequest
	GetCveId() *string
	SetDesktopIdList(v []*string) *DescribeVulDesktopsRequest
	GetDesktopIdList() []*string
	SetIncludeFixResult(v bool) *DescribeVulDesktopsRequest
	GetIncludeFixResult() *bool
	SetLanguage(v string) *DescribeVulDesktopsRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeVulDesktopsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVulDesktopsRequest
	GetNextToken() *string
	SetOnlyCurrentMonthFixAttempted(v bool) *DescribeVulDesktopsRequest
	GetOnlyCurrentMonthFixAttempted() *bool
	SetPageNumber(v int32) *DescribeVulDesktopsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVulDesktopsRequest
	GetPageSize() *int32
	SetPatchId(v string) *DescribeVulDesktopsRequest
	GetPatchId() *string
	SetRegionId(v string) *DescribeVulDesktopsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVulDesktopsRequest
	GetResourceGroupId() *string
	SetSearchRegionId(v string) *DescribeVulDesktopsRequest
	GetSearchRegionId() *string
	SetStatusList(v []*string) *DescribeVulDesktopsRequest
	GetStatusList() []*string
	SetVulLevel(v string) *DescribeVulDesktopsRequest
	GetVulLevel() *string
}

type DescribeVulDesktopsRequest struct {
	// The CVE ID.
	//
	// example:
	//
	// CVE-2026-43284
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The list of cloud computer IDs.
	DesktopIdList []*string `json:"DesktopIdList,omitempty" xml:"DesktopIdList,omitempty" type:"Repeated"`
	// Specifies whether to include patch update results.
	//
	// example:
	//
	// false
	IncludeFixResult *bool `json:"IncludeFixResult,omitempty" xml:"IncludeFixResult,omitempty"`
	// The language type of the returned information.
	//
	// example:
	//
	// ch
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of entries per page in a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to include only cloud computers on which fix tasks were executed in the current month.
	//
	// example:
	//
	// false
	OnlyCurrentMonthFixAttempted *bool `json:"OnlyCurrentMonthFixAttempted,omitempty" xml:"OnlyCurrentMonthFixAttempted,omitempty"`
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The patch ID.
	//
	// example:
	//
	// KB5082063
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by WUYING Workspace.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-f3s3dgt8dtb0vlqc8
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The region ID used to filter cloud computer information for a specific region.
	//
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// The list of vulnerability status details.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The security level of the intrusion prevention event. Valid values:
	//
	// - **low**: Low risk.
	//
	// - **medium**: Medium risk.
	//
	// - **critical**: High risk.
	//
	// > If you do not set this parameter, vulnerabilities of all security levels are queried.
	//
	// example:
	//
	// low
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeVulDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulDesktopsRequest) GetCveId() *string {
	return s.CveId
}

func (s *DescribeVulDesktopsRequest) GetDesktopIdList() []*string {
	return s.DesktopIdList
}

func (s *DescribeVulDesktopsRequest) GetIncludeFixResult() *bool {
	return s.IncludeFixResult
}

func (s *DescribeVulDesktopsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeVulDesktopsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVulDesktopsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVulDesktopsRequest) GetOnlyCurrentMonthFixAttempted() *bool {
	return s.OnlyCurrentMonthFixAttempted
}

func (s *DescribeVulDesktopsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVulDesktopsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVulDesktopsRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *DescribeVulDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVulDesktopsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVulDesktopsRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeVulDesktopsRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *DescribeVulDesktopsRequest) GetVulLevel() *string {
	return s.VulLevel
}

func (s *DescribeVulDesktopsRequest) SetCveId(v string) *DescribeVulDesktopsRequest {
	s.CveId = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetDesktopIdList(v []*string) *DescribeVulDesktopsRequest {
	s.DesktopIdList = v
	return s
}

func (s *DescribeVulDesktopsRequest) SetIncludeFixResult(v bool) *DescribeVulDesktopsRequest {
	s.IncludeFixResult = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetLanguage(v string) *DescribeVulDesktopsRequest {
	s.Language = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetMaxResults(v int32) *DescribeVulDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetNextToken(v string) *DescribeVulDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetOnlyCurrentMonthFixAttempted(v bool) *DescribeVulDesktopsRequest {
	s.OnlyCurrentMonthFixAttempted = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetPageNumber(v int32) *DescribeVulDesktopsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetPageSize(v int32) *DescribeVulDesktopsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetPatchId(v string) *DescribeVulDesktopsRequest {
	s.PatchId = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetRegionId(v string) *DescribeVulDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetResourceGroupId(v string) *DescribeVulDesktopsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetSearchRegionId(v string) *DescribeVulDesktopsRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeVulDesktopsRequest) SetStatusList(v []*string) *DescribeVulDesktopsRequest {
	s.StatusList = v
	return s
}

func (s *DescribeVulDesktopsRequest) SetVulLevel(v string) *DescribeVulDesktopsRequest {
	s.VulLevel = &v
	return s
}

func (s *DescribeVulDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
