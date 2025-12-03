// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppKeyResponseBody) *CreateAppKeyResponse
	GetBody() *CreateAppKeyResponseBody
}

type CreateAppKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppKeyResponse) GetBody() *CreateAppKeyResponseBody {
	return s.Body
}

func (s *CreateAppKeyResponse) SetHeaders(v map[string]*string) *CreateAppKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAppKeyResponse) SetStatusCode(v int32) *CreateAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppKeyResponse) SetBody(v *CreateAppKeyResponseBody) *CreateAppKeyResponse {
	s.Body = v
	return s
}

func (s *CreateAppKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
