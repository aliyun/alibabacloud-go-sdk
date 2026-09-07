// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentMJobInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentMJobInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentMJobInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentMJobInfoResponseBody) *ListAgentMJobInfoResponse
	GetBody() *ListAgentMJobInfoResponseBody
}

type ListAgentMJobInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentMJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentMJobInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentMJobInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAgentMJobInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentMJobInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentMJobInfoResponse) GetBody() *ListAgentMJobInfoResponseBody {
	return s.Body
}

func (s *ListAgentMJobInfoResponse) SetHeaders(v map[string]*string) *ListAgentMJobInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAgentMJobInfoResponse) SetStatusCode(v int32) *ListAgentMJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentMJobInfoResponse) SetBody(v *ListAgentMJobInfoResponseBody) *ListAgentMJobInfoResponse {
	s.Body = v
	return s
}

func (s *ListAgentMJobInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
