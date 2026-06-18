// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBranchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranches(v *ListBranchesResponseBodyBranches) *ListBranchesResponseBody
	GetBranches() *ListBranchesResponseBodyBranches
	SetMaxResults(v int32) *ListBranchesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListBranchesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListBranchesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBranchesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBranchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBranchesResponseBody
	GetTotalCount() *int32
}

type ListBranchesResponseBody struct {
	Branches *ListBranchesResponseBodyBranches `json:"Branches,omitempty" xml:"Branches,omitempty" type:"Struct"`
	// The maximum number of records returned in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. It is not required for the first query. For subsequent queries, use the NextToken returned from the previous query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// Valid values:
	//
	// - 10
	//
	// - 20
	//
	// - 50
	//
	// - 100
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of branches that match the query conditions.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBranchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBranchesResponseBody) GetBranches() *ListBranchesResponseBodyBranches {
	return s.Branches
}

func (s *ListBranchesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBranchesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBranchesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBranchesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBranchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBranchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBranchesResponseBody) SetBranches(v *ListBranchesResponseBodyBranches) *ListBranchesResponseBody {
	s.Branches = v
	return s
}

func (s *ListBranchesResponseBody) SetMaxResults(v int32) *ListBranchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBranchesResponseBody) SetNextToken(v string) *ListBranchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBranchesResponseBody) SetPageNumber(v int32) *ListBranchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBranchesResponseBody) SetPageSize(v int32) *ListBranchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBranchesResponseBody) SetRequestId(v string) *ListBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBranchesResponseBody) SetTotalCount(v int32) *ListBranchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBranchesResponseBody) Validate() error {
	if s.Branches != nil {
		if err := s.Branches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBranchesResponseBodyBranches struct {
	Branch []*ListBranchesResponseBodyBranchesBranch `json:"Branch,omitempty" xml:"Branch,omitempty" type:"Repeated"`
}

func (s ListBranchesResponseBodyBranches) String() string {
	return dara.Prettify(s)
}

func (s ListBranchesResponseBodyBranches) GoString() string {
	return s.String()
}

func (s *ListBranchesResponseBodyBranches) GetBranch() []*ListBranchesResponseBodyBranchesBranch {
	return s.Branch
}

func (s *ListBranchesResponseBodyBranches) SetBranch(v []*ListBranchesResponseBodyBranchesBranch) *ListBranchesResponseBodyBranches {
	s.Branch = v
	return s
}

func (s *ListBranchesResponseBodyBranches) Validate() error {
	if s.Branch != nil {
		for _, item := range s.Branch {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBranchesResponseBodyBranchesBranch struct {
	BranchId         *string                                     `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	BranchName       *string                                     `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	CreateTime       *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiresAt        *string                                     `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	InitSource       *string                                     `json:"InitSource,omitempty" xml:"InitSource,omitempty"`
	IsDefault        *bool                                       `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ParentBranchId   *string                                     `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	ParentBranchName *string                                     `json:"ParentBranchName,omitempty" xml:"ParentBranchName,omitempty"`
	ParentLSN        *string                                     `json:"ParentLSN,omitempty" xml:"ParentLSN,omitempty"`
	ParentTimestamp  *string                                     `json:"ParentTimestamp,omitempty" xml:"ParentTimestamp,omitempty"`
	ProjectId        *string                                     `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Protected        *bool                                       `json:"Protected,omitempty" xml:"Protected,omitempty"`
	ServiceType      *string                                     `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status           *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags             *ListBranchesResponseBodyBranchesBranchTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListBranchesResponseBodyBranchesBranch) String() string {
	return dara.Prettify(s)
}

func (s ListBranchesResponseBodyBranchesBranch) GoString() string {
	return s.String()
}

func (s *ListBranchesResponseBodyBranchesBranch) GetBranchId() *string {
	return s.BranchId
}

func (s *ListBranchesResponseBodyBranchesBranch) GetBranchName() *string {
	return s.BranchName
}

func (s *ListBranchesResponseBodyBranchesBranch) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBranchesResponseBodyBranchesBranch) GetDescription() *string {
	return s.Description
}

func (s *ListBranchesResponseBodyBranchesBranch) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *ListBranchesResponseBodyBranchesBranch) GetInitSource() *string {
	return s.InitSource
}

func (s *ListBranchesResponseBodyBranchesBranch) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListBranchesResponseBodyBranchesBranch) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *ListBranchesResponseBodyBranchesBranch) GetParentBranchName() *string {
	return s.ParentBranchName
}

func (s *ListBranchesResponseBodyBranchesBranch) GetParentLSN() *string {
	return s.ParentLSN
}

func (s *ListBranchesResponseBodyBranchesBranch) GetParentTimestamp() *string {
	return s.ParentTimestamp
}

func (s *ListBranchesResponseBodyBranchesBranch) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListBranchesResponseBodyBranchesBranch) GetProtected() *bool {
	return s.Protected
}

func (s *ListBranchesResponseBodyBranchesBranch) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListBranchesResponseBodyBranchesBranch) GetStatus() *string {
	return s.Status
}

func (s *ListBranchesResponseBodyBranchesBranch) GetTags() *ListBranchesResponseBodyBranchesBranchTags {
	return s.Tags
}

func (s *ListBranchesResponseBodyBranchesBranch) SetBranchId(v string) *ListBranchesResponseBodyBranchesBranch {
	s.BranchId = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetBranchName(v string) *ListBranchesResponseBodyBranchesBranch {
	s.BranchName = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetCreateTime(v string) *ListBranchesResponseBodyBranchesBranch {
	s.CreateTime = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetDescription(v string) *ListBranchesResponseBodyBranchesBranch {
	s.Description = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetExpiresAt(v string) *ListBranchesResponseBodyBranchesBranch {
	s.ExpiresAt = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetInitSource(v string) *ListBranchesResponseBodyBranchesBranch {
	s.InitSource = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetIsDefault(v bool) *ListBranchesResponseBodyBranchesBranch {
	s.IsDefault = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetParentBranchId(v string) *ListBranchesResponseBodyBranchesBranch {
	s.ParentBranchId = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetParentBranchName(v string) *ListBranchesResponseBodyBranchesBranch {
	s.ParentBranchName = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetParentLSN(v string) *ListBranchesResponseBodyBranchesBranch {
	s.ParentLSN = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetParentTimestamp(v string) *ListBranchesResponseBodyBranchesBranch {
	s.ParentTimestamp = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetProjectId(v string) *ListBranchesResponseBodyBranchesBranch {
	s.ProjectId = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetProtected(v bool) *ListBranchesResponseBodyBranchesBranch {
	s.Protected = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetServiceType(v string) *ListBranchesResponseBodyBranchesBranch {
	s.ServiceType = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetStatus(v string) *ListBranchesResponseBodyBranchesBranch {
	s.Status = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) SetTags(v *ListBranchesResponseBodyBranchesBranchTags) *ListBranchesResponseBodyBranchesBranch {
	s.Tags = v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranch) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBranchesResponseBodyBranchesBranchTags struct {
	Tag []*ListBranchesResponseBodyBranchesBranchTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListBranchesResponseBodyBranchesBranchTags) String() string {
	return dara.Prettify(s)
}

func (s ListBranchesResponseBodyBranchesBranchTags) GoString() string {
	return s.String()
}

func (s *ListBranchesResponseBodyBranchesBranchTags) GetTag() []*ListBranchesResponseBodyBranchesBranchTagsTag {
	return s.Tag
}

func (s *ListBranchesResponseBodyBranchesBranchTags) SetTag(v []*ListBranchesResponseBodyBranchesBranchTagsTag) *ListBranchesResponseBodyBranchesBranchTags {
	s.Tag = v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranchTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBranchesResponseBodyBranchesBranchTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBranchesResponseBodyBranchesBranchTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListBranchesResponseBodyBranchesBranchTagsTag) GoString() string {
	return s.String()
}

func (s *ListBranchesResponseBodyBranchesBranchTagsTag) GetKey() *string {
	return s.Key
}

func (s *ListBranchesResponseBodyBranchesBranchTagsTag) GetValue() *string {
	return s.Value
}

func (s *ListBranchesResponseBodyBranchesBranchTagsTag) SetKey(v string) *ListBranchesResponseBodyBranchesBranchTagsTag {
	s.Key = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranchTagsTag) SetValue(v string) *ListBranchesResponseBodyBranchesBranchTagsTag {
	s.Value = &v
	return s
}

func (s *ListBranchesResponseBodyBranchesBranchTagsTag) Validate() error {
	return dara.Validate(s)
}
