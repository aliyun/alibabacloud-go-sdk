// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextModerationPlusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *TextModerationPlusRequest
	GetService() *string
	SetServiceParameters(v string) *TextModerationPlusRequest
	GetServiceParameters() *string
}

type TextModerationPlusRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// ugc_moderation_byllm
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The set of parameters required for the moderation service. The value must be a JSON string.
	//
	// example:
	//
	// {"content":"Test content"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s TextModerationPlusRequest) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusRequest) GoString() string {
	return s.String()
}

func (s *TextModerationPlusRequest) GetService() *string {
	return s.Service
}

func (s *TextModerationPlusRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *TextModerationPlusRequest) SetService(v string) *TextModerationPlusRequest {
	s.Service = &v
	return s
}

func (s *TextModerationPlusRequest) SetServiceParameters(v string) *TextModerationPlusRequest {
	s.ServiceParameters = &v
	return s
}

func (s *TextModerationPlusRequest) Validate() error {
	return dara.Validate(s)
}
