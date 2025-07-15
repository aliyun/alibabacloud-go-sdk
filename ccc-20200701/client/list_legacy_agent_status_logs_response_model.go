// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAgentStatusLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLegacyAgentStatusLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLegacyAgentStatusLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListLegacyAgentStatusLogsResponseBody) *ListLegacyAgentStatusLogsResponse
	GetBody() *ListLegacyAgentStatusLogsResponseBody
}

type ListLegacyAgentStatusLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLegacyAgentStatusLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLegacyAgentStatusLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentStatusLogsResponse) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentStatusLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLegacyAgentStatusLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLegacyAgentStatusLogsResponse) GetBody() *ListLegacyAgentStatusLogsResponseBody {
	return s.Body
}

func (s *ListLegacyAgentStatusLogsResponse) SetHeaders(v map[string]*string) *ListLegacyAgentStatusLogsResponse {
	s.Headers = v
	return s
}

func (s *ListLegacyAgentStatusLogsResponse) SetStatusCode(v int32) *ListLegacyAgentStatusLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLegacyAgentStatusLogsResponse) SetBody(v *ListLegacyAgentStatusLogsResponseBody) *ListLegacyAgentStatusLogsResponse {
	s.Body = v
	return s
}

func (s *ListLegacyAgentStatusLogsResponse) Validate() error {
	return dara.Validate(s)
}
