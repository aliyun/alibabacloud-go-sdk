// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayErrorAccessLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayErrorAccessLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayErrorAccessLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayErrorAccessLogsResponseBody) *ListGatewayErrorAccessLogsResponse
	GetBody() *ListGatewayErrorAccessLogsResponseBody
}

type ListGatewayErrorAccessLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayErrorAccessLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayErrorAccessLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayErrorAccessLogsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayErrorAccessLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayErrorAccessLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayErrorAccessLogsResponse) GetBody() *ListGatewayErrorAccessLogsResponseBody {
	return s.Body
}

func (s *ListGatewayErrorAccessLogsResponse) SetHeaders(v map[string]*string) *ListGatewayErrorAccessLogsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayErrorAccessLogsResponse) SetStatusCode(v int32) *ListGatewayErrorAccessLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayErrorAccessLogsResponse) SetBody(v *ListGatewayErrorAccessLogsResponseBody) *ListGatewayErrorAccessLogsResponse {
	s.Body = v
	return s
}

func (s *ListGatewayErrorAccessLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
