// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *VideoModerationResultRequest
	GetService() *string
	SetServiceParameters(v string) *VideoModerationResultRequest
	GetServiceParameters() *string
}

type VideoModerationResultRequest struct {
	// The service code for video moderation.
	//
	// example:
	//
	// videoDetection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameter set required by the moderation service. taskId specifies the task ID of the moderation task to query. Only one taskId can be specified per request.
	//
	// example:
	//
	// {"taskId":"au_f_8PoWiZKoLbczp5HRn69VdT-1y8@U5"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VideoModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationResultRequest) GoString() string {
	return s.String()
}

func (s *VideoModerationResultRequest) GetService() *string {
	return s.Service
}

func (s *VideoModerationResultRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *VideoModerationResultRequest) SetService(v string) *VideoModerationResultRequest {
	s.Service = &v
	return s
}

func (s *VideoModerationResultRequest) SetServiceParameters(v string) *VideoModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

func (s *VideoModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
