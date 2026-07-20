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
	// example:
	//
	// cn-beijing
	RegionId   *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// {\\"contactIds\\":[\\"job-658854766790086656\\",\\"job-658854801112113152\\"]}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 0,10,13
	Flags *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// example:
	//
	// 2
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// AND
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// example:
	//
	// AAAAAZak7VOTMl2OSt/xmc4J6gbg4Z5eXuWnrvKgOsGARL76TVbKERXHXKNFurqjtfDdRw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// true
	NoGroupTrace *bool `json:"NoGroupTrace,omitempty" xml:"NoGroupTrace,omitempty"`
	// example:
	//
	// 99
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1587359978118481
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
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
