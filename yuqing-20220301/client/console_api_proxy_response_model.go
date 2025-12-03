// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleApiProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConsoleApiProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConsoleApiProxyResponse
	GetStatusCode() *int32
	SetBody(v *ConsoleApiProxyResponseBody) *ConsoleApiProxyResponse
	GetBody() *ConsoleApiProxyResponseBody
}

type ConsoleApiProxyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsoleApiProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleApiProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s ConsoleApiProxyResponse) GoString() string {
	return s.String()
}

func (s *ConsoleApiProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConsoleApiProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConsoleApiProxyResponse) GetBody() *ConsoleApiProxyResponseBody {
	return s.Body
}

func (s *ConsoleApiProxyResponse) SetHeaders(v map[string]*string) *ConsoleApiProxyResponse {
	s.Headers = v
	return s
}

func (s *ConsoleApiProxyResponse) SetStatusCode(v int32) *ConsoleApiProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsoleApiProxyResponse) SetBody(v *ConsoleApiProxyResponseBody) *ConsoleApiProxyResponse {
	s.Body = v
	return s
}

func (s *ConsoleApiProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
