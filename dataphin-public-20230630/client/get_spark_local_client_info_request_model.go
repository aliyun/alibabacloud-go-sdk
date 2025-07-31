// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkLocalClientInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvEnum(v string) *GetSparkLocalClientInfoRequest
	GetEnvEnum() *string
	SetOpTenantId(v int64) *GetSparkLocalClientInfoRequest
	GetOpTenantId() *int64
	SetProjectId(v string) *GetSparkLocalClientInfoRequest
	GetProjectId() *string
}

type GetSparkLocalClientInfoRequest struct {
	// This parameter is required.
	EnvEnum *string `json:"EnvEnum,omitempty" xml:"EnvEnum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetSparkLocalClientInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkLocalClientInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSparkLocalClientInfoRequest) GetEnvEnum() *string {
	return s.EnvEnum
}

func (s *GetSparkLocalClientInfoRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSparkLocalClientInfoRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSparkLocalClientInfoRequest) SetEnvEnum(v string) *GetSparkLocalClientInfoRequest {
	s.EnvEnum = &v
	return s
}

func (s *GetSparkLocalClientInfoRequest) SetOpTenantId(v int64) *GetSparkLocalClientInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSparkLocalClientInfoRequest) SetProjectId(v string) *GetSparkLocalClientInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSparkLocalClientInfoRequest) Validate() error {
	return dara.Validate(s)
}
