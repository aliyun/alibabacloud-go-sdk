// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *PausePipelineRequest
	GetReason() *string
}

type PausePipelineRequest struct {
	// The reason for pausing the pipeline.
	//
	// example:
	//
	// manual maintenance
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s PausePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s PausePipelineRequest) GoString() string {
	return s.String()
}

func (s *PausePipelineRequest) GetReason() *string {
	return s.Reason
}

func (s *PausePipelineRequest) SetReason(v string) *PausePipelineRequest {
	s.Reason = &v
	return s
}

func (s *PausePipelineRequest) Validate() error {
	return dara.Validate(s)
}
