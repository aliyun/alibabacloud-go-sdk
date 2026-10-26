// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectAuthorizationResponseBody) *GetProjectAuthorizationResponse
	GetBody() *GetProjectAuthorizationResponseBody
}

type GetProjectAuthorizationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GetProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectAuthorizationResponse) GetBody() *GetProjectAuthorizationResponseBody {
	return s.Body
}

func (s *GetProjectAuthorizationResponse) SetHeaders(v map[string]*string) *GetProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GetProjectAuthorizationResponse) SetStatusCode(v int32) *GetProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectAuthorizationResponse) SetBody(v *GetProjectAuthorizationResponseBody) *GetProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *GetProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
