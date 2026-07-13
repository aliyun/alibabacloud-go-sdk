// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetTokenTrendResponseBody) *GetTokenTrendResponse
	GetBody() *GetTokenTrendResponseBody
}

type GetTokenTrendResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenTrendResponse) GoString() string {
	return s.String()
}

func (s *GetTokenTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenTrendResponse) GetBody() *GetTokenTrendResponseBody {
	return s.Body
}

func (s *GetTokenTrendResponse) SetHeaders(v map[string]*string) *GetTokenTrendResponse {
	s.Headers = v
	return s
}

func (s *GetTokenTrendResponse) SetStatusCode(v int32) *GetTokenTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenTrendResponse) SetBody(v *GetTokenTrendResponseBody) *GetTokenTrendResponse {
	s.Body = v
	return s
}

func (s *GetTokenTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
