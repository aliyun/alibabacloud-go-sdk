// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAbacAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAbacAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAbacAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *CreateAbacAuthorizationResponseBody) *CreateAbacAuthorizationResponse
	GetBody() *CreateAbacAuthorizationResponseBody
}

type CreateAbacAuthorizationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAbacAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAbacAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAbacAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *CreateAbacAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAbacAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAbacAuthorizationResponse) GetBody() *CreateAbacAuthorizationResponseBody {
	return s.Body
}

func (s *CreateAbacAuthorizationResponse) SetHeaders(v map[string]*string) *CreateAbacAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *CreateAbacAuthorizationResponse) SetStatusCode(v int32) *CreateAbacAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAbacAuthorizationResponse) SetBody(v *CreateAbacAuthorizationResponseBody) *CreateAbacAuthorizationResponse {
	s.Body = v
	return s
}

func (s *CreateAbacAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
