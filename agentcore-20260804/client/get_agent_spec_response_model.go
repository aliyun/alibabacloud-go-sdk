// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSpecResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSpecResponseBody) *GetAgentSpecResponse
	GetBody() *GetAgentSpecResponseBody
}

type GetAgentSpecResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSpecResponse) GetBody() *GetAgentSpecResponseBody {
	return s.Body
}

func (s *GetAgentSpecResponse) SetHeaders(v map[string]*string) *GetAgentSpecResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSpecResponse) SetStatusCode(v int32) *GetAgentSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSpecResponse) SetBody(v *GetAgentSpecResponseBody) *GetAgentSpecResponse {
	s.Body = v
	return s
}

func (s *GetAgentSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
