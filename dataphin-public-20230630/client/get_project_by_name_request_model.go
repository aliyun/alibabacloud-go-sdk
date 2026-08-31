// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetProjectByNameRequest
	GetName() *string
	SetOpTenantId(v int64) *GetProjectByNameRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetProjectByNameRequest
	GetOpUserId() *string
}

type GetProjectByNameRequest struct {
	// The project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetProjectByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectByNameRequest) GoString() string {
	return s.String()
}

func (s *GetProjectByNameRequest) GetName() *string {
	return s.Name
}

func (s *GetProjectByNameRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetProjectByNameRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetProjectByNameRequest) SetName(v string) *GetProjectByNameRequest {
	s.Name = &v
	return s
}

func (s *GetProjectByNameRequest) SetOpTenantId(v int64) *GetProjectByNameRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetProjectByNameRequest) SetOpUserId(v string) *GetProjectByNameRequest {
	s.OpUserId = &v
	return s
}

func (s *GetProjectByNameRequest) Validate() error {
	return dara.Validate(s)
}
