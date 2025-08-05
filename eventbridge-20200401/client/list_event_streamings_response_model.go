// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventStreamingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventStreamingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventStreamingsResponse
	GetStatusCode() *int32
	SetBody(v *ListEventStreamingsResponseBody) *ListEventStreamingsResponse
	GetBody() *ListEventStreamingsResponseBody
}

type ListEventStreamingsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventStreamingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventStreamingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsResponse) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventStreamingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventStreamingsResponse) GetBody() *ListEventStreamingsResponseBody {
	return s.Body
}

func (s *ListEventStreamingsResponse) SetHeaders(v map[string]*string) *ListEventStreamingsResponse {
	s.Headers = v
	return s
}

func (s *ListEventStreamingsResponse) SetStatusCode(v int32) *ListEventStreamingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventStreamingsResponse) SetBody(v *ListEventStreamingsResponseBody) *ListEventStreamingsResponse {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponse) Validate() error {
	return dara.Validate(s)
}
