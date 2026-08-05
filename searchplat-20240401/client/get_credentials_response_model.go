// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *GetCredentialsResponseBody) *GetCredentialsResponse
	GetBody() *GetCredentialsResponseBody
}

type GetCredentialsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCredentialsResponse) GetBody() *GetCredentialsResponseBody {
	return s.Body
}

func (s *GetCredentialsResponse) SetHeaders(v map[string]*string) *GetCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetCredentialsResponse) SetStatusCode(v int32) *GetCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCredentialsResponse) SetBody(v *GetCredentialsResponseBody) *GetCredentialsResponse {
	s.Body = v
	return s
}

func (s *GetCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
