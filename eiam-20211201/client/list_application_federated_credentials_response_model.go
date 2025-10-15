// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationFederatedCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationFederatedCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationFederatedCredentialsResponseBody) *ListApplicationFederatedCredentialsResponse
	GetBody() *ListApplicationFederatedCredentialsResponseBody
}

type ListApplicationFederatedCredentialsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationFederatedCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationFederatedCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationFederatedCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationFederatedCredentialsResponse) GetBody() *ListApplicationFederatedCredentialsResponseBody {
	return s.Body
}

func (s *ListApplicationFederatedCredentialsResponse) SetHeaders(v map[string]*string) *ListApplicationFederatedCredentialsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponse) SetStatusCode(v int32) *ListApplicationFederatedCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationFederatedCredentialsResponse) SetBody(v *ListApplicationFederatedCredentialsResponseBody) *ListApplicationFederatedCredentialsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationFederatedCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
