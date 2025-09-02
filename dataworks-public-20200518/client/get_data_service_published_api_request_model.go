// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServicePublishedApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *GetDataServicePublishedApiRequest
	GetApiId() *int64
	SetProjectId(v int64) *GetDataServicePublishedApiRequest
	GetProjectId() *int64
	SetTenantId(v int64) *GetDataServicePublishedApiRequest
	GetTenantId() *int64
}

type GetDataServicePublishedApiRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// example:
	//
	// 10003
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDataServicePublishedApiRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServicePublishedApiRequest) GoString() string {
	return s.String()
}

func (s *GetDataServicePublishedApiRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *GetDataServicePublishedApiRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServicePublishedApiRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataServicePublishedApiRequest) SetApiId(v int64) *GetDataServicePublishedApiRequest {
	s.ApiId = &v
	return s
}

func (s *GetDataServicePublishedApiRequest) SetProjectId(v int64) *GetDataServicePublishedApiRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServicePublishedApiRequest) SetTenantId(v int64) *GetDataServicePublishedApiRequest {
	s.TenantId = &v
	return s
}

func (s *GetDataServicePublishedApiRequest) Validate() error {
	return dara.Validate(s)
}
