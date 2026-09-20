// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *DeleteDataServiceApiRequest
	GetApiId() *int64
	SetProjectId(v int64) *DeleteDataServiceApiRequest
	GetProjectId() *int64
	SetTenantId(v int64) *DeleteDataServiceApiRequest
	GetTenantId() *int64
}

type DeleteDataServiceApiRequest struct {
	// The ID of the DataService API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The workspace ID.
	//
	// You can obtain the workspace ID from PageResult.ProjectList[].ProjectId in the response of ListProjects.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console), go to the DataStudio page, click the username in the upper-right corner, and choose Menu > User Info to obtain the tenant ID.
	//
	// You can also obtain the tenant ID from Data.Apis[].TenantId in the response of ListDataServiceApis.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *DeleteDataServiceApiRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDataServiceApiRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DeleteDataServiceApiRequest) SetApiId(v int64) *DeleteDataServiceApiRequest {
	s.ApiId = &v
	return s
}

func (s *DeleteDataServiceApiRequest) SetProjectId(v int64) *DeleteDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDataServiceApiRequest) SetTenantId(v int64) *DeleteDataServiceApiRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}
