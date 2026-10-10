// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetGroupSourceRequest
	GetGroupId() *string
	SetSourceId(v string) *GetGroupSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *GetGroupSourceRequest
	GetTenantId() *string
}

type GetGroupSourceRequest struct {
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The ID of the personal FILE data source to be replaced. The ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGroupSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupSourceRequest) GoString() string {
	return s.String()
}

func (s *GetGroupSourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetGroupSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGroupSourceRequest) SetGroupId(v string) *GetGroupSourceRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupSourceRequest) SetSourceId(v string) *GetGroupSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetGroupSourceRequest) SetTenantId(v string) *GetGroupSourceRequest {
	s.TenantId = &v
	return s
}

func (s *GetGroupSourceRequest) Validate() error {
	return dara.Validate(s)
}
