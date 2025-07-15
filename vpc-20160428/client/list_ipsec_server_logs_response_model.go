// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpsecServerLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpsecServerLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpsecServerLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpsecServerLogsResponseBody) *ListIpsecServerLogsResponse
	GetBody() *ListIpsecServerLogsResponseBody
}

type ListIpsecServerLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpsecServerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpsecServerLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServerLogsResponse) GoString() string {
	return s.String()
}

func (s *ListIpsecServerLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpsecServerLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpsecServerLogsResponse) GetBody() *ListIpsecServerLogsResponseBody {
	return s.Body
}

func (s *ListIpsecServerLogsResponse) SetHeaders(v map[string]*string) *ListIpsecServerLogsResponse {
	s.Headers = v
	return s
}

func (s *ListIpsecServerLogsResponse) SetStatusCode(v int32) *ListIpsecServerLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpsecServerLogsResponse) SetBody(v *ListIpsecServerLogsResponseBody) *ListIpsecServerLogsResponse {
	s.Body = v
	return s
}

func (s *ListIpsecServerLogsResponse) Validate() error {
	return dara.Validate(s)
}
