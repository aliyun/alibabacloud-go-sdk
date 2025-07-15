// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyQueueEventLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLegacyQueueEventLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLegacyQueueEventLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListLegacyQueueEventLogsResponseBody) *ListLegacyQueueEventLogsResponse
	GetBody() *ListLegacyQueueEventLogsResponseBody
}

type ListLegacyQueueEventLogsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLegacyQueueEventLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLegacyQueueEventLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyQueueEventLogsResponse) GoString() string {
	return s.String()
}

func (s *ListLegacyQueueEventLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLegacyQueueEventLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLegacyQueueEventLogsResponse) GetBody() *ListLegacyQueueEventLogsResponseBody {
	return s.Body
}

func (s *ListLegacyQueueEventLogsResponse) SetHeaders(v map[string]*string) *ListLegacyQueueEventLogsResponse {
	s.Headers = v
	return s
}

func (s *ListLegacyQueueEventLogsResponse) SetStatusCode(v int32) *ListLegacyQueueEventLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLegacyQueueEventLogsResponse) SetBody(v *ListLegacyQueueEventLogsResponseBody) *ListLegacyQueueEventLogsResponse {
	s.Body = v
	return s
}

func (s *ListLegacyQueueEventLogsResponse) Validate() error {
	return dara.Validate(s)
}
