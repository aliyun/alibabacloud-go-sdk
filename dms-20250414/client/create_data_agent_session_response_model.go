// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAgentSessionResponseBody) *CreateDataAgentSessionResponse
	GetBody() *CreateDataAgentSessionResponseBody
}

type CreateDataAgentSessionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAgentSessionResponse) GetBody() *CreateDataAgentSessionResponseBody {
	return s.Body
}

func (s *CreateDataAgentSessionResponse) SetHeaders(v map[string]*string) *CreateDataAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAgentSessionResponse) SetStatusCode(v int32) *CreateDataAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAgentSessionResponse) SetBody(v *CreateDataAgentSessionResponseBody) *CreateDataAgentSessionResponse {
	s.Body = v
	return s
}

func (s *CreateDataAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
