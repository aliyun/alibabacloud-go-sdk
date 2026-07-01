// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchComputeClusterId(v string) *CreateAgenticDBBranchResponseBody
	GetBranchComputeClusterId() *string
	SetBranchId(v string) *CreateAgenticDBBranchResponseBody
	GetBranchId() *string
	SetBranchName(v string) *CreateAgenticDBBranchResponseBody
	GetBranchName() *string
	SetDBClusterId(v string) *CreateAgenticDBBranchResponseBody
	GetDBClusterId() *string
	SetParentBranchId(v string) *CreateAgenticDBBranchResponseBody
	GetParentBranchId() *string
	SetParentBranchName(v string) *CreateAgenticDBBranchResponseBody
	GetParentBranchName() *string
	SetProjectId(v string) *CreateAgenticDBBranchResponseBody
	GetProjectId() *string
	SetProjectName(v string) *CreateAgenticDBBranchResponseBody
	GetProjectName() *string
	SetRequestId(v string) *CreateAgenticDBBranchResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateAgenticDBBranchResponseBody
	GetStatus() *string
}

type CreateAgenticDBBranchResponseBody struct {
	// The ID of the compute cluster associated with the branch.
	//
	// example:
	//
	// pc-g0lsayq8c5qe
	BranchComputeClusterId *string `json:"BranchComputeClusterId,omitempty" xml:"BranchComputeClusterId,omitempty"`
	// The branch ID.
	//
	// example:
	//
	// br-7g8h9i0j1k2l
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The name of the branch.
	//
	// example:
	//
	// feature-analytics
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The AgenticDB cluster ID.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the parent branch.
	//
	// example:
	//
	// br-1a2b3c4d5e6f
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// The name of the parent branch.
	//
	// example:
	//
	// main
	ParentBranchName *string `json:"ParentBranchName,omitempty" xml:"ParentBranchName,omitempty"`
	// The ID of the project to which the branch belongs.
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the project to which the branch belongs.
	//
	// example:
	//
	// analytics-prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-1234567890AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the branch.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAgenticDBBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBBranchResponseBody) GetBranchComputeClusterId() *string {
	return s.BranchComputeClusterId
}

func (s *CreateAgenticDBBranchResponseBody) GetBranchId() *string {
	return s.BranchId
}

func (s *CreateAgenticDBBranchResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *CreateAgenticDBBranchResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgenticDBBranchResponseBody) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *CreateAgenticDBBranchResponseBody) GetParentBranchName() *string {
	return s.ParentBranchName
}

func (s *CreateAgenticDBBranchResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateAgenticDBBranchResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateAgenticDBBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgenticDBBranchResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateAgenticDBBranchResponseBody) SetBranchComputeClusterId(v string) *CreateAgenticDBBranchResponseBody {
	s.BranchComputeClusterId = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetBranchId(v string) *CreateAgenticDBBranchResponseBody {
	s.BranchId = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetBranchName(v string) *CreateAgenticDBBranchResponseBody {
	s.BranchName = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetDBClusterId(v string) *CreateAgenticDBBranchResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetParentBranchId(v string) *CreateAgenticDBBranchResponseBody {
	s.ParentBranchId = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetParentBranchName(v string) *CreateAgenticDBBranchResponseBody {
	s.ParentBranchName = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetProjectId(v string) *CreateAgenticDBBranchResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetProjectName(v string) *CreateAgenticDBBranchResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetRequestId(v string) *CreateAgenticDBBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) SetStatus(v string) *CreateAgenticDBBranchResponseBody {
	s.Status = &v
	return s
}

func (s *CreateAgenticDBBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
