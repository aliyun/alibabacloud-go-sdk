// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStatusCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizDate(v string) *GetInstanceStatusCountRequest
	GetBizDate() *string
	SetProjectEnv(v string) *GetInstanceStatusCountRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *GetInstanceStatusCountRequest
	GetProjectId() *int64
}

type GetInstanceStatusCountRequest struct {
	// The data timestamp of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-01-01
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The runtime environment. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetInstanceStatusCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusCountRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusCountRequest) GetBizDate() *string {
	return s.BizDate
}

func (s *GetInstanceStatusCountRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetInstanceStatusCountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceStatusCountRequest) SetBizDate(v string) *GetInstanceStatusCountRequest {
	s.BizDate = &v
	return s
}

func (s *GetInstanceStatusCountRequest) SetProjectEnv(v string) *GetInstanceStatusCountRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetInstanceStatusCountRequest) SetProjectId(v int64) *GetInstanceStatusCountRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceStatusCountRequest) Validate() error {
	return dara.Validate(s)
}
