// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityProxyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityProxyResponseBody) *CreateSecurityProxyResponse
	GetBody() *CreateSecurityProxyResponseBody
}

type CreateSecurityProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityProxyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityProxyResponse) GetBody() *CreateSecurityProxyResponseBody {
	return s.Body
}

func (s *CreateSecurityProxyResponse) SetHeaders(v map[string]*string) *CreateSecurityProxyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityProxyResponse) SetStatusCode(v int32) *CreateSecurityProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityProxyResponse) SetBody(v *CreateSecurityProxyResponseBody) *CreateSecurityProxyResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityProxyResponse) Validate() error {
	return dara.Validate(s)
}
