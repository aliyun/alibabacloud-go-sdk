// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *VoiceModerationCancelRequest
	GetService() *string
	SetServiceParameters(v string) *VoiceModerationCancelRequest
	GetServiceParameters() *string
}

type VoiceModerationCancelRequest struct {
	// The type of moderation service. Valid values include \\`nickname_detection\\` for user nicknames. Other values are to be determined.
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The ID of the task that you want to cancel.
	//
	// example:
	//
	// {
	//
	//         "taskId": "xxxxx-xxxx"
	//
	//     }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationCancelRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelRequest) GetService() *string {
	return s.Service
}

func (s *VoiceModerationCancelRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *VoiceModerationCancelRequest) SetService(v string) *VoiceModerationCancelRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationCancelRequest) SetServiceParameters(v string) *VoiceModerationCancelRequest {
	s.ServiceParameters = &v
	return s
}

func (s *VoiceModerationCancelRequest) Validate() error {
	return dara.Validate(s)
}
