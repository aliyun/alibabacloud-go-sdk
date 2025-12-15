// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootProxyResponse
	GetStatusCode() *int32
	SetBody(v *RebootProxyResponseBody) *RebootProxyResponse
	GetBody() *RebootProxyResponseBody
}

type RebootProxyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootProxyResponse) GoString() string {
	return s.String()
}

func (s *RebootProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootProxyResponse) GetBody() *RebootProxyResponseBody {
	return s.Body
}

func (s *RebootProxyResponse) SetHeaders(v map[string]*string) *RebootProxyResponse {
	s.Headers = v
	return s
}

func (s *RebootProxyResponse) SetStatusCode(v int32) *RebootProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootProxyResponse) SetBody(v *RebootProxyResponseBody) *RebootProxyResponse {
	s.Body = v
	return s
}

func (s *RebootProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
