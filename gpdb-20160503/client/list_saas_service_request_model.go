// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSaasServiceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSaasServiceRequest
	GetNextToken() *string
	SetRegionId(v string) *ListSaasServiceRequest
	GetRegionId() *string
	SetServiceType(v string) *ListSaasServiceRequest
	GetServiceType() *string
	SetWorkspaceId(v string) *ListSaasServiceRequest
	GetWorkspaceId() *string
}

type ListSaasServiceRequest struct {
	// The maximum number of entries to return in this request. Default value: 10.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query to begin with during a paginated query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service type, with the following value:
	//
	// - **memroy**
	//
	// - **drama**
	//
	// This parameter is required.
	//
	// example:
	//
	// drama
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListSaasServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSaasServiceRequest) GoString() string {
	return s.String()
}

func (s *ListSaasServiceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSaasServiceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSaasServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSaasServiceRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListSaasServiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSaasServiceRequest) SetMaxResults(v int32) *ListSaasServiceRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSaasServiceRequest) SetNextToken(v string) *ListSaasServiceRequest {
	s.NextToken = &v
	return s
}

func (s *ListSaasServiceRequest) SetRegionId(v string) *ListSaasServiceRequest {
	s.RegionId = &v
	return s
}

func (s *ListSaasServiceRequest) SetServiceType(v string) *ListSaasServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *ListSaasServiceRequest) SetWorkspaceId(v string) *ListSaasServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSaasServiceRequest) Validate() error {
	return dara.Validate(s)
}
