// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTempFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTempFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateTempFileResponseBody) *CreateTempFileResponse
	GetBody() *CreateTempFileResponseBody
}

type CreateTempFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTempFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTempFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileResponse) GoString() string {
	return s.String()
}

func (s *CreateTempFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTempFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTempFileResponse) GetBody() *CreateTempFileResponseBody {
	return s.Body
}

func (s *CreateTempFileResponse) SetHeaders(v map[string]*string) *CreateTempFileResponse {
	s.Headers = v
	return s
}

func (s *CreateTempFileResponse) SetStatusCode(v int32) *CreateTempFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTempFileResponse) SetBody(v *CreateTempFileResponseBody) *CreateTempFileResponse {
	s.Body = v
	return s
}

func (s *CreateTempFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
