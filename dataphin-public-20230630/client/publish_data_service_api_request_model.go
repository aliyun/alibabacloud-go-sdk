// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *PublishDataServiceApiRequest
	GetApiId() *int64
	SetOpTenantId(v int64) *PublishDataServiceApiRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *PublishDataServiceApiRequest
	GetProjectId() *int32
	SetVersionId(v string) *PublishDataServiceApiRequest
	GetVersionId() *string
}

type PublishDataServiceApiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
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
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s PublishDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *PublishDataServiceApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *PublishDataServiceApiRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PublishDataServiceApiRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *PublishDataServiceApiRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PublishDataServiceApiRequest) SetApiId(v int64) *PublishDataServiceApiRequest {
	s.ApiId = &v
	return s
}

func (s *PublishDataServiceApiRequest) SetOpTenantId(v int64) *PublishDataServiceApiRequest {
	s.OpTenantId = &v
	return s
}

func (s *PublishDataServiceApiRequest) SetProjectId(v int32) *PublishDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *PublishDataServiceApiRequest) SetVersionId(v string) *PublishDataServiceApiRequest {
	s.VersionId = &v
	return s
}

func (s *PublishDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}
