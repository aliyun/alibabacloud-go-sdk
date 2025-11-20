// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodesRequest
	GetNextToken() *string
	SetParentNodeId(v string) *ListNodesRequest
	GetParentNodeId() *string
	SetTenantContext(v *ListNodesRequestTenantContext) *ListNodesRequest
	GetTenantContext() *ListNodesRequestTenantContext
	SetWithPermissionRole(v bool) *ListNodesRequest
	GetWithPermissionRole() *bool
}

type ListNodesRequest struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// next_token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MNDoBb60VLBPraakI1Ywxyyn8lemrZQ3
	ParentNodeId  *string                        `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	TenantContext *ListNodesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// false
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodesRequest) GetParentNodeId() *string {
	return s.ParentNodeId
}

func (s *ListNodesRequest) GetTenantContext() *ListNodesRequestTenantContext {
	return s.TenantContext
}

func (s *ListNodesRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *ListNodesRequest) SetMaxResults(v int32) *ListNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodesRequest) SetNextToken(v string) *ListNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodesRequest) SetParentNodeId(v string) *ListNodesRequest {
	s.ParentNodeId = &v
	return s
}

func (s *ListNodesRequest) SetTenantContext(v *ListNodesRequestTenantContext) *ListNodesRequest {
	s.TenantContext = v
	return s
}

func (s *ListNodesRequest) SetWithPermissionRole(v bool) *ListNodesRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *ListNodesRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListNodesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListNodesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListNodesRequestTenantContext) SetTenantId(v string) *ListNodesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListNodesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
