// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentMJobInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentMJobInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentMJobInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentMJobInfoResponseBody) *GetAgentMJobInfoResponse
	GetBody() *GetAgentMJobInfoResponseBody
}

type GetAgentMJobInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentMJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentMJobInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentMJobInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAgentMJobInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentMJobInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentMJobInfoResponse) GetBody() *GetAgentMJobInfoResponseBody {
	return s.Body
}

func (s *GetAgentMJobInfoResponse) SetHeaders(v map[string]*string) *GetAgentMJobInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAgentMJobInfoResponse) SetStatusCode(v int32) *GetAgentMJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentMJobInfoResponse) SetBody(v *GetAgentMJobInfoResponseBody) *GetAgentMJobInfoResponse {
	s.Body = v
	return s
}

func (s *GetAgentMJobInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
