// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstallRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentInstallRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentInstallRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentInstallRecordsResponseBody) *ListAgentInstallRecordsResponse
	GetBody() *ListAgentInstallRecordsResponseBody
}

type ListAgentInstallRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentInstallRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentInstallRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstallRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentInstallRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentInstallRecordsResponse) GetBody() *ListAgentInstallRecordsResponseBody {
	return s.Body
}

func (s *ListAgentInstallRecordsResponse) SetHeaders(v map[string]*string) *ListAgentInstallRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentInstallRecordsResponse) SetStatusCode(v int32) *ListAgentInstallRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentInstallRecordsResponse) SetBody(v *ListAgentInstallRecordsResponseBody) *ListAgentInstallRecordsResponse {
	s.Body = v
	return s
}

func (s *ListAgentInstallRecordsResponse) Validate() error {
	return dara.Validate(s)
}
