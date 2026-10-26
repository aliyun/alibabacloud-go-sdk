// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachProjectManagersToProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachProjectManagersToProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachProjectManagersToProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *AttachProjectManagersToProjectAuthorizationResponseBody) *AttachProjectManagersToProjectAuthorizationResponse
	GetBody() *AttachProjectManagersToProjectAuthorizationResponseBody
}

type AttachProjectManagersToProjectAuthorizationResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachProjectManagersToProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachProjectManagersToProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachProjectManagersToProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *AttachProjectManagersToProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachProjectManagersToProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachProjectManagersToProjectAuthorizationResponse) GetBody() *AttachProjectManagersToProjectAuthorizationResponseBody {
	return s.Body
}

func (s *AttachProjectManagersToProjectAuthorizationResponse) SetHeaders(v map[string]*string) *AttachProjectManagersToProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationResponse) SetStatusCode(v int32) *AttachProjectManagersToProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationResponse) SetBody(v *AttachProjectManagersToProjectAuthorizationResponseBody) *AttachProjectManagersToProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
