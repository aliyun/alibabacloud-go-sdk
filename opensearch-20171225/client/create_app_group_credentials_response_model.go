// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppGroupCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppGroupCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppGroupCredentialsResponseBody) *CreateAppGroupCredentialsResponse
	GetBody() *CreateAppGroupCredentialsResponseBody
}

type CreateAppGroupCredentialsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppGroupCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppGroupCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupCredentialsResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppGroupCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppGroupCredentialsResponse) GetBody() *CreateAppGroupCredentialsResponseBody {
	return s.Body
}

func (s *CreateAppGroupCredentialsResponse) SetHeaders(v map[string]*string) *CreateAppGroupCredentialsResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGroupCredentialsResponse) SetStatusCode(v int32) *CreateAppGroupCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppGroupCredentialsResponse) SetBody(v *CreateAppGroupCredentialsResponseBody) *CreateAppGroupCredentialsResponse {
	s.Body = v
	return s
}

func (s *CreateAppGroupCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
