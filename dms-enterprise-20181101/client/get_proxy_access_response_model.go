// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProxyAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProxyAccessResponse
	GetStatusCode() *int32
	SetBody(v *GetProxyAccessResponseBody) *GetProxyAccessResponse
	GetBody() *GetProxyAccessResponseBody
}

type GetProxyAccessResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProxyAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProxyAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProxyAccessResponse) GoString() string {
	return s.String()
}

func (s *GetProxyAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProxyAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProxyAccessResponse) GetBody() *GetProxyAccessResponseBody {
	return s.Body
}

func (s *GetProxyAccessResponse) SetHeaders(v map[string]*string) *GetProxyAccessResponse {
	s.Headers = v
	return s
}

func (s *GetProxyAccessResponse) SetStatusCode(v int32) *GetProxyAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProxyAccessResponse) SetBody(v *GetProxyAccessResponseBody) *GetProxyAccessResponse {
	s.Body = v
	return s
}

func (s *GetProxyAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
