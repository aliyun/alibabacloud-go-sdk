// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachProjectsFromProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachProjectsFromProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachProjectsFromProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *DetachProjectsFromProjectAuthorizationResponseBody) *DetachProjectsFromProjectAuthorizationResponse
	GetBody() *DetachProjectsFromProjectAuthorizationResponseBody
}

type DetachProjectsFromProjectAuthorizationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachProjectsFromProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachProjectsFromProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachProjectsFromProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DetachProjectsFromProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachProjectsFromProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachProjectsFromProjectAuthorizationResponse) GetBody() *DetachProjectsFromProjectAuthorizationResponseBody {
	return s.Body
}

func (s *DetachProjectsFromProjectAuthorizationResponse) SetHeaders(v map[string]*string) *DetachProjectsFromProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DetachProjectsFromProjectAuthorizationResponse) SetStatusCode(v int32) *DetachProjectsFromProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachProjectsFromProjectAuthorizationResponse) SetBody(v *DetachProjectsFromProjectAuthorizationResponseBody) *DetachProjectsFromProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *DetachProjectsFromProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
