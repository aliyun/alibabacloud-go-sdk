// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentInfoResponseBody) *GetAgentInfoResponse
	GetBody() *GetAgentInfoResponseBody
}

type GetAgentInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAgentInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentInfoResponse) GetBody() *GetAgentInfoResponseBody {
	return s.Body
}

func (s *GetAgentInfoResponse) SetHeaders(v map[string]*string) *GetAgentInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAgentInfoResponse) SetStatusCode(v int32) *GetAgentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentInfoResponse) SetBody(v *GetAgentInfoResponseBody) *GetAgentInfoResponse {
	s.Body = v
	return s
}

func (s *GetAgentInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
