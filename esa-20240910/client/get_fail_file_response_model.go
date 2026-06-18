// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFailFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFailFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFailFileResponse
	GetStatusCode() *int32
	SetBody(v *GetFailFileResponseBody) *GetFailFileResponse
	GetBody() *GetFailFileResponseBody
}

type GetFailFileResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFailFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFailFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFailFileResponse) GoString() string {
	return s.String()
}

func (s *GetFailFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFailFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFailFileResponse) GetBody() *GetFailFileResponseBody {
	return s.Body
}

func (s *GetFailFileResponse) SetHeaders(v map[string]*string) *GetFailFileResponse {
	s.Headers = v
	return s
}

func (s *GetFailFileResponse) SetStatusCode(v int32) *GetFailFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFailFileResponse) SetBody(v *GetFailFileResponseBody) *GetFailFileResponse {
	s.Body = v
	return s
}

func (s *GetFailFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
