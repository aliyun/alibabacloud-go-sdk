// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ArchiveAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ArchiveAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *ArchiveAgentSessionResponseBody) *ArchiveAgentSessionResponse
	GetBody() *ArchiveAgentSessionResponseBody
}

type ArchiveAgentSessionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ArchiveAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ArchiveAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s ArchiveAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *ArchiveAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ArchiveAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ArchiveAgentSessionResponse) GetBody() *ArchiveAgentSessionResponseBody {
	return s.Body
}

func (s *ArchiveAgentSessionResponse) SetHeaders(v map[string]*string) *ArchiveAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *ArchiveAgentSessionResponse) SetStatusCode(v int32) *ArchiveAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *ArchiveAgentSessionResponse) SetBody(v *ArchiveAgentSessionResponseBody) *ArchiveAgentSessionResponse {
	s.Body = v
	return s
}

func (s *ArchiveAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
