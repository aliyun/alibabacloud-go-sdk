// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentListResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentListResponseBody) *GetAgentListResponse
	GetBody() *GetAgentListResponseBody
}

type GetAgentListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentListResponse) GoString() string {
	return s.String()
}

func (s *GetAgentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentListResponse) GetBody() *GetAgentListResponseBody {
	return s.Body
}

func (s *GetAgentListResponse) SetHeaders(v map[string]*string) *GetAgentListResponse {
	s.Headers = v
	return s
}

func (s *GetAgentListResponse) SetStatusCode(v int32) *GetAgentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentListResponse) SetBody(v *GetAgentListResponseBody) *GetAgentListResponse {
	s.Body = v
	return s
}

func (s *GetAgentListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
