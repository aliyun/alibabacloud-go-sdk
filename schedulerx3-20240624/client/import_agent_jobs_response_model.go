// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAgentJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportAgentJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportAgentJobsResponse
	GetStatusCode() *int32
	SetId(v string) *ImportAgentJobsResponse
	GetId() *string
	SetEvent(v string) *ImportAgentJobsResponse
	GetEvent() *string
	SetBody(v *ImportAgentJobsResponseBody) *ImportAgentJobsResponse
	GetBody() *ImportAgentJobsResponseBody
}

type ImportAgentJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                      `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                      `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ImportAgentJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportAgentJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportAgentJobsResponse) GoString() string {
	return s.String()
}

func (s *ImportAgentJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportAgentJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportAgentJobsResponse) GetId() *string {
	return s.Id
}

func (s *ImportAgentJobsResponse) GetEvent() *string {
	return s.Event
}

func (s *ImportAgentJobsResponse) GetBody() *ImportAgentJobsResponseBody {
	return s.Body
}

func (s *ImportAgentJobsResponse) SetHeaders(v map[string]*string) *ImportAgentJobsResponse {
	s.Headers = v
	return s
}

func (s *ImportAgentJobsResponse) SetStatusCode(v int32) *ImportAgentJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAgentJobsResponse) SetId(v string) *ImportAgentJobsResponse {
	s.Id = &v
	return s
}

func (s *ImportAgentJobsResponse) SetEvent(v string) *ImportAgentJobsResponse {
	s.Event = &v
	return s
}

func (s *ImportAgentJobsResponse) SetBody(v *ImportAgentJobsResponseBody) *ImportAgentJobsResponse {
	s.Body = v
	return s
}

func (s *ImportAgentJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
