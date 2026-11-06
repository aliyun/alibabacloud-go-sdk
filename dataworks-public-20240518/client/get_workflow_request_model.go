// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *GetWorkflowRequest
	GetEnvType() *string
	SetId(v int64) *GetWorkflowRequest
	GetId() *int64
}

type GetWorkflowRequest struct {
	// The project environment. Valid values:
	//
	// - Prod: Production environment.
	//
	// - Dev: Development environment.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The unique identifier of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *GetWorkflowRequest) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowRequest) SetEnvType(v string) *GetWorkflowRequest {
	s.EnvType = &v
	return s
}

func (s *GetWorkflowRequest) SetId(v int64) *GetWorkflowRequest {
	s.Id = &v
	return s
}

func (s *GetWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
