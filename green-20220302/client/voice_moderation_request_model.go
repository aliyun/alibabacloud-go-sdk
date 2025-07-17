// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *VoiceModerationRequest
	GetService() *string
	SetServiceParameters(v string) *VoiceModerationRequest
	GetServiceParameters() *string
}

type VoiceModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"url": "http://aliyundoc.com/test.flv", "dataId": "data1234"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationRequest) GetService() *string {
	return s.Service
}

func (s *VoiceModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *VoiceModerationRequest) SetService(v string) *VoiceModerationRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationRequest) SetServiceParameters(v string) *VoiceModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *VoiceModerationRequest) Validate() error {
	return dara.Validate(s)
}
