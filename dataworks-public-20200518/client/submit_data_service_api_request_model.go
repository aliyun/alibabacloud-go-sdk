// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *SubmitDataServiceApiRequest
	GetApiId() *int64
	SetProjectId(v int64) *SubmitDataServiceApiRequest
	GetProjectId() *int64
	SetTenantId(v int64) *SubmitDataServiceApiRequest
	GetTenantId() *int64
}

type SubmitDataServiceApiRequest struct {
	// The API ID. You can call the [ListDataServiceApis](~~ListDataServiceApis~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the DataWorks console. Find your workspace and go to the [DataService Studio](https://ds-cn-shanghai.data.aliyun.com/) page. On the DataService Studio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s SubmitDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *SubmitDataServiceApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *SubmitDataServiceApiRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SubmitDataServiceApiRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *SubmitDataServiceApiRequest) SetApiId(v int64) *SubmitDataServiceApiRequest {
	s.ApiId = &v
	return s
}

func (s *SubmitDataServiceApiRequest) SetProjectId(v int64) *SubmitDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitDataServiceApiRequest) SetTenantId(v int64) *SubmitDataServiceApiRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}
