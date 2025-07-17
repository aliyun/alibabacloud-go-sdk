// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceParameters(v string) *ManualModerationResultRequest
	GetServiceParameters() *string
}

type ManualModerationResultRequest struct {
	// Set of parameters required by the service, in JSON string format.
	//
	// - TaskId: The task ID returned when the task was submitted.
	//
	// example:
	//
	// {\\"TaskId\\":\\"e5f2d886-4c23-440d-999c-bd98acde11b6\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ManualModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResultRequest) GoString() string {
	return s.String()
}

func (s *ManualModerationResultRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ManualModerationResultRequest) SetServiceParameters(v string) *ManualModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ManualModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
