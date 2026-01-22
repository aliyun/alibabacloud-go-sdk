// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSecretResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSecretResponseBody) *GetAppSecretResponse
	GetBody() *GetAppSecretResponseBody
}

type GetAppSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSecretResponse) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSecretResponse) GetBody() *GetAppSecretResponseBody {
	return s.Body
}

func (s *GetAppSecretResponse) SetHeaders(v map[string]*string) *GetAppSecretResponse {
	s.Headers = v
	return s
}

func (s *GetAppSecretResponse) SetStatusCode(v int32) *GetAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSecretResponse) SetBody(v *GetAppSecretResponseBody) *GetAppSecretResponse {
	s.Body = v
	return s
}

func (s *GetAppSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
