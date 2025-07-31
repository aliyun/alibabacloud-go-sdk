// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDevObjectDependencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectFrom(v string) *GetDevObjectDependencyRequest
	GetObjectFrom() *string
	SetObjectId(v string) *GetDevObjectDependencyRequest
	GetObjectId() *string
	SetObjectType(v string) *GetDevObjectDependencyRequest
	GetObjectType() *string
	SetOpTenantId(v int64) *GetDevObjectDependencyRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetDevObjectDependencyRequest
	GetProjectId() *int64
}

type GetDevObjectDependencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DATA_PROCESS
	ObjectFrom *string `json:"ObjectFrom,omitempty" xml:"ObjectFrom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7026498387616064
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7026498387616064
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7021037162911616L
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDevObjectDependencyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDevObjectDependencyRequest) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyRequest) GetObjectFrom() *string {
	return s.ObjectFrom
}

func (s *GetDevObjectDependencyRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetDevObjectDependencyRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetDevObjectDependencyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDevObjectDependencyRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDevObjectDependencyRequest) SetObjectFrom(v string) *GetDevObjectDependencyRequest {
	s.ObjectFrom = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetObjectId(v string) *GetDevObjectDependencyRequest {
	s.ObjectId = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetObjectType(v string) *GetDevObjectDependencyRequest {
	s.ObjectType = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetOpTenantId(v int64) *GetDevObjectDependencyRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetProjectId(v int64) *GetDevObjectDependencyRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDevObjectDependencyRequest) Validate() error {
	return dara.Validate(s)
}
