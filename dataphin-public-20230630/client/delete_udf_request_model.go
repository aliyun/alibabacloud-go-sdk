// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DeleteUdfRequest
	GetComment() *string
	SetId(v int64) *DeleteUdfRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteUdfRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *DeleteUdfRequest
	GetProjectId() *int64
}

type DeleteUdfRequest struct {
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

func (s DeleteUdfRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfRequest) GoString() string {
	return s.String()
}

func (s *DeleteUdfRequest) GetComment() *string {
	return s.Comment
}

func (s *DeleteUdfRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteUdfRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteUdfRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteUdfRequest) SetComment(v string) *DeleteUdfRequest {
	s.Comment = &v
	return s
}

func (s *DeleteUdfRequest) SetId(v int64) *DeleteUdfRequest {
	s.Id = &v
	return s
}

func (s *DeleteUdfRequest) SetOpTenantId(v int64) *DeleteUdfRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteUdfRequest) SetProjectId(v int64) *DeleteUdfRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteUdfRequest) Validate() error {
	return dara.Validate(s)
}
