// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *CreateCredentialsResponseBody) *CreateCredentialsResponse
	GetBody() *CreateCredentialsResponseBody
}

type CreateCredentialsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialsResponse) GoString() string {
	return s.String()
}

func (s *CreateCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCredentialsResponse) GetBody() *CreateCredentialsResponseBody {
	return s.Body
}

func (s *CreateCredentialsResponse) SetHeaders(v map[string]*string) *CreateCredentialsResponse {
	s.Headers = v
	return s
}

func (s *CreateCredentialsResponse) SetStatusCode(v int32) *CreateCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCredentialsResponse) SetBody(v *CreateCredentialsResponseBody) *CreateCredentialsResponse {
	s.Body = v
	return s
}

func (s *CreateCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
