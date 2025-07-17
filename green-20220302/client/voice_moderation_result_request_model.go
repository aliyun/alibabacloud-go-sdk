// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *VoiceModerationResultRequest
	GetService() *string
	SetServiceParameters(v string) *VoiceModerationResultRequest
	GetServiceParameters() *string
}

type VoiceModerationResultRequest struct {
	// The type of the moderation service. Valid values: nickname_detection: user nickname
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters of API requests that are sent from API Gateway to the backend service.
	//
	// For more information, see [ServiceParameter](https://help.aliyun.com/document_detail/43988.html).
	//
	// example:
	//
	// {"taskId":"xxxxx-xxxx"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResultRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultRequest) GetService() *string {
	return s.Service
}

func (s *VoiceModerationResultRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *VoiceModerationResultRequest) SetService(v string) *VoiceModerationResultRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationResultRequest) SetServiceParameters(v string) *VoiceModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

func (s *VoiceModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
