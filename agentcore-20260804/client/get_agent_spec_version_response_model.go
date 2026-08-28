// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSpecVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSpecVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSpecVersionResponseBody) *GetAgentSpecVersionResponse
	GetBody() *GetAgentSpecVersionResponseBody
}

type GetAgentSpecVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSpecVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSpecVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSpecVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSpecVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSpecVersionResponse) GetBody() *GetAgentSpecVersionResponseBody {
	return s.Body
}

func (s *GetAgentSpecVersionResponse) SetHeaders(v map[string]*string) *GetAgentSpecVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSpecVersionResponse) SetStatusCode(v int32) *GetAgentSpecVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSpecVersionResponse) SetBody(v *GetAgentSpecVersionResponseBody) *GetAgentSpecVersionResponse {
	s.Body = v
	return s
}

func (s *GetAgentSpecVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
