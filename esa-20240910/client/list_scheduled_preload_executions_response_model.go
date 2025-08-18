// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPreloadExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduledPreloadExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduledPreloadExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduledPreloadExecutionsResponseBody) *ListScheduledPreloadExecutionsResponse
	GetBody() *ListScheduledPreloadExecutionsResponseBody
}

type ListScheduledPreloadExecutionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPreloadExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPreloadExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduledPreloadExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduledPreloadExecutionsResponse) GetBody() *ListScheduledPreloadExecutionsResponseBody {
	return s.Body
}

func (s *ListScheduledPreloadExecutionsResponse) SetHeaders(v map[string]*string) *ListScheduledPreloadExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPreloadExecutionsResponse) SetStatusCode(v int32) *ListScheduledPreloadExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponse) SetBody(v *ListScheduledPreloadExecutionsResponseBody) *ListScheduledPreloadExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListScheduledPreloadExecutionsResponse) Validate() error {
	return dara.Validate(s)
}
