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
	SetOpUserId(v string) *GetDevObjectDependencyRequest
	GetOpUserId() *string
	SetProjectId(v int64) *GetDevObjectDependencyRequest
	GetProjectId() *int64
}

type GetDevObjectDependencyRequest struct {
	// The object source. Valid values:
	//
	// - BLACK_BOX: logical table
	//
	// - DATA_PROCESS: computing node
	//
	// - PIPELINE: integration node
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_PROCESS
	ObjectFrom *string `json:"ObjectFrom,omitempty" xml:"ObjectFrom,omitempty"`
	// The object ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7026498387616064
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The object type. Valid values:
	//
	// - FILE_ID: file ID of a node
	//
	// - MODEL_ID: logical table
	//
	// This parameter is required.
	//
	// example:
	//
	// 7026498387616064
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The project ID.
	//
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

func (s *GetDevObjectDependencyRequest) GetOpUserId() *string {
	return s.OpUserId
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

func (s *GetDevObjectDependencyRequest) SetOpUserId(v string) *GetDevObjectDependencyRequest {
	s.OpUserId = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetProjectId(v int64) *GetDevObjectDependencyRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDevObjectDependencyRequest) Validate() error {
	return dara.Validate(s)
}
