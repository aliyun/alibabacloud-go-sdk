// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetWorkflowInstanceRequest
	GetId() *int64
}

type GetWorkflowInstanceRequest struct {
	// The ID of the workflow instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetWorkflowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceRequest) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowInstanceRequest) SetId(v int64) *GetWorkflowInstanceRequest {
	s.Id = &v
	return s
}

func (s *GetWorkflowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
