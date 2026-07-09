// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminatePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *TerminatePipelineRequest
	GetReason() *string
}

type TerminatePipelineRequest struct {
	// example:
	//
	// project deprecated
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s TerminatePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminatePipelineRequest) GoString() string {
	return s.String()
}

func (s *TerminatePipelineRequest) GetReason() *string {
	return s.Reason
}

func (s *TerminatePipelineRequest) SetReason(v string) *TerminatePipelineRequest {
	s.Reason = &v
	return s
}

func (s *TerminatePipelineRequest) Validate() error {
	return dara.Validate(s)
}
