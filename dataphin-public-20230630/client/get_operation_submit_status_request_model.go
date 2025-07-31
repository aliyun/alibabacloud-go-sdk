// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationSubmitStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetOperationSubmitStatusRequest
	GetEnv() *string
	SetJobId(v string) *GetOperationSubmitStatusRequest
	GetJobId() *string
	SetOpTenantId(v int64) *GetOperationSubmitStatusRequest
	GetOpTenantId() *int64
}

type GetOperationSubmitStatusRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1324444131
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetOperationSubmitStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationSubmitStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusRequest) GetEnv() *string {
	return s.Env
}

func (s *GetOperationSubmitStatusRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetOperationSubmitStatusRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetOperationSubmitStatusRequest) SetEnv(v string) *GetOperationSubmitStatusRequest {
	s.Env = &v
	return s
}

func (s *GetOperationSubmitStatusRequest) SetJobId(v string) *GetOperationSubmitStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetOperationSubmitStatusRequest) SetOpTenantId(v int64) *GetOperationSubmitStatusRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetOperationSubmitStatusRequest) Validate() error {
	return dara.Validate(s)
}
