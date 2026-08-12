// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudCenterInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeCloudCenterInstancesRequest
	GetRegionId() *string
	SetSdkRequest(v *DescribeCloudCenterInstancesRequestSdkRequest) *DescribeCloudCenterInstancesRequest
	GetSdkRequest() *DescribeCloudCenterInstancesRequestSdkRequest
}

type DescribeCloudCenterInstancesRequest struct {
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequest *DescribeCloudCenterInstancesRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s DescribeCloudCenterInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudCenterInstancesRequest) GetSdkRequest() *DescribeCloudCenterInstancesRequestSdkRequest {
	return s.SdkRequest
}

func (s *DescribeCloudCenterInstancesRequest) SetRegionId(v string) *DescribeCloudCenterInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetSdkRequest(v *DescribeCloudCenterInstancesRequestSdkRequest) *DescribeCloudCenterInstancesRequest {
	s.SdkRequest = v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudCenterInstancesRequestSdkRequest struct {
	// The search conditions for assets. This parameter is in JSON format. Note that the parameter names are case-sensitive.
	//
	// example:
	//
	// {\\"contactIds\\":[\\"job-658854766790086656\\",\\"job-658854801112113152\\"]}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The asset vendor. Separate multiple asset vendors with commas (,). Valid values:
	//
	// example:
	//
	// 0,10,13
	Flags *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// The importance level of the asset. Valid values:
	//
	// example:
	//
	// 2
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The language of the response. Valid values:
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The logical relationship between multiple search conditions. Valid values:
	//
	// example:
	//
	// AND
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The type of asset to query. Valid values:
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// The pagination token used to retrieve the next page of results. If no more results are available, this field is not returned.
	//
	// example:
	//
	// AAAAAZak7VOTMl2OSt/xmc4J6gbg4Z5eXuWnrvKgOsGARL76TVbKERXHXKNFurqjtfDdRw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to internationalize the default group name **Ungrouped**. Default value: **false**. Valid values:
	//
	// example:
	//
	// true
	NoGroupTrace *bool `json:"NoGroupTrace,omitempty" xml:"NoGroupTrace,omitempty"`
	// The number of assets to display per page in a paged query. Default value: 20, which indicates that 20 asset records are displayed per page.
	//
	// example:
	//
	// 99
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the member accounts in the resource folder.
	//
	// example:
	//
	// 1587359978118481
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to use the NextToken method to retrieve the vulnerability list data. If this parameter is used, TotalCount is no longer returned. Valid values:
	//
	// example:
	//
	// true
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
}

func (s DescribeCloudCenterInstancesRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetFlags() *string {
	return s.Flags
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetImportance() *int32 {
	return s.Importance
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetNoGroupTrace() *bool {
	return s.NoGroupTrace
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetResourceDirectoryAccountId() *string {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetCriteria(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetCurrentPage(v int32) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetFlags(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.Flags = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetImportance(v int32) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.Importance = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetLang(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetLogicalExp(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetMachineTypes(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetNextToken(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetNoGroupTrace(v bool) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.NoGroupTrace = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetPageSize(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetResourceDirectoryAccountId(v string) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) SetUseNextToken(v bool) *DescribeCloudCenterInstancesRequestSdkRequest {
	s.UseNextToken = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
