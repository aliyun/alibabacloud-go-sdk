// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *GetScheduledTaskExecutionDetailRequest
	GetExecutionId() *string
	SetTenantId(v string) *GetScheduledTaskExecutionDetailRequest
	GetTenantId() *string
}

type GetScheduledTaskExecutionDetailRequest struct {
	// The execution record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	// The tenant ID that takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScheduledTaskExecutionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionDetailRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *GetScheduledTaskExecutionDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskExecutionDetailRequest) SetExecutionId(v string) *GetScheduledTaskExecutionDetailRequest {
	s.ExecutionId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailRequest) SetTenantId(v string) *GetScheduledTaskExecutionDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailRequest) Validate() error {
	return dara.Validate(s)
}
