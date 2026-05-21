// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTempFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTempFileResponse
	GetStatusCode() *int32
	SetBody(v *GetTempFileResponseBody) *GetTempFileResponse
	GetBody() *GetTempFileResponseBody
}

type GetTempFileResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTempFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTempFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTempFileResponse) GoString() string {
	return s.String()
}

func (s *GetTempFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTempFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTempFileResponse) GetBody() *GetTempFileResponseBody {
	return s.Body
}

func (s *GetTempFileResponse) SetHeaders(v map[string]*string) *GetTempFileResponse {
	s.Headers = v
	return s
}

func (s *GetTempFileResponse) SetStatusCode(v int32) *GetTempFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTempFileResponse) SetBody(v *GetTempFileResponseBody) *GetTempFileResponse {
	s.Body = v
	return s
}

func (s *GetTempFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
