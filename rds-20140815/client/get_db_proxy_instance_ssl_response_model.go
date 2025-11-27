// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDbProxyInstanceSslResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDbProxyInstanceSslResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDbProxyInstanceSslResponse
	GetStatusCode() *int32
	SetBody(v *GetDbProxyInstanceSslResponseBody) *GetDbProxyInstanceSslResponse
	GetBody() *GetDbProxyInstanceSslResponseBody
}

type GetDbProxyInstanceSslResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDbProxyInstanceSslResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDbProxyInstanceSslResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDbProxyInstanceSslResponse) GoString() string {
	return s.String()
}

func (s *GetDbProxyInstanceSslResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDbProxyInstanceSslResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDbProxyInstanceSslResponse) GetBody() *GetDbProxyInstanceSslResponseBody {
	return s.Body
}

func (s *GetDbProxyInstanceSslResponse) SetHeaders(v map[string]*string) *GetDbProxyInstanceSslResponse {
	s.Headers = v
	return s
}

func (s *GetDbProxyInstanceSslResponse) SetStatusCode(v int32) *GetDbProxyInstanceSslResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDbProxyInstanceSslResponse) SetBody(v *GetDbProxyInstanceSslResponseBody) *GetDbProxyInstanceSslResponse {
	s.Body = v
	return s
}

func (s *GetDbProxyInstanceSslResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
