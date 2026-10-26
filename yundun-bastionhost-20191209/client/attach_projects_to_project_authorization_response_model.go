// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachProjectsToProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachProjectsToProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachProjectsToProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *AttachProjectsToProjectAuthorizationResponseBody) *AttachProjectsToProjectAuthorizationResponse
	GetBody() *AttachProjectsToProjectAuthorizationResponseBody
}

type AttachProjectsToProjectAuthorizationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachProjectsToProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachProjectsToProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachProjectsToProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *AttachProjectsToProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachProjectsToProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachProjectsToProjectAuthorizationResponse) GetBody() *AttachProjectsToProjectAuthorizationResponseBody {
	return s.Body
}

func (s *AttachProjectsToProjectAuthorizationResponse) SetHeaders(v map[string]*string) *AttachProjectsToProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *AttachProjectsToProjectAuthorizationResponse) SetStatusCode(v int32) *AttachProjectsToProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachProjectsToProjectAuthorizationResponse) SetBody(v *AttachProjectsToProjectAuthorizationResponseBody) *AttachProjectsToProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *AttachProjectsToProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
