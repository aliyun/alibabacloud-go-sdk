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
	// The client unique code of the workflow, which is used to implement asynchronous operations and idempotence. If you do not specify this parameter during creation, the system automatically generates one. This code is uniquely bound to the resource ID. If you specify this parameter during update or deletion, it must be the same as the client unique code specified during creation.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production
	//
	// - Dev: development
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
