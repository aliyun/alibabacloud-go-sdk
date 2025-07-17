// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *VideoModerationCancelRequest
	GetService() *string
	SetServiceParameters(v string) *VideoModerationCancelRequest
	GetServiceParameters() *string
}

type VideoModerationCancelRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// videoDetection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {\\"taskId\\":\\"vi_s_4O9gp7GfNQdx9GOqdekFmk-1z2RJT\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VideoModerationCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationCancelRequest) GoString() string {
	return s.String()
}

func (s *VideoModerationCancelRequest) GetService() *string {
	return s.Service
}

func (s *VideoModerationCancelRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *VideoModerationCancelRequest) SetService(v string) *VideoModerationCancelRequest {
	s.Service = &v
	return s
}

func (s *VideoModerationCancelRequest) SetServiceParameters(v string) *VideoModerationCancelRequest {
	s.ServiceParameters = &v
	return s
}

func (s *VideoModerationCancelRequest) Validate() error {
	return dara.Validate(s)
}
