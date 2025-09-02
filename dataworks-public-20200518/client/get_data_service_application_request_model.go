// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *GetDataServiceApplicationRequest
	GetApplicationId() *int64
	SetProjectId(v int64) *GetDataServiceApplicationRequest
	GetProjectId() *int64
	SetTenantId(v int64) *GetDataServiceApplicationRequest
	GetTenantId() *int64
}

type GetDataServiceApplicationRequest struct {
	// The ID of the application. You can view the information about the application in the API Gateway console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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

func (s GetDataServiceApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApplicationRequest) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *GetDataServiceApplicationRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServiceApplicationRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataServiceApplicationRequest) SetApplicationId(v int64) *GetDataServiceApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetDataServiceApplicationRequest) SetProjectId(v int64) *GetDataServiceApplicationRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceApplicationRequest) SetTenantId(v int64) *GetDataServiceApplicationRequest {
	s.TenantId = &v
	return s
}

func (s *GetDataServiceApplicationRequest) Validate() error {
	return dara.Validate(s)
}
