// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggTaskGroupStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpdateAggTaskGroupStatusRequest
	GetStatus() *string
}

type UpdateAggTaskGroupStatusRequest struct {
	// Status of the aggregation task group, either “Running” or “Stopped”. Default is Running.
	//
	// This parameter is required.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateAggTaskGroupStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggTaskGroupStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggTaskGroupStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAggTaskGroupStatusRequest) SetStatus(v string) *UpdateAggTaskGroupStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateAggTaskGroupStatusRequest) Validate() error {
	return dara.Validate(s)
}
