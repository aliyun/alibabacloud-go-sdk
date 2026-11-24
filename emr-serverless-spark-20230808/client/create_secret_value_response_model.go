// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecretValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecretValueResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecretValueResponseBody) *CreateSecretValueResponse
	GetBody() *CreateSecretValueResponseBody
}

type CreateSecretValueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecretValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecretValueResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretValueResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecretValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecretValueResponse) GetBody() *CreateSecretValueResponseBody {
	return s.Body
}

func (s *CreateSecretValueResponse) SetHeaders(v map[string]*string) *CreateSecretValueResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretValueResponse) SetStatusCode(v int32) *CreateSecretValueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecretValueResponse) SetBody(v *CreateSecretValueResponseBody) *CreateSecretValueResponse {
	s.Body = v
	return s
}

func (s *CreateSecretValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
