// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *DeleteMaterialTaskRequest
	GetTaskIds() []*string
}

type DeleteMaterialTaskRequest struct {
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DeleteMaterialTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterialTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DeleteMaterialTaskRequest) SetTaskIds(v []*string) *DeleteMaterialTaskRequest {
	s.TaskIds = v
	return s
}

func (s *DeleteMaterialTaskRequest) Validate() error {
	return dara.Validate(s)
}
