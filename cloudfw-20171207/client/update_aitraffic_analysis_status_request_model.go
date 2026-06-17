// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAITrafficAnalysisStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpdateAITrafficAnalysisStatusRequest
	GetStatus() *string
}

type UpdateAITrafficAnalysisStatusRequest struct {
	// The enabling status. Valid values:
	//
	// - **open**: Enables the feature.
	//
	// - **close**: Disables the feature.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAITrafficAnalysisStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAITrafficAnalysisStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAITrafficAnalysisStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAITrafficAnalysisStatusRequest) SetStatus(v string) *UpdateAITrafficAnalysisStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateAITrafficAnalysisStatusRequest) Validate() error {
	return dara.Validate(s)
}
