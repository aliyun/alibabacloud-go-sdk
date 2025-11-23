// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProxyResponse
	GetStatusCode() *int32
	SetBody(v *CreateProxyResponseBody) *CreateProxyResponse
	GetBody() *CreateProxyResponseBody
}

type CreateProxyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyResponse) GoString() string {
	return s.String()
}

func (s *CreateProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProxyResponse) GetBody() *CreateProxyResponseBody {
	return s.Body
}

func (s *CreateProxyResponse) SetHeaders(v map[string]*string) *CreateProxyResponse {
	s.Headers = v
	return s
}

func (s *CreateProxyResponse) SetStatusCode(v int32) *CreateProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProxyResponse) SetBody(v *CreateProxyResponseBody) *CreateProxyResponse {
	s.Body = v
	return s
}

func (s *CreateProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
