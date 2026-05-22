// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPreloadExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteScheduledPreloadExecutionRequest
	GetId() *string
}

type DeleteScheduledPreloadExecutionRequest struct {
	// The ID of the prefetch plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// DeleteScheduledPreloadExecution
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScheduledPreloadExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadExecutionRequest) GetId() *string {
	return s.Id
}

func (s *DeleteScheduledPreloadExecutionRequest) SetId(v string) *DeleteScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

func (s *DeleteScheduledPreloadExecutionRequest) Validate() error {
	return dara.Validate(s)
}
