// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstanceLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstanceLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstanceLogsResponseBody) *ListServiceInstanceLogsResponse
	GetBody() *ListServiceInstanceLogsResponseBody
}

type ListServiceInstanceLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstanceLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstanceLogsResponse) GetBody() *ListServiceInstanceLogsResponseBody {
	return s.Body
}

func (s *ListServiceInstanceLogsResponse) SetHeaders(v map[string]*string) *ListServiceInstanceLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetStatusCode(v int32) *ListServiceInstanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetBody(v *ListServiceInstanceLogsResponseBody) *ListServiceInstanceLogsResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstanceLogsResponse) Validate() error {
	return dara.Validate(s)
}
