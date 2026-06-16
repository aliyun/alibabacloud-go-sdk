// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *TextModerationRequest
	GetService() *string
	SetServiceParameters(v string) *TextModerationRequest
	GetServiceParameters() *string
}

type TextModerationRequest struct {
	// The type of moderation service. Valid values:
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters for the moderation service. The value must be a JSON string.
	//
	// example:
	//
	// {"content":"The map is still black"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s TextModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s TextModerationRequest) GoString() string {
	return s.String()
}

func (s *TextModerationRequest) GetService() *string {
	return s.Service
}

func (s *TextModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *TextModerationRequest) SetService(v string) *TextModerationRequest {
	s.Service = &v
	return s
}

func (s *TextModerationRequest) SetServiceParameters(v string) *TextModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *TextModerationRequest) Validate() error {
	return dara.Validate(s)
}
