// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DeleteResourceRequest
	GetComment() *string
	SetId(v int64) *DeleteResourceRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteResourceRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *DeleteResourceRequest
	GetProjectId() *int64
}

type DeleteResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10300010201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// 1030111021
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) GetComment() *string {
	return s.Comment
}

func (s *DeleteResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteResourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteResourceRequest) SetComment(v string) *DeleteResourceRequest {
	s.Comment = &v
	return s
}

func (s *DeleteResourceRequest) SetId(v int64) *DeleteResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteResourceRequest) SetOpTenantId(v int64) *DeleteResourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteResourceRequest) SetProjectId(v int64) *DeleteResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteResourceRequest) Validate() error {
	return dara.Validate(s)
}
