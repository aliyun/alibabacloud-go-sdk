// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartScheduledPreloadExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StartScheduledPreloadExecutionRequest
	GetId() *string
}

type StartScheduledPreloadExecutionRequest struct {
	// The ID of the prefetch plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// StartScheduledPreloadExecution
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StartScheduledPreloadExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartScheduledPreloadExecutionRequest) GetId() *string {
	return s.Id
}

func (s *StartScheduledPreloadExecutionRequest) SetId(v string) *StartScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

func (s *StartScheduledPreloadExecutionRequest) Validate() error {
	return dara.Validate(s)
}
