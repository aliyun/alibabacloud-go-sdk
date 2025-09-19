// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCycleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DeleteCycleTaskRequest
	GetConfigId() *string
}

type DeleteCycleTaskRequest struct {
	// The ID of the task configuration.
	//
	// >  You can call the [DescribeCycleTaskList](~~DescribeCycleTaskList~~) operation to query the IDs of task configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 435f626256ebf564cf5ba966a539****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s DeleteCycleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCycleTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteCycleTaskRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteCycleTaskRequest) SetConfigId(v string) *DeleteCycleTaskRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteCycleTaskRequest) Validate() error {
	return dara.Validate(s)
}
