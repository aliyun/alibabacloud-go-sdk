// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveSpacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDriveSpacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDriveSpacesRequest
	GetNextToken() *string
	SetSpaceType(v string) *ListDriveSpacesRequest
	GetSpaceType() *string
	SetTenantContext(v *ListDriveSpacesRequestTenantContext) *ListDriveSpacesRequest
	GetTenantContext() *ListDriveSpacesRequestTenantContext
}

type ListDriveSpacesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// fekaf
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// org
	SpaceType     *string                              `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	TenantContext *ListDriveSpacesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ListDriveSpacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDriveSpacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDriveSpacesRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListDriveSpacesRequest) GetTenantContext() *ListDriveSpacesRequestTenantContext {
	return s.TenantContext
}

func (s *ListDriveSpacesRequest) SetMaxResults(v int32) *ListDriveSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDriveSpacesRequest) SetNextToken(v string) *ListDriveSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDriveSpacesRequest) SetSpaceType(v string) *ListDriveSpacesRequest {
	s.SpaceType = &v
	return s
}

func (s *ListDriveSpacesRequest) SetTenantContext(v *ListDriveSpacesRequestTenantContext) *ListDriveSpacesRequest {
	s.TenantContext = v
	return s
}

func (s *ListDriveSpacesRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDriveSpacesRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListDriveSpacesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListDriveSpacesRequestTenantContext) SetTenantId(v string) *ListDriveSpacesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListDriveSpacesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
