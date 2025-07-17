// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientUniqueCode(v string) *DeleteWorkflowRequest
	GetClientUniqueCode() *string
	SetEnvType(v string) *DeleteWorkflowRequest
	GetEnvType() *string
	SetId(v int64) *DeleteWorkflowRequest
	GetId() *int64
}

type DeleteWorkflowRequest struct {
	// The unique code of the client. This parameter is used to create a workflow asynchronously and implement the idempotence of the workflow. If you do not specify this parameter when you create the workflow, the system automatically generates a unique code. The unique code is uniquely associated with the workflow ID. If you specify this parameter when you update or delete the workflow, the value of this parameter must be the unique code that is used to create the workflow.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowRequest) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *DeleteWorkflowRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *DeleteWorkflowRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteWorkflowRequest) SetClientUniqueCode(v string) *DeleteWorkflowRequest {
	s.ClientUniqueCode = &v
	return s
}

func (s *DeleteWorkflowRequest) SetEnvType(v string) *DeleteWorkflowRequest {
	s.EnvType = &v
	return s
}

func (s *DeleteWorkflowRequest) SetId(v int64) *DeleteWorkflowRequest {
	s.Id = &v
	return s
}

func (s *DeleteWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
