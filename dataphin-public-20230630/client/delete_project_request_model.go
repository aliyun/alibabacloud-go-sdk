// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteProjectRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteProjectRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteProjectRequest
	GetOpUserId() *string
}

type DeleteProjectRequest struct {
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteProjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteProjectRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteProjectRequest) SetId(v int64) *DeleteProjectRequest {
	s.Id = &v
	return s
}

func (s *DeleteProjectRequest) SetOpTenantId(v int64) *DeleteProjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteProjectRequest) SetOpUserId(v string) *DeleteProjectRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteProjectRequest) Validate() error {
	return dara.Validate(s)
}
