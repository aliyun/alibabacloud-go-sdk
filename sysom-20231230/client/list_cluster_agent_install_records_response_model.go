// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAgentInstallRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterAgentInstallRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterAgentInstallRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterAgentInstallRecordsResponseBody) *ListClusterAgentInstallRecordsResponse
	GetBody() *ListClusterAgentInstallRecordsResponseBody
}

type ListClusterAgentInstallRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterAgentInstallRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterAgentInstallRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAgentInstallRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterAgentInstallRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterAgentInstallRecordsResponse) GetBody() *ListClusterAgentInstallRecordsResponseBody {
	return s.Body
}

func (s *ListClusterAgentInstallRecordsResponse) SetHeaders(v map[string]*string) *ListClusterAgentInstallRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterAgentInstallRecordsResponse) SetStatusCode(v int32) *ListClusterAgentInstallRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponse) SetBody(v *ListClusterAgentInstallRecordsResponseBody) *ListClusterAgentInstallRecordsResponse {
	s.Body = v
	return s
}

func (s *ListClusterAgentInstallRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
