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
	SetProjectId(v int64) *PublishDataServiceApiRequest
	GetProjectId() *int64
	SetTenantId(v int64) *PublishDataServiceApiRequest
	GetTenantId() *int64
}

type PublishDataServiceApiRequest struct {
	// The API ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
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

func (s *PublishDataServiceApiRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *PublishDataServiceApiRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *PublishDataServiceApiRequest) SetApiId(v int64) *PublishDataServiceApiRequest {
	s.ApiId = &v
	return s
}

func (s *PublishDataServiceApiRequest) SetProjectId(v int64) *PublishDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *PublishDataServiceApiRequest) SetTenantId(v int64) *PublishDataServiceApiRequest {
	s.TenantId = &v
	return s
}

func (s *PublishDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}
