// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityProxyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityProxyResponseBody) *DeleteSecurityProxyResponse
	GetBody() *DeleteSecurityProxyResponseBody
}

type DeleteSecurityProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityProxyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityProxyResponse) GetBody() *DeleteSecurityProxyResponseBody {
	return s.Body
}

func (s *DeleteSecurityProxyResponse) SetHeaders(v map[string]*string) *DeleteSecurityProxyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityProxyResponse) SetStatusCode(v int32) *DeleteSecurityProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityProxyResponse) SetBody(v *DeleteSecurityProxyResponseBody) *DeleteSecurityProxyResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
