// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v *DescribeBranchResponseBodyBranch) *DescribeBranchResponseBody
	GetBranch() *DescribeBranchResponseBodyBranch
	SetRequestId(v string) *DescribeBranchResponseBody
	GetRequestId() *string
}

type DescribeBranchResponseBody struct {
	// The branch list. Each element represents a Supabase branch.
	Branch *DescribeBranchResponseBodyBranch `json:"Branch,omitempty" xml:"Branch,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBranchResponseBody) GetBranch() *DescribeBranchResponseBodyBranch {
	return s.Branch
}

func (s *DescribeBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBranchResponseBody) SetBranch(v *DescribeBranchResponseBodyBranch) *DescribeBranchResponseBody {
	s.Branch = v
	return s
}

func (s *DescribeBranchResponseBody) SetRequestId(v string) *DescribeBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBranchResponseBody) Validate() error {
	if s.Branch != nil {
		if err := s.Branch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBranchResponseBodyBranch struct {
	// The branch ID, which uniquely identifies a Supabase branch.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The branch name.
	//
	// example:
	//
	// dev
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The connection information of the compute node associated with the branch.
	//
	// example:
	//
	// postgresql://user:password@host:5432/db
	ComputeEndpoint *string `json:"ComputeEndpoint,omitempty" xml:"ComputeEndpoint,omitempty"`
	// The time when the branch was created, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-04-08T09:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The branch description.
	//
	// example:
	//
	// test branch
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the branch expires and is automatically deleted, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-10-08T09:11:12Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The initialization source of the branch.
	//
	// Valid values:
	//
	// - ParentData: Copies the schema and data from the parent branch. This is the default value.
	//
	// - SchemaOnly: Copies only the schema structure.
	//
	// example:
	//
	// ParentData
	InitSource *string `json:"InitSource,omitempty" xml:"InitSource,omitempty"`
	// Indicates whether this is the default branch.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The parent branch ID, which specifies the parent branch of a new branch or a query condition.
	//
	// example:
	//
	// br-main
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// The parent branch name. This value is empty or displayed as - for the primary branch.
	//
	// example:
	//
	// main
	ParentBranchName *string `json:"ParentBranchName,omitempty" xml:"ParentBranchName,omitempty"`
	// The Log Sequence Number (LSN) of the parent branch at the time this branch was created.
	//
	// example:
	//
	// 0/3522648
	ParentLSN *string `json:"ParentLSN,omitempty" xml:"ParentLSN,omitempty"`
	// The data synchronization point in time selected from the parent branch when this branch was created, in ISO 8601 UTC format.
	//
	// Note:
	//
	// - For child branches, this value indicates the point in time of the parent branch selected during creation.
	//
	// - If no parent branch exists, the value 1970-01-01T00:00:00.000Z is returned.
	//
	// example:
	//
	// 2026-04-08T09:11:12Z
	ParentTimestamp *string `json:"ParentTimestamp,omitempty" xml:"ParentTimestamp,omitempty"`
	// The Supabase project ID that corresponds to the primary branch.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Indicates whether branch protection is enabled. A value of true indicates that branch protection is enabled. A value of false indicates that branch protection is disabled.
	//
	// example:
	//
	// false
	Protected *bool `json:"Protected,omitempty" xml:"Protected,omitempty"`
	// The service type.
	//
	// Valid values:
	//
	// - Supabase: Supabase service.
	//
	// - Memory: Memory service.
	//
	// example:
	//
	// Supabase
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The branch status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of branch tags.
	Tags []*DescribeBranchResponseBodyBranchTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeBranchResponseBodyBranch) String() string {
	return dara.Prettify(s)
}

func (s DescribeBranchResponseBodyBranch) GoString() string {
	return s.String()
}

func (s *DescribeBranchResponseBodyBranch) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeBranchResponseBodyBranch) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeBranchResponseBodyBranch) GetComputeEndpoint() *string {
	return s.ComputeEndpoint
}

func (s *DescribeBranchResponseBodyBranch) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeBranchResponseBodyBranch) GetDescription() *string {
	return s.Description
}

func (s *DescribeBranchResponseBodyBranch) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *DescribeBranchResponseBodyBranch) GetInitSource() *string {
	return s.InitSource
}

func (s *DescribeBranchResponseBodyBranch) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeBranchResponseBodyBranch) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *DescribeBranchResponseBodyBranch) GetParentBranchName() *string {
	return s.ParentBranchName
}

func (s *DescribeBranchResponseBodyBranch) GetParentLSN() *string {
	return s.ParentLSN
}

func (s *DescribeBranchResponseBodyBranch) GetParentTimestamp() *string {
	return s.ParentTimestamp
}

func (s *DescribeBranchResponseBodyBranch) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeBranchResponseBodyBranch) GetProtected() *bool {
	return s.Protected
}

func (s *DescribeBranchResponseBodyBranch) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeBranchResponseBodyBranch) GetStatus() *string {
	return s.Status
}

func (s *DescribeBranchResponseBodyBranch) GetTags() []*DescribeBranchResponseBodyBranchTags {
	return s.Tags
}

func (s *DescribeBranchResponseBodyBranch) SetBranchId(v string) *DescribeBranchResponseBodyBranch {
	s.BranchId = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetBranchName(v string) *DescribeBranchResponseBodyBranch {
	s.BranchName = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetComputeEndpoint(v string) *DescribeBranchResponseBodyBranch {
	s.ComputeEndpoint = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetCreateTime(v string) *DescribeBranchResponseBodyBranch {
	s.CreateTime = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetDescription(v string) *DescribeBranchResponseBodyBranch {
	s.Description = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetExpiresAt(v string) *DescribeBranchResponseBodyBranch {
	s.ExpiresAt = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetInitSource(v string) *DescribeBranchResponseBodyBranch {
	s.InitSource = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetIsDefault(v bool) *DescribeBranchResponseBodyBranch {
	s.IsDefault = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetParentBranchId(v string) *DescribeBranchResponseBodyBranch {
	s.ParentBranchId = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetParentBranchName(v string) *DescribeBranchResponseBodyBranch {
	s.ParentBranchName = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetParentLSN(v string) *DescribeBranchResponseBodyBranch {
	s.ParentLSN = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetParentTimestamp(v string) *DescribeBranchResponseBodyBranch {
	s.ParentTimestamp = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetProjectId(v string) *DescribeBranchResponseBodyBranch {
	s.ProjectId = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetProtected(v bool) *DescribeBranchResponseBodyBranch {
	s.Protected = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetServiceType(v string) *DescribeBranchResponseBodyBranch {
	s.ServiceType = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetStatus(v string) *DescribeBranchResponseBodyBranch {
	s.Status = &v
	return s
}

func (s *DescribeBranchResponseBodyBranch) SetTags(v []*DescribeBranchResponseBodyBranchTags) *DescribeBranchResponseBodyBranch {
	s.Tags = v
	return s
}

func (s *DescribeBranchResponseBodyBranch) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBranchResponseBodyBranchTags struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// dev
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBranchResponseBodyBranchTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeBranchResponseBodyBranchTags) GoString() string {
	return s.String()
}

func (s *DescribeBranchResponseBodyBranchTags) GetKey() *string {
	return s.Key
}

func (s *DescribeBranchResponseBodyBranchTags) GetValue() *string {
	return s.Value
}

func (s *DescribeBranchResponseBodyBranchTags) SetKey(v string) *DescribeBranchResponseBodyBranchTags {
	s.Key = &v
	return s
}

func (s *DescribeBranchResponseBodyBranchTags) SetValue(v string) *DescribeBranchResponseBodyBranchTags {
	s.Value = &v
	return s
}

func (s *DescribeBranchResponseBodyBranchTags) Validate() error {
	return dara.Validate(s)
}
