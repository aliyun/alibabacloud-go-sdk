// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *AbolishDataServiceApiRequest
	GetApiId() *int64
	SetProjectId(v int64) *AbolishDataServiceApiRequest
	GetProjectId() *int64
	SetTenantId(v int64) *AbolishDataServiceApiRequest
	GetTenantId() *int64
}

type AbolishDataServiceApiRequest struct {
	// The ID of the DataService Studio API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the workspace.
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

func (s AbolishDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s AbolishDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *AbolishDataServiceApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *AbolishDataServiceApiRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AbolishDataServiceApiRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *AbolishDataServiceApiRequest) SetApiId(v int64) *AbolishDataServiceApiRequest {
	s.ApiId = &v
	return s
}

func (s *AbolishDataServiceApiRequest) SetProjectId(v int64) *AbolishDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *AbolishDataServiceApiRequest) SetTenantId(v int64) *AbolishDataServiceApiRequest {
	s.TenantId = &v
	return s
}

func (s *AbolishDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}
