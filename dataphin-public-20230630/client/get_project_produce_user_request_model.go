// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectProduceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetProjectProduceUserRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetProjectProduceUserRequest
	GetProjectId() *int64
}

type GetProjectProduceUserRequest struct {
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
	// 131311111321
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectProduceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectProduceUserRequest) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetProjectProduceUserRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectProduceUserRequest) SetOpTenantId(v int64) *GetProjectProduceUserRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetProjectProduceUserRequest) SetProjectId(v int64) *GetProjectProduceUserRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectProduceUserRequest) Validate() error {
	return dara.Validate(s)
}
