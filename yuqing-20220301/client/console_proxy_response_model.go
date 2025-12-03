// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConsoleProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConsoleProxyResponse
	GetStatusCode() *int32
	SetBody(v *ConsoleProxyResponseBody) *ConsoleProxyResponse
	GetBody() *ConsoleProxyResponseBody
}

type ConsoleProxyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsoleProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s ConsoleProxyResponse) GoString() string {
	return s.String()
}

func (s *ConsoleProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConsoleProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConsoleProxyResponse) GetBody() *ConsoleProxyResponseBody {
	return s.Body
}

func (s *ConsoleProxyResponse) SetHeaders(v map[string]*string) *ConsoleProxyResponse {
	s.Headers = v
	return s
}

func (s *ConsoleProxyResponse) SetStatusCode(v int32) *ConsoleProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsoleProxyResponse) SetBody(v *ConsoleProxyResponseBody) *ConsoleProxyResponse {
	s.Body = v
	return s
}

func (s *ConsoleProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
