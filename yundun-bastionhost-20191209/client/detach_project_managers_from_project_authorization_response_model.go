// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachProjectManagersFromProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachProjectManagersFromProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachProjectManagersFromProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *DetachProjectManagersFromProjectAuthorizationResponseBody) *DetachProjectManagersFromProjectAuthorizationResponse
	GetBody() *DetachProjectManagersFromProjectAuthorizationResponseBody
}

type DetachProjectManagersFromProjectAuthorizationResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachProjectManagersFromProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachProjectManagersFromProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachProjectManagersFromProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DetachProjectManagersFromProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachProjectManagersFromProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachProjectManagersFromProjectAuthorizationResponse) GetBody() *DetachProjectManagersFromProjectAuthorizationResponseBody {
	return s.Body
}

func (s *DetachProjectManagersFromProjectAuthorizationResponse) SetHeaders(v map[string]*string) *DetachProjectManagersFromProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DetachProjectManagersFromProjectAuthorizationResponse) SetStatusCode(v int32) *DetachProjectManagersFromProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachProjectManagersFromProjectAuthorizationResponse) SetBody(v *DetachProjectManagersFromProjectAuthorizationResponseBody) *DetachProjectManagersFromProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *DetachProjectManagersFromProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
