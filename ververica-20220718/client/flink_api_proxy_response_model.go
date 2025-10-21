// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlinkApiProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlinkApiProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlinkApiProxyResponse
	GetStatusCode() *int32
	SetBody(v *FlinkApiProxyResponseBody) *FlinkApiProxyResponse
	GetBody() *FlinkApiProxyResponseBody
}

type FlinkApiProxyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlinkApiProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlinkApiProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s FlinkApiProxyResponse) GoString() string {
	return s.String()
}

func (s *FlinkApiProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlinkApiProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlinkApiProxyResponse) GetBody() *FlinkApiProxyResponseBody {
	return s.Body
}

func (s *FlinkApiProxyResponse) SetHeaders(v map[string]*string) *FlinkApiProxyResponse {
	s.Headers = v
	return s
}

func (s *FlinkApiProxyResponse) SetStatusCode(v int32) *FlinkApiProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *FlinkApiProxyResponse) SetBody(v *FlinkApiProxyResponseBody) *FlinkApiProxyResponse {
	s.Body = v
	return s
}

func (s *FlinkApiProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
