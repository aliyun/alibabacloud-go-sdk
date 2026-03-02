// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPdpLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPdpLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListPdpLogsResponseBody) *ListPdpLogsResponse
	GetBody() *ListPdpLogsResponseBody
}

type ListPdpLogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPdpLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPdpLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLogsResponse) GoString() string {
	return s.String()
}

func (s *ListPdpLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPdpLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPdpLogsResponse) GetBody() *ListPdpLogsResponseBody {
	return s.Body
}

func (s *ListPdpLogsResponse) SetHeaders(v map[string]*string) *ListPdpLogsResponse {
	s.Headers = v
	return s
}

func (s *ListPdpLogsResponse) SetStatusCode(v int32) *ListPdpLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPdpLogsResponse) SetBody(v *ListPdpLogsResponseBody) *ListPdpLogsResponse {
	s.Body = v
	return s
}

func (s *ListPdpLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
