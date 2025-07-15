// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAgentEventLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLegacyAgentEventLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLegacyAgentEventLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListLegacyAgentEventLogsResponseBody) *ListLegacyAgentEventLogsResponse
	GetBody() *ListLegacyAgentEventLogsResponseBody
}

type ListLegacyAgentEventLogsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLegacyAgentEventLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLegacyAgentEventLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentEventLogsResponse) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentEventLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLegacyAgentEventLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLegacyAgentEventLogsResponse) GetBody() *ListLegacyAgentEventLogsResponseBody {
	return s.Body
}

func (s *ListLegacyAgentEventLogsResponse) SetHeaders(v map[string]*string) *ListLegacyAgentEventLogsResponse {
	s.Headers = v
	return s
}

func (s *ListLegacyAgentEventLogsResponse) SetStatusCode(v int32) *ListLegacyAgentEventLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLegacyAgentEventLogsResponse) SetBody(v *ListLegacyAgentEventLogsResponseBody) *ListLegacyAgentEventLogsResponse {
	s.Body = v
	return s
}

func (s *ListLegacyAgentEventLogsResponse) Validate() error {
	return dara.Validate(s)
}
