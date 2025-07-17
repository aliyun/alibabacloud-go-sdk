// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetTaskInstanceRequest
	GetId() *int64
}

type GetTaskInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceRequest) GetId() *int64 {
	return s.Id
}

func (s *GetTaskInstanceRequest) SetId(v int64) *GetTaskInstanceRequest {
	s.Id = &v
	return s
}

func (s *GetTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
