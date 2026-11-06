// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentVersionResponseBody) *GetAgentVersionResponse
	GetBody() *GetAgentVersionResponseBody
}

type GetAgentVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAgentVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentVersionResponse) GetBody() *GetAgentVersionResponseBody {
	return s.Body
}

func (s *GetAgentVersionResponse) SetHeaders(v map[string]*string) *GetAgentVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAgentVersionResponse) SetStatusCode(v int32) *GetAgentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentVersionResponse) SetBody(v *GetAgentVersionResponseBody) *GetAgentVersionResponse {
	s.Body = v
	return s
}

func (s *GetAgentVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
