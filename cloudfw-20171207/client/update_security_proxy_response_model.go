// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecurityProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecurityProxyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecurityProxyResponseBody) *UpdateSecurityProxyResponse
	GetBody() *UpdateSecurityProxyResponseBody
}

type UpdateSecurityProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecurityProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecurityProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityProxyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecurityProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecurityProxyResponse) GetBody() *UpdateSecurityProxyResponseBody {
	return s.Body
}

func (s *UpdateSecurityProxyResponse) SetHeaders(v map[string]*string) *UpdateSecurityProxyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityProxyResponse) SetStatusCode(v int32) *UpdateSecurityProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityProxyResponse) SetBody(v *UpdateSecurityProxyResponseBody) *UpdateSecurityProxyResponse {
	s.Body = v
	return s
}

func (s *UpdateSecurityProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
