// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *QueryMaterialTaskDetailRequest
	GetTaskId() *string
}

type QueryMaterialTaskDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c3r127e325at9yd
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryMaterialTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskDetailRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryMaterialTaskDetailRequest) SetTaskId(v string) *QueryMaterialTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryMaterialTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
